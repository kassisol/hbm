package allow

import (
	"fmt"
	"path"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/juliengk/go-log"
	logdriver "github.com/juliengk/go-log/driver"
	"github.com/juliengk/go-mount"
	"github.com/juliengk/go-utils"
	"github.com/juliengk/go-utils/json"
	"github.com/kassisol/hbm/allow/types"
	"github.com/kassisol/hbm/storage"
	"github.com/kassisol/hbm/storage/driver"
	"github.com/kassisol/hbm/version"
)

func AllowContainerCreate(req authorization.Request, config *types.Config) *types.AllowResult {
	type ContainerCreateConfig struct {
		container.Config
		HostConfig container.HostConfig
	}

	l, _ := log.NewDriver("standard", nil)

	cc := &ContainerCreateConfig{}

	if err := json.Decode(req.RequestBody, cc); err != nil {
		return &types.AllowResult{Allow: false, Error: err.Error()}
	}

	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", config.AppPath)
	if err != nil {
		l.WithFields(logdriver.Fields{
			"storagedriver": "sqlite",
			"logdriver":     "standard",
			"version":       version.Version,
		}).Fatal(err)
	}
	defer s.End()

	if cc.HostConfig.Privileged {
		if !s.ValidatePolicy(config.Username, "config", "container_create_privileged", "") {
			return &types.AllowResult{Allow: false, Msg: "--privileged param is not allowed"}
		}
	}

	if cc.HostConfig.IpcMode == "host" {
		if !s.ValidatePolicy(config.Username, "config", "container_create_ipc_host", "") {
			return &types.AllowResult{Allow: false, Msg: "--ipc=\"host\" param is not allowed"}
		}
	}

	if cc.HostConfig.NetworkMode == "host" {
		if !s.ValidatePolicy(config.Username, "config", "container_create_net_host", "") {
			return &types.AllowResult{Allow: false, Msg: "--net=\"host\" param is not allowed"}
		}
	}

	if cc.HostConfig.PidMode == "host" {
		if !s.ValidatePolicy(config.Username, "config", "container_create_pid_host", "") {
			return &types.AllowResult{Allow: false, Msg: "--pid=\"host\" param is not allowed"}
		}
	}

	if cc.HostConfig.UsernsMode == "host" {
		if !s.ValidatePolicy(config.Username, "config", "container_create_userns_host", "") {
			return &types.AllowResult{Allow: false, Msg: "--userns=\"host\" param is not allowed"}
		}
	}

	if cc.HostConfig.UTSMode == "host" {
		if !s.ValidatePolicy(config.Username, "config", "container_create_uts_host", "") {
			return &types.AllowResult{Allow: false, Msg: "--uts=\"host\" param is not allowed"}
		}
	}

	if len(cc.HostConfig.CapAdd) > 0 {
		for _, c := range cc.HostConfig.CapAdd {
			if !s.ValidatePolicy(config.Username, "cap", c, "") {
				return &types.AllowResult{Allow: false, Msg: fmt.Sprintf("Capability %s is not allowed", c)}
			}
		}
	}

	if len(cc.HostConfig.Devices) > 0 {
		for _, dev := range cc.HostConfig.Devices {
			if !s.ValidatePolicy(config.Username, "device", dev.PathOnHost, "") {
				return &types.AllowResult{Allow: false, Msg: fmt.Sprintf("Device %s is not allowed to be exported", dev.PathOnHost)}
			}
		}
	}

	if len(cc.HostConfig.DNS) > 0 {
		for _, dns := range cc.HostConfig.DNS {
			if !s.ValidatePolicy(config.Username, "dns", dns, "") {
				return &types.AllowResult{Allow: false, Msg: fmt.Sprintf("DNS server %s is not allowed", dns)}
			}
		}
	}

	if cc.HostConfig.PublishAllPorts {
		if !s.ValidatePolicy(config.Username, "config", "container_publish_all", "") {
			return &types.AllowResult{Allow: false, Msg: "--publish-all param is not allowed"}
		}
	}

	if len(cc.HostConfig.PortBindings) > 0 {
		for _, pbs := range cc.HostConfig.PortBindings {
			for _, pb := range pbs {
				spb := GetPortBindingString(&pb)

				if !s.ValidatePolicy(config.Username, "port", spb, "") {
					return &types.AllowResult{Allow: false, Msg: fmt.Sprintf("Port %s is not allowed to be pubished", spb)}
				}
			}
		}
	}

	if len(cc.HostConfig.Binds) > 0 {
		s.End()

		for _, b := range cc.HostConfig.Binds {
			vol := strings.Split(b, ":")

			if !AllowVolume(vol[0], config) {
				return &types.AllowResult{Allow: false, Msg: fmt.Sprintf("Volume %s is not allowed to be mounted", b)}
			}
		}
	}

	if len(cc.HostConfig.LogConfig.Type) > 0 {
		if !s.ValidatePolicy(config.Username, "logdriver", cc.HostConfig.LogConfig.Type, "") {
			return &types.AllowResult{Allow: false, Msg: fmt.Sprintf("Log driver %s is not allowed", cc.HostConfig.LogConfig.Type)}
		}
	}

	if len(cc.HostConfig.LogConfig.Config) > 0 {
		for k, v := range cc.HostConfig.LogConfig.Config {
			los := fmt.Sprintf("%s=%s", k, v)

			if !s.ValidatePolicy(config.Username, "logopt", los, "") {
				return &types.AllowResult{Allow: false, Msg: fmt.Sprintf("Log driver %s is not allowed", los)}
			}
		}
	}

	if len(cc.User) > 0 {
		if cc.Config.User == "root" && !s.ValidatePolicy(config.Username, "config", "container_create_user_root", "") {
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
	defer utils.RecoverFunc()

	l, _ := log.NewDriver("standard", nil)

	s, err := storage.NewDriver("sqlite", config.AppPath)
	if err != nil {
		l.WithFields(logdriver.Fields{
			"storagedriver": "sqlite",
			"logdriver":     "standard",
			"version":       version.Version,
		}).Fatal(err)
	}
	defer s.End()

	vo := driver.VolumeOptions{
		Recursive: false,
	}
	if AllowMount(vol) {
		vo.NoSuid = true
	}
	jsonVO := json.Encode(vo)
	opts := jsonVO.String()

	if s.ValidatePolicy(config.Username, "volume", vol, opts) {
		return true
	}

	v := strings.Split(vol, "/")

	p := make([]string, len(v))
	p[0] = "/"

	for i := 1; i < len(v); i++ {
		p = append(p, v[i])

		vo = driver.VolumeOptions{
			Recursive: true,
		}
		if AllowMount(vol) {
			vo.NoSuid = true
		} else {
			vo.NoSuid = false
		}
		jsonVO = json.Encode(vo)
		opts = jsonVO.String()

		if s.ValidatePolicy(config.Username, "volume", path.Join(p...), opts) {
			return true
		}
	}

	return false
}

func AllowMount(vol string) bool {
	result := false

	entries, err := mount.New()
	if err != nil {
		return false
	}

	entry, err := entries.Find(vol)
	if err == nil {
		result = entry.FindOption("nosuid")
	}

	return result
}
