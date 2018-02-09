package allow

import (
	"fmt"

	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/juliengk/go-log"
	"github.com/juliengk/go-log/driver"
	"github.com/juliengk/go-utils"
	"github.com/juliengk/go-utils/json"
	"github.com/kassisol/hbm/docker/allow/types"
	policyobj "github.com/kassisol/hbm/object/policy"
	"github.com/kassisol/hbm/version"
)

func ServiceCreate(req authorization.Request, config *types.Config) *types.AllowResult {
	svc := &swarm.ServiceSpec{}

	err := json.Decode(req.RequestBody, svc)
	if err != nil {
		return &types.AllowResult{Allow: false, Error: err.Error()}
	}

	defer utils.RecoverFunc()

	l, _ := log.NewDriver("standard", nil)

	p, err := policyobj.New("sqlite", config.AppPath)
	if err != nil {
		l.WithFields(driver.Fields{
			"storagedriver": "sqlite",
			"logdriver":     "standard",
			"version":       version.Version,
		}).Fatal(err)
	}
	defer p.End()

	if svc.EndpointSpec != nil {
		if len(svc.EndpointSpec.Ports) > 0 {
			for _, port := range svc.EndpointSpec.Ports {
				if !p.Validate(config.Username, "port", string(port.PublishedPort), "") {
					return &types.AllowResult{Allow: false, Msg: fmt.Sprintf("Port %s is not allowed to be published", port.PublishedPort)}
				}
			}
		}
	}

	if svc.TaskTemplate.ContainerSpec != nil {
		if len(svc.TaskTemplate.ContainerSpec.Mounts) > 0 {
			for _, mount := range svc.TaskTemplate.ContainerSpec.Mounts {
				if mount.Type == "bind" {
					if len(mount.Source) > 0 {
						if !AllowVolume(mount.Source, config) {
							return &types.AllowResult{Allow: false, Msg: fmt.Sprintf("Volume %s is not allowed to be mounted", mount.Source)}
						}
					}
				}
			}
		}

		if len(svc.TaskTemplate.ContainerSpec.User) > 0 {
			if svc.TaskTemplate.ContainerSpec.User == "root" && !p.Validate(config.Username, "config", "container_create_user_root", "") {
				return &types.AllowResult{Allow: false, Msg: "Running as user \"root\" is not allowed. Please use --user=\"someuser\" param."}
			}
		}
	}

	if svc.TaskTemplate.LogDriver != nil {
		if len(svc.TaskTemplate.LogDriver.Name) > 0 {
			if !p.Validate(config.Username, "logdriver", svc.TaskTemplate.LogDriver.Name, "") {
				return &types.AllowResult{Allow: false, Msg: fmt.Sprintf("Log driver %s is not allowed", svc.TaskTemplate.LogDriver.Name)}
			}
		}

		if len(svc.TaskTemplate.LogDriver.Options) > 0 {
			for k, v := range svc.TaskTemplate.LogDriver.Options {
				los := fmt.Sprintf("%s=%s", k, v)

				if !p.Validate(config.Username, "logopt", los, "") {
					return &types.AllowResult{Allow: false, Msg: fmt.Sprintf("Log driver %s is not allowed", los)}
				}
			}
		}
	}

	return &types.AllowResult{Allow: true}
}
