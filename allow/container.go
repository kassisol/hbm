package allow

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"path"
	"strings"

	"github.com/docker/engine-api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/allow/types"
	"github.com/kassisol/hbm/pkg/db"
)

func AllowContainerCreate(req authorization.Request, config *types.Config) *types.AllowResult {
	type ContainerCreateConfig struct {
		container.Config
		HostConfig container.HostConfig
	}
	cc := &ContainerCreateConfig{}

	if req.RequestBody == nil {
		return &types.AllowResult{Allow: false, Error: "Malformed request"}
	}
	if err := json.NewDecoder(bytes.NewReader(req.RequestBody)).Decode(cc); err != nil {
		return &types.AllowResult{Allow: false, Error: err.Error()}
	}

	defer db.RecoverFunc()

	d, err := db.NewDB(config.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer d.Conn.Close()

	if cc.HostConfig.Privileged {
		if !d.KeyExists("config", "container_create_privileged") {
			return &types.AllowResult{Allow: false, Msg: "--privileged param is not allowed"}
		}
	}

	if cc.HostConfig.IpcMode == "host" {
		if !d.KeyExists("config", "container_create_ipc_host") {
			return &types.AllowResult{Allow: false, Msg: "--ipc=\"host\" param is not allowed"}
		}
	}

	if cc.HostConfig.NetworkMode == "host" {
		if !d.KeyExists("config", "container_create_net_host") {
			return &types.AllowResult{Allow: false, Msg: "--net=\"host\" param is not allowed"}
		}
	}

	if cc.HostConfig.PidMode == "host" {
		if !d.KeyExists("config", "container_create_pid_host") {
			return &types.AllowResult{Allow: false, Msg: "--pid=\"host\" param is not allowed"}
		}
	}

	if cc.HostConfig.UsernsMode == "host" {
		if !d.KeyExists("config", "container_create_userns_host") {
			return &types.AllowResult{Allow: false, Msg: "--userns=\"host\" param is not allowed"}
		}
	}

	if cc.HostConfig.UTSMode == "host" {
		if !d.KeyExists("config", "container_create_uts_host") {
			return &types.AllowResult{Allow: false, Msg: "--uts=\"host\" param is not allowed"}
		}
	}

	if len(cc.HostConfig.CapAdd) > 0 {
		for _, c := range cc.HostConfig.CapAdd {
			if !d.KeyExists("cap", c) {
				return &types.AllowResult{Allow: false, Msg: fmt.Sprintf("Capability %s is not allowed", c)}
			}
		}
	}

	if len(cc.HostConfig.Devices) > 0 {
		for _, dev := range cc.HostConfig.Devices {
			if !d.KeyExists("device", dev.PathOnHost) {
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

				if !d.KeyExists("port", spb) {
					return &types.AllowResult{Allow: false, Msg: fmt.Sprintf("Port %s is not allowed to be pubished", spb)}
				}
			}
		}
	}

	if len(cc.HostConfig.Binds) > 0 {
		d.Conn.Close()

		for _, b := range cc.HostConfig.Binds {
			vol := strings.Split(b, ":")

			if !AllowVolume(vol[0], config) {
				return &types.AllowResult{Allow: false, Msg: fmt.Sprintf("Volume %s is not allowed to be mounted", b)}
			}
		}
	}

	if len(cc.User) > 0 {
		if cc.Config.User == "root" && !d.KeyExists("config", "container_create_user_root") {
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

func AllowVolume(vol string, config *types.Config) bool {
	defer db.RecoverFunc()

	d, err := db.NewDB(config.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer d.Conn.Close()

	if d.KeyExists("volume", vol) {
		return true
	}

	v := strings.Split(vol, "/")

	p := make([]string, len(v))
	p[0] = "/"

	for i := 1; i < len(v); i++ {
		p = append(p, v[i])

		if d.KeyExistsRecursive("volume", path.Join(p...)) {
			return true
		}
	}

	return false
}
