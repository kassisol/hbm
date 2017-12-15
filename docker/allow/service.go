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
	"github.com/kassisol/hbm/storage"
	"github.com/kassisol/hbm/version"
)

func AllowServiceCreate(req authorization.Request, config *types.Config) *types.AllowResult {
	svc := &swarm.Service{}

	err := json.Decode(req.RequestBody, svc)
	if err != nil {
		return &types.AllowResult{Allow: false, Error: err.Error()}
	}

	defer utils.RecoverFunc()

	l, _ := log.NewDriver("standard", nil)

	s, err := storage.NewDriver("sqlite", config.AppPath)
	if err != nil {
		l.WithFields(driver.Fields{
			"storagedriver": "sqlite",
			"logdriver":     "standard",
			"version":       version.Version,
		}).Fatal(err)
	}
	defer s.End()

	if svc.Spec.EndpointSpec != nil {
		if len(svc.Spec.EndpointSpec.Ports) > 0 {
			for _, port := range svc.Spec.EndpointSpec.Ports {
				if !s.ValidatePolicy(config.Username, "port", string(port.PublishedPort), "") {
					return &types.AllowResult{Allow: false, Msg: fmt.Sprintf("Port %s is not allowed to be pubished", port.PublishedPort)}
				}
			}
		}
	}

	if len(svc.Spec.TaskTemplate.ContainerSpec.Mounts) > 0 {
		for _, mount := range svc.Spec.TaskTemplate.ContainerSpec.Mounts {
			if mount.Type == "bind" {
				if len(mount.Source) > 0 {
					if !AllowVolume(mount.Source, config) {
						return &types.AllowResult{Allow: false, Msg: fmt.Sprintf("Volume %s is not allowed to be mounted", mount.Source)}
					}
				}
			}
		}
	}

	if svc.Spec.TaskTemplate.LogDriver != nil {
		if len(svc.Spec.TaskTemplate.LogDriver.Name) > 0 {
			if !s.ValidatePolicy(config.Username, "logdriver", svc.Spec.TaskTemplate.LogDriver.Name, "") {
				return &types.AllowResult{Allow: false, Msg: fmt.Sprintf("Log driver %s is not allowed", svc.Spec.TaskTemplate.LogDriver.Name)}
			}
		}

		if len(svc.Spec.TaskTemplate.LogDriver.Options) > 0 {
			for k, v := range svc.Spec.TaskTemplate.LogDriver.Options {
				los := fmt.Sprintf("%s=%s", k, v)

				if !s.ValidatePolicy(config.Username, "logopt", los, "") {
					return &types.AllowResult{Allow: false, Msg: fmt.Sprintf("Log driver %s is not allowed", los)}
				}
			}
		}
	}

	if len(svc.Spec.TaskTemplate.ContainerSpec.User) > 0 {
		if svc.Spec.TaskTemplate.ContainerSpec.User == "root" && !s.ValidatePolicy(config.Username, "config", "container_create_user_root", "") {
			return &types.AllowResult{Allow: false, Msg: "Running as user \"root\" is not allowed. Please use --user=\"someuser\" param."}
		}
	}

	return &types.AllowResult{Allow: true}
}
