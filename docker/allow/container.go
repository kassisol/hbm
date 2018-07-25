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

	if len(cc.HostConfig.Binds) > 0 {
		for _, b := range cc.HostConfig.Binds {
			vol := strings.Split(b, ":")

			if !AllowVolume(vol[0], config) {
				return &types.AllowResult{
					Allow: false,
					Msg: map[string]string{
						"text":           fmt.Sprintf("Volume %s is not allowed to be mounted", b),
						"resource_type":  "volume",
						"resource_value": b,
					},
				}
			}
		}
	}

	if len(cc.HostConfig.LogConfig.Type) > 0 {
		if !p.Validate(config.Username, "logdriver", cc.HostConfig.LogConfig.Type, "") {
			return &types.AllowResult{
				Allow: false,
				Msg: map[string]string{
					"text":           fmt.Sprintf("Log driver %s is not allowed", cc.HostConfig.LogConfig.Type),
					"resource_type":  "logdriver",
					"resource_value": cc.HostConfig.LogConfig.Type,
				},
			}
		}
	}

	if len(cc.HostConfig.LogConfig.Config) > 0 {
		for k, v := range cc.HostConfig.LogConfig.Config {
			los := fmt.Sprintf("%s=%s", k, v)

			if !p.Validate(config.Username, "logopt", los, "") {
				return &types.AllowResult{
					Allow: false,
					Msg: map[string]string{
						"text":           fmt.Sprintf("Log option %s is not allowed", los),
						"resource_type":  "logopt",
						"resource_value": los,
					},
				}
			}
		}
	}

	if cc.HostConfig.NetworkMode == "host" {
		if !p.Validate(config.Username, "config", "container_create_param_net_host", "") {
			return &types.AllowResult{
				Allow: false,
				Msg: map[string]string{
					"text":           "--net=\"host\" param is not allowed",
					"resource_type":  "config",
					"resource_value": "container_create_param_net_host",
				},
			}
		}
	}

	if len(cc.HostConfig.PortBindings) > 0 {
		for _, pbs := range cc.HostConfig.PortBindings {
			for _, pb := range pbs {
				spb := GetPortBindingString(&pb)

				if !p.Validate(config.Username, "port", spb, "") {
					return &types.AllowResult{
						Allow: false,
						Msg: map[string]string{
							"text":           fmt.Sprintf("Port %s is not allowed to be published", spb),
							"resource_type":  "port",
							"resource_value": spb,
						},
					}
				}
			}
		}
	}

	if len(cc.HostConfig.VolumeDriver) > 0 {
		if !p.Validate(config.Username, "volumedriver", cc.HostConfig.VolumeDriver, "") {
			return &types.AllowResult{
				Allow: false,
				Msg: map[string]string{
					"text":           fmt.Sprintf("Volume driver %s is not allowed", cc.HostConfig.VolumeDriver),
					"resource_type":  "volumedriver",
					"resource_value": cc.HostConfig.VolumeDriver,
				},
			}
		}
	}

	if len(cc.HostConfig.CapAdd) > 0 {
		for _, c := range cc.HostConfig.CapAdd {
			if !p.Validate(config.Username, "capability", c, "") {
				return &types.AllowResult{
					Allow: false,
					Msg: map[string]string{
						"text":           fmt.Sprintf("Capability %s is not allowed", c),
						"resource_type":  "capability",
						"resource_value": c,
					},
				}
			}
		}
	}

	if len(cc.HostConfig.DNS) > 0 {
		for _, dns := range cc.HostConfig.DNS {
			if !p.Validate(config.Username, "dns", dns, "") {
				return &types.AllowResult{
					Allow: false,
					Msg: map[string]string{
						"text":           fmt.Sprintf("DNS server %s is not allowed", dns),
						"resource_type":  "dns",
						"resource_value": dns,
					},
				}
			}
		}
	}

	if cc.HostConfig.IpcMode == "host" {
		if !p.Validate(config.Username, "config", "container_create_param_ipc_host", "") {
			return &types.AllowResult{
				Allow: false,
				Msg: map[string]string{
					"text":           "--ipc=\"host\" param is not allowed",
					"resource_type":  "config",
					"resource_value": "container_create_param_ipc_host",
				},
			}
		}
	}

	if cc.HostConfig.OomScoreAdj != 0 {
		if !p.Validate(config.Username, "config", "container_create_param_oom_score_adj", "") {
			return &types.AllowResult{
				Allow: false,
				Msg: map[string]string{
					"text":           "--oom-score-adj param is not allowed",
					"resource_type":  "config",
					"resource_value": "container_create_param_oom_score_adj",
				},
			}
		}
	}

	if cc.HostConfig.PidMode == "host" {
		if !p.Validate(config.Username, "config", "container_create_param_pid_host", "") {
			return &types.AllowResult{
				Allow: false,
				Msg: map[string]string{
					"text":           "--pid=\"host\" param is not allowed",
					"resource_type":  "config",
					"resource_value": "container_create_param_pid_host",
				},
			}
		}
	}

	if cc.HostConfig.Privileged {
		if !p.Validate(config.Username, "config", "container_create_param_privileged", "") {
			return &types.AllowResult{
				Allow: false,
				Msg: map[string]string{
					"text":           "--privileged param is not allowed",
					"resource_type":  "config",
					"resource_value": "container_create_param_privileged",
				},
			}
		}
	}

	if cc.HostConfig.PublishAllPorts {
		if !p.Validate(config.Username, "config", "container_create_param_publish_all", "") {
			return &types.AllowResult{
				Allow: false,
				Msg: map[string]string{
					"text":           "--publish-all param is not allowed",
					"resource_type":  "config",
					"resource_value": "container_create_param_publish_all",
				},
			}
		}
	}

	if len(cc.HostConfig.SecurityOpt) > 0 {
		if !p.Validate(config.Username, "config", "container_create_param_security_opt", "") {
			return &types.AllowResult{
				Allow: false,
				Msg: map[string]string{
					"text":           "--security-opt param is not allowed",
					"resource_type":  "config",
					"resource_value": "container_create_param_security_opt",
				},
			}
		}
	}

	if len(cc.HostConfig.Tmpfs) > 0 {
		if !p.Validate(config.Username, "config", "container_create_param_tmpfs", "") {
			return &types.AllowResult{
				Allow: false,
				Msg: map[string]string{
					"text":           "--tmpfs param is not allowed",
					"resource_type":  "config",
					"resource_value": "container_create_param_tmpfs",
				},
			}
		}
	}

	if cc.HostConfig.UTSMode == "host" {
		if !p.Validate(config.Username, "config", "container_create_param_uts_host", "") {
			return &types.AllowResult{
				Allow: false,
				Msg: map[string]string{
					"text":           "--uts=\"host\" param is not allowed",
					"resource_type":  "config",
					"resource_value": "container_create_param_uts_host",
				},
			}
		}
	}

	if cc.HostConfig.UsernsMode == "host" {
		if !p.Validate(config.Username, "config", "container_create_param_userns_host", "") {
			return &types.AllowResult{
				Allow: false,
				Msg: map[string]string{
					"text":           "--userns=\"host\" param is not allowed",
					"resource_type":  "config",
					"resource_value": "container_create_param_userns_host",
				},
			}
		}
	}

	if len(cc.HostConfig.Sysctls) > 0 {
		if !p.Validate(config.Username, "config", "container_create_param_sysctl", "") {
			return &types.AllowResult{
				Allow: false,
				Msg: map[string]string{
					"text":           "--sysctl param is not allowed",
					"resource_type":  "config",
					"resource_value": "container_create_param_sysctl",
				},
			}
		}
	}

	if len(cc.HostConfig.Runtime) > 0 {
		if !p.Validate(config.Username, "runtime", cc.HostConfig.Runtime, "") {
			return &types.AllowResult{
				Allow: false,
				Msg: map[string]string{
					"text":           fmt.Sprintf("Runtime %s is not allowed", cc.HostConfig.Runtime),
					"resource_type":  "runtime",
					"resource_value": cc.HostConfig.Runtime,
				},
			}
		}
	}

	if len(cc.HostConfig.Devices) > 0 {
		for _, dev := range cc.HostConfig.Devices {
			if !p.Validate(config.Username, "device", dev.PathOnHost, "") {
				return &types.AllowResult{
					Allow: false,
					Msg: map[string]string{
						"text":           fmt.Sprintf("Device %s is not allowed to be exported", dev.PathOnHost),
						"resource_type":  "device",
						"resource_value": dev.PathOnHost,
					},
				}
			}
		}
	}

	if cc.HostConfig.OomKillDisable != nil && *cc.HostConfig.OomKillDisable {
		if !p.Validate(config.Username, "config", "container_create_param_oom_kill_disable", "") {
			return &types.AllowResult{
				Allow: false,
				Msg: map[string]string{
					"text":           "--oom-kill-disable param is not allowed",
					"resource_type":  "config",
					"resource_value": "container_create_param_oom_kill_disable",
				},
			}
		}
	}

	if len(cc.HostConfig.Mounts) > 0 {
		for _, mount := range cc.HostConfig.Mounts {
			if mount.Type == "bind" {
				if len(mount.Source) > 0 {
					if !AllowVolume(mount.Source, config) {
						return &types.AllowResult{
							Allow: false,
							Msg: map[string]string{
								"text":           fmt.Sprintf("Volume %s is not allowed to be mounted", mount.Source),
								"resource_type":  "volume",
								"resource_value": mount.Source,
							},
						}
					}
				}
			}

			if mount.Type == "tmpfs" {
				if !p.Validate(config.Username, "config", "container_create_param_tmpfs", "") {
					return &types.AllowResult{
						Allow: false,
						Msg: map[string]string{
							"text":           "--tmpfs param is not allowed",
							"resource_type":  "config",
							"resource_value": "container_create_param_tmpfs",
						},
					}
				}
			}
		}
	}

	if len(cc.User) > 0 {
		if cc.Config.User == "root" && !p.Validate(config.Username, "config", "container_create_param_user_root", "") {
			return &types.AllowResult{
				Allow: false,
				Msg: map[string]string{
					"text":           "Running as user \"root\" is not allowed. Please use --user=\"someuser\" param.",
					"resource_type":  "config",
					"resource_value": "container_create_param_user_root",
				},
			}
		}
	}

	if !AllowImage(cc.Image, config) {
		return &types.AllowResult{
			Allow: false,
			Msg: map[string]string{
				"text":           fmt.Sprintf("Image %s is not allowed", cc.Image),
				"resource_type":  "image",
				"resource_value": cc.Image,
			},
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
	opts := strings.TrimSpace(jsonVO.String())

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
		opts = strings.TrimSpace(jsonVO.String())

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
