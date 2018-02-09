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
	"github.com/kassisol/hbm/docker/allow/types"
	policyobj "github.com/kassisol/hbm/object/policy"
	objtypes "github.com/kassisol/hbm/object/types"
	"github.com/kassisol/hbm/version"
)

func ContainerCreate(req authorization.Request, config *types.Config) *types.AllowResult {
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

	p, err := policyobj.New("sqlite", config.AppPath)
	if err != nil {
		l.WithFields(logdriver.Fields{
			"storagedriver": "sqlite",
			"logdriver":     "standard",
			"version":       version.Version,
		}).Fatal(err)
	}
	defer p.End()

	if cc.HostConfig.Privileged {
		if !p.Validate(config.Username, "config", "container_create_param_privileged", "") {
			return &types.AllowResult{Allow: false, Msg: "--privileged param is not allowed"}
		}
	}

	if cc.HostConfig.IpcMode == "host" {
		if !p.Validate(config.Username, "config", "container_create_param_ipc_host", "") {
			return &types.AllowResult{Allow: false, Msg: "--ipc=\"host\" param is not allowed"}
		}
	}

	if cc.HostConfig.NetworkMode == "host" {
		if !p.Validate(config.Username, "config", "container_create_param_net_host", "") {
			return &types.AllowResult{Allow: false, Msg: "--net=\"host\" param is not allowed"}
		}
	}

	if cc.HostConfig.PidMode == "host" {
		if !p.Validate(config.Username, "config", "container_create_param_pid_host", "") {
			return &types.AllowResult{Allow: false, Msg: "--pid=\"host\" param is not allowed"}
		}
	}

	if cc.HostConfig.UsernsMode == "host" {
		if !p.Validate(config.Username, "config", "container_create_param_userns_host", "") {
			return &types.AllowResult{Allow: false, Msg: "--userns=\"host\" param is not allowed"}
		}
	}

	if cc.HostConfig.UTSMode == "host" {
		if !p.Validate(config.Username, "config", "container_create_param_uts_host", "") {
			return &types.AllowResult{Allow: false, Msg: "--uts=\"host\" param is not allowed"}
		}
	}

	if len(cc.HostConfig.SecurityOpt) > 0 {
		if !p.Validate(config.Username, "config", "container_create_param_security_opt", "") {
			return &types.AllowResult{Allow: false, Msg: "--security-opt param is not allowed"}
		}
	}

	if len(cc.HostConfig.Sysctls) > 0 {
		if !p.Validate(config.Username, "config", "container_create_param_sysctl", "") {
			return &types.AllowResult{Allow: false, Msg: "--sysctl param is not allowed"}
		}
	}

	if len(cc.HostConfig.CapAdd) > 0 {
		for _, c := range cc.HostConfig.CapAdd {
			if !p.Validate(config.Username, "capability", c, "") {
				return &types.AllowResult{Allow: false, Msg: fmt.Sprintf("Capability %s is not allowed", c)}
			}
		}
	}

	if len(cc.HostConfig.Devices) > 0 {
		for _, dev := range cc.HostConfig.Devices {
			if !p.Validate(config.Username, "device", dev.PathOnHost, "") {
				return &types.AllowResult{Allow: false, Msg: fmt.Sprintf("Device %s is not allowed to be exported", dev.PathOnHost)}
			}
		}
	}

	if len(cc.HostConfig.DNS) > 0 {
		for _, dns := range cc.HostConfig.DNS {
			if !p.Validate(config.Username, "dns", dns, "") {
				return &types.AllowResult{Allow: false, Msg: fmt.Sprintf("DNS server %s is not allowed", dns)}
			}
		}
	}

	if cc.HostConfig.PublishAllPorts {
		if !p.Validate(config.Username, "config", "container_create_param_publish_all", "") {
			return &types.AllowResult{Allow: false, Msg: "--publish-all param is not allowed"}
		}
	}

	if len(cc.HostConfig.PortBindings) > 0 {
		for _, pbs := range cc.HostConfig.PortBindings {
			for _, pb := range pbs {
				spb := GetPortBindingString(&pb)

				if !p.Validate(config.Username, "port", spb, "") {
					return &types.AllowResult{Allow: false, Msg: fmt.Sprintf("Port %s is not allowed to be published", spb)}
				}
			}
		}
	}

	if len(cc.HostConfig.Binds) > 0 {
		p.End()

		for _, b := range cc.HostConfig.Binds {
			vol := strings.Split(b, ":")

			if !AllowVolume(vol[0], config) {
				return &types.AllowResult{Allow: false, Msg: fmt.Sprintf("Volume %s is not allowed to be mounted", b)}
			}
		}
	}

	if len(cc.HostConfig.LogConfig.Type) > 0 {
		if !p.Validate(config.Username, "logdriver", cc.HostConfig.LogConfig.Type, "") {
			return &types.AllowResult{Allow: false, Msg: fmt.Sprintf("Log driver %s is not allowed", cc.HostConfig.LogConfig.Type)}
		}
	}

	if len(cc.HostConfig.LogConfig.Config) > 0 {
		for k, v := range cc.HostConfig.LogConfig.Config {
			los := fmt.Sprintf("%s=%s", k, v)

			if !p.Validate(config.Username, "logopt", los, "") {
				return &types.AllowResult{Allow: false, Msg: fmt.Sprintf("Log driver %s is not allowed", los)}
			}
		}
	}

	if len(cc.User) > 0 {
		if cc.Config.User == "root" && !p.Validate(config.Username, "config", "container_create_user_root", "") {
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

	p, err := policyobj.New("sqlite", config.AppPath)
	if err != nil {
		l.WithFields(logdriver.Fields{
			"storagedriver": "sqlite",
			"logdriver":     "standard",
			"version":       version.Version,
		}).Fatal(err)
	}
	defer p.End()

	vo := objtypes.VolumeOptions{
		Recursive: false,
	}
	if AllowMount(vol) {
		vo.NoSuid = true
	}
	jsonVO := json.Encode(vo)
	opts := jsonVO.String()

	if p.Validate(config.Username, "volume", vol, opts) {
		return true
	}

	v := strings.Split(vol, "/")

	val := make([]string, len(v))
	val[0] = "/"

	for i := 1; i < len(v); i++ {
		val = append(val, v[i])

		vo = objtypes.VolumeOptions{
			Recursive: true,
		}
		if AllowMount(vol) {
			vo.NoSuid = true
		} else {
			vo.NoSuid = false
		}
		jsonVO = json.Encode(vo)
		opts = jsonVO.String()

		if p.Validate(config.Username, "volume", path.Join(val...), opts) {
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
