package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"path"
	"strings"

	//	"github.com/davecgh/go-spew/spew"
	"github.com/docker/engine-api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/harbourmaster/hbm/api/types"
	"github.com/harbourmaster/hbm/pkg/db"
)

type CreateContainerConfig struct {
	Config     container.Config
	HostConfig container.HostConfig
}

func AllowContainerCreate(req authorization.Request, config *types.Config) *types.AllowResult {
	if req.User != "" {
		for _, val := range []string{"config", "cap", "device", "dns", "port"} {
			initUserBucket(val, req.User, config.AppPath)
		}
	}

	if req.RequestBody == nil {
		return &types.AllowResult{Allow: false, Error: "Malformed request"}
	}

	cc := &CreateContainerConfig{}

	if req.RequestBody == nil {
		return &types.AllowResult{Allow: false, Error: "Malformed request"}
	}
	if err := json.NewDecoder(bytes.NewReader(req.RequestBody)).Decode(cc); err != nil {
		return &types.AllowResult{Allow: false, Error: err.Error()}
	}
	//	log.Debug(spew.Sdump(cc))

	result := validateContainerConfig(req.User, cc, config)
	if result.Allow {
		return result
	}

	if req.User == "" {
		return &types.AllowResult{Allow: false, Msg: result.Msg + " for anonymous user"}
	} else {
		return &types.AllowResult{Allow: false, Msg: result.Msg + " for user " + req.User}
	}
}

func validateContainerConfig(user string, cc *CreateContainerConfig, config *types.Config) *types.AllowResult {

	configBuckets := []string{"config"}
	capBuckets := []string{"cap"}
	deviceBuckets := []string{"device"}
	portBuckets := []string{"port"}
	volumeBuckets := []string{"volume"}

	if user != "" {
		configBuckets = append(configBuckets, configBuckets[0]+"_"+user)
		capBuckets = append(capBuckets, capBuckets[0]+"_"+user)
		deviceBuckets = append(deviceBuckets, deviceBuckets[0]+"_"+user)
		volumeBuckets = append(volumeBuckets, volumeBuckets[0]+"_"+user)
	}

	defer db.RecoverFunc()

	d, err := db.NewDB(config.AppPath)
	if err != nil {
		log.Debug("In validateContainerConfig")
		log.Fatal(err)
	}
	defer d.Conn.Close()

	if cc.HostConfig.Privileged {
		if !d.KeyExistsInBuckets(configBuckets, "container_create_privileged") {
			return &types.AllowResult{Allow: false, Msg: "--privileged param is not allowed"}
		}
	}

	if cc.HostConfig.IpcMode == "host" {
		if !d.KeyExistsInBuckets(configBuckets, "container_create_ipc_host") {
			return &types.AllowResult{Allow: false, Msg: "--ipc=\"host\" param is not allowed"}
		}
	}

	if cc.HostConfig.NetworkMode == "host" {
		if !d.KeyExistsInBuckets(configBuckets, "container_create_net_host") {
			return &types.AllowResult{Allow: false, Msg: "--net=\"host\" param is not allowed"}
		}
	}

	if cc.HostConfig.PidMode == "host" {
		if !d.KeyExistsInBuckets(configBuckets, "container_create_pid_host") {
			return &types.AllowResult{Allow: false, Msg: "--pid=\"host\" param is not allowed"}
		}
	}

	if cc.HostConfig.UsernsMode == "host" {
		if !d.KeyExistsInBuckets(configBuckets, "container_create_userns_host") {
			return &types.AllowResult{Allow: false, Msg: "--userns=\"host\" param is not allowed"}
		}
	}

	if cc.HostConfig.UTSMode == "host" {
		if !d.KeyExistsInBuckets(configBuckets, "container_create_uts_host") {
			return &types.AllowResult{Allow: false, Msg: "--uts=\"host\" param is not allowed"}
		}
	}

	if len(cc.HostConfig.CapAdd) > 0 {
		for _, c := range cc.HostConfig.CapAdd {
			if !d.KeyExistsInBuckets(capBuckets, c) {
				return &types.AllowResult{Allow: false, Msg: fmt.Sprintf("Capability %s is not allowed", c)}
			}
		}
	}

	if len(cc.HostConfig.Devices) > 0 {
		for _, dev := range cc.HostConfig.Devices {
			if !d.KeyExistsInBuckets(deviceBuckets, dev.PathOnHost) {
				return &types.AllowResult{Allow: false, Msg: fmt.Sprintf("Device %s is not allowed to be exported", dev.PathOnHost)}
			}
		}
	}

	if len(cc.HostConfig.DNS) > 0 {
		for _, dns := range cc.HostConfig.DNS {
			if !d.KeyExists("dns", dns) {
				return &types.AllowResult{Allow: false, Msg: fmt.Sprintf("DNS server %s is not allowed", dns)}
			}
		}
	}

	if len(cc.HostConfig.PortBindings) > 0 {
		for _, pbs := range cc.HostConfig.PortBindings {
			for _, pb := range pbs {
				spb := GetPortBindingString(&pb)

				if !d.KeyExistsInBuckets(portBuckets, spb) {
					return &types.AllowResult{Allow: false, Msg: fmt.Sprintf("Port %s is not allowed to be pubished", spb)}
				}
			}
		}
	}

	if len(cc.HostConfig.Binds) > 0 {
		d.Conn.Close()

		for _, b := range cc.HostConfig.Binds {
			vol := strings.Split(b, ":")

			if !AllowVolume(volumeBuckets, vol[0], config) {
				return &types.AllowResult{Allow: false, Msg: fmt.Sprintf("Volume %s is not allowed to be mounted", b)}
			}
		}
	}

	if len(cc.Config.User) > 0 {
		if cc.Config.User == "root" && !d.KeyExistsInBuckets(configBuckets, "container_create_user_root") {
			return &types.AllowResult{Allow: false, Msg: "Running as user \"root\" is not allowed. Please use --user=\"someuser\" param."}
		}
	}

	return &types.AllowResult{Allow: true}
}

func GetPortBindingString(pb *nat.PortBinding) string {
	result := pb.HostPort

	if len(pb.HostIP) > 0 {
		result = fmt.Sprintf("%s:%s", pb.HostIP, pb.HostPort)
	}

	return result
}

func AllowVolume(volumeBuckets []string, vol string, config *types.Config) bool {
	defer db.RecoverFunc()

	d, err := db.NewDB(config.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer d.Conn.Close()

	if d.KeyExistsInBuckets(volumeBuckets, vol) {
		log.Debug("Passing on allowed volume")
		return true
	}

	v := strings.Split(vol, "/")

	p := make([]string, len(v))
	p[0] = "/"

	for i := 1; i < len(v); i++ {
		p = append(p, v[i])

		if d.KeyExistsInBucketsRecursive(volumeBuckets, path.Join(p...)) {
			log.Debug("Passing on allowed volume")
			return true
		}
	}

	return false
}
