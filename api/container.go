package api

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
	"github.com/harbourmaster/hbm/api/types"
	"github.com/harbourmaster/hbm/pkg/db"
)

func AllowContainerCreate(req authorization.Request, config *types.Config) (string, string) {
	if req.RequestBody == nil {
		return "", "Malformed request"
	}

	type CreateContainer struct {
		HostConfig container.HostConfig
	}
	cc := &CreateContainer{}

	if err := json.NewDecoder(bytes.NewReader(req.RequestBody)).Decode(cc); err != nil {
		return "", err.Error()
	}

	defer db.RecoverFunc()

	d, err := db.NewDB(config.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer d.Conn.Close()

	if cc.HostConfig.Privileged {
		if ! d.KeyExists("config", "container_create_privileged") {
			return "--privileged param is not allowed", ""
		}
	}

	if cc.HostConfig.IpcMode == "host" {
		if ! d.KeyExists("config", "container_create_ipc_host") {
			return "--ipc=\"host\" param is not allowed", ""
		}
	}

	if cc.HostConfig.NetworkMode == "host" {
		if ! d.KeyExists("config", "container_create_net_host") {
			return "--net=\"host\" param is not allowed", ""
		}
	}

	if cc.HostConfig.PidMode == "host" {
		if ! d.KeyExists("config", "container_create_pid_host") {
			return "--pid=\"host\" param is not allowed", ""
		}
	}

	if cc.HostConfig.UsernsMode == "host" {
		if ! d.KeyExists("config", "container_create_userns_host") {
			return "--userns=\"host\" param is not allowed", ""
		}
	}

	if cc.HostConfig.UTSMode == "host" {
		if ! d.KeyExists("config", "container_create_uts_host") {
			return "--uts=\"host\" param is not allowed", ""
		}
	}

	if len(cc.HostConfig.CapAdd) > 0 {
		for _, c := range cc.HostConfig.CapAdd {
			if ! d.KeyExists("cap", c) {
				return fmt.Sprintf("Capability %s is not allowed", c), ""
			}
		}
	}

	if len(cc.HostConfig.Devices) > 0 {
		for _, dev := range cc.HostConfig.Devices {
			if ! d.KeyExists("device", dev.PathOnHost) {
				return fmt.Sprintf("Device %s is not allowed to be exported", dev.PathOnHost), ""
			}
		}
	}

	if len(cc.HostConfig.DNS) > 0 {
		for _, dns := range cc.HostConfig.DNS {
			if ! d.KeyExists("dns", dns) {
				return fmt.Sprintf("DNS server %s is not allowed", dns), ""
			}
		}
	}

	if len(cc.HostConfig.PortBindings) > 0 {
                for _, pbs := range cc.HostConfig.PortBindings {
			for _, pb := range pbs {
				spb := GetPortBindingString(&pb)

				if ! d.KeyExists("port", spb) {
					return fmt.Sprintf("Port %s is not allowed to be pubished", spb), ""
				}
			}
                }
        }

	if len(cc.HostConfig.Binds) > 0 {
		d.Conn.Close()

		for _, b := range cc.HostConfig.Binds {
			vol := strings.Split(b, ":")

			if ! AllowVolume(vol[0], config) {
				return fmt.Sprintf("Volume %s is not allowed to be mounted", b), ""
			}
		}
	}

	return "", ""
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
