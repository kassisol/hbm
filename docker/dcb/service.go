package dcb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/pkg/cmdbuilder"
)

func ServiceList(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("service")
	cmd.Add("ls")

	cmd.GetParams(req.RequestURI)

	cmd.AddFilters()

	return cmd.String()
}

func ServiceCreate(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("service")
	cmd.Add("create")

	svc := &swarm.ServiceSpec{}

	if req.RequestBody != nil {
		if err := json.NewDecoder(bytes.NewReader(req.RequestBody)).Decode(svc); err != nil {
			panic(err)
		}
	}

	if svc.TaskTemplate.Placement != nil {
		if len(svc.TaskTemplate.Placement.Constraints) > 0 {
			for _, constraint := range svc.TaskTemplate.Placement.Constraints {
				cmd.Add(fmt.Sprintf("--constraint %s", constraint))
			}
		}
	}

	if len(svc.Annotations.Labels) > 0 {
		for k, v := range svc.Annotations.Labels {
			cmd.Add(fmt.Sprintf("--label \"%s=%s\"", k, v))
		}
	}

	if svc.TaskTemplate.ContainerSpec != nil {
		if len(svc.TaskTemplate.ContainerSpec.Env) > 0 {
			for _, env := range svc.TaskTemplate.ContainerSpec.Env {
				cmd.Add(fmt.Sprintf("--env %s", env))
			}
		}

		if len(svc.TaskTemplate.ContainerSpec.Labels) > 0 {
			for k, v := range svc.TaskTemplate.ContainerSpec.Labels {
				cmd.Add(fmt.Sprintf("--container-label \"%s=%s\"", k, v))
			}
		}
	}

	if svc.TaskTemplate.Resources != nil {
		if svc.TaskTemplate.Resources.Limits != nil {
			if svc.TaskTemplate.Resources.Limits.NanoCPUs > 0 {
				cmd.Add(fmt.Sprintf("--limit-cpu %s", svc.TaskTemplate.Resources.Limits.NanoCPUs))
			}

			if svc.TaskTemplate.Resources.Limits.MemoryBytes > 0 {
				cmd.Add(fmt.Sprintf("--limit-memory %s", svc.TaskTemplate.Resources.Limits.MemoryBytes))
			}
		}

		if svc.TaskTemplate.Resources.Reservations != nil {
			if svc.TaskTemplate.Resources.Reservations.NanoCPUs > 0 {
				cmd.Add(fmt.Sprintf("--reserve-cpu %s", svc.TaskTemplate.Resources.Reservations.NanoCPUs))
			}

			if svc.TaskTemplate.Resources.Reservations.MemoryBytes > 0 {
				cmd.Add(fmt.Sprintf("--reserve-memory %s", svc.TaskTemplate.Resources.Reservations.MemoryBytes))
			}
		}
	}

	if svc.TaskTemplate.LogDriver != nil {
		if len(svc.TaskTemplate.LogDriver.Name) > 0 {
			cmd.Add(fmt.Sprintf("--log-driver %s", svc.TaskTemplate.LogDriver.Name))
		}

		if len(svc.TaskTemplate.LogDriver.Options) > 0 {
			for k, v := range svc.TaskTemplate.LogDriver.Options {
				cmd.Add(fmt.Sprintf("--log-opt \"%s=%s\"", k, v))
			}
		}
	}

	if svc.TaskTemplate.ContainerSpec != nil {
		if len(svc.TaskTemplate.ContainerSpec.Mounts) > 0 {
			for _, mount := range svc.TaskTemplate.ContainerSpec.Mounts {
				str := []string{}

				if mount.Type == "bind" {
					str = append(str, "type=bind")
				}

				if len(mount.Source) > 0 {
					str = append(str, fmt.Sprintf("source=%s", mount.Source))
				}

				if len(mount.Target) > 0 {
					if mount.ReadOnly {
						str = append(str, fmt.Sprintf("destination=%s:ro", mount.Target))
					} else {
						str = append(str, fmt.Sprintf("destination=%s", mount.Target))
					}
				}

				if mount.BindOptions != nil {
					if len(mount.BindOptions.Propagation) > 0 {
						str = append(str, fmt.Sprintf("bind-propagation=%s", mount.BindOptions.Propagation))
					}
				}

				if mount.VolumeOptions != nil {
					if !mount.VolumeOptions.NoCopy {
						str = append(str, "volume-nocopy=false")
					}

					if len(mount.VolumeOptions.Labels) > 0 {
						for k, v := range mount.VolumeOptions.Labels {
							str = append(str, fmt.Sprintf("volume-label=\"%s=%s\"", k, v))
						}
					}

					if len(mount.VolumeOptions.DriverConfig.Name) > 0 {
						str = append(str, fmt.Sprintf("volume-driver=%s", mount.VolumeOptions.DriverConfig.Name))
					}
					if mount.VolumeOptions.DriverConfig != nil {
						if len(mount.VolumeOptions.DriverConfig.Options) > 0 {
							for k, v := range mount.VolumeOptions.DriverConfig.Options {
								str = append(str, fmt.Sprintf("volume-opt=\"%s=%s\"", k, v))
							}
						}
					}

					cmd.Add(strings.Join(str, ","))
				}
			}
		}
	}

	if len(svc.Annotations.Name) > 0 {
		cmd.Add(fmt.Sprintf("--name %s", svc.Annotations.Name))
	}

	if len(svc.Networks) > 0 {
		nt := []string{}

		for _, network := range svc.Networks {
			nt = append(nt, network.Target)
		}

		cmd.Add(strings.Join(nt, ","))
	}

	if svc.EndpointSpec != nil {
		if len(svc.EndpointSpec.Mode) > 0 {
			cmd.Add(fmt.Sprintf("--endpoint-mode %s", svc.EndpointSpec.Mode))
		}

		if len(svc.EndpointSpec.Ports) > 0 {
			for _, port := range svc.EndpointSpec.Ports {
				pc := fmt.Sprintf("%s:%s", port.TargetPort, port.PublishedPort)
				if len(port.Protocol) > 0 {
					pc = fmt.Sprintf("%s/%s", port.Protocol, pc)
				}

				cmd.Add(fmt.Sprintf("--publish %s", pc))
			}
		}
	}

	if svc.Mode.Replicated != nil {
		/*
			if svc.Mode.Replicated.Replicas != nil {
				cmd.Add("--mode replicated")
			}

		*/
		if svc.Mode.Replicated.Replicas != nil {
			cmd.Add(fmt.Sprintf("--replicas %d", *svc.Mode.Replicated.Replicas))
		}
	}

	if svc.Mode.Global != nil {
		cmd.Add("--mode global")
	}

	if svc.TaskTemplate.RestartPolicy != nil {
		if len(svc.TaskTemplate.RestartPolicy.Condition) > 0 {
			cmd.Add(fmt.Sprintf("--restart-condition %s", svc.TaskTemplate.RestartPolicy.Condition))
		}

		if svc.TaskTemplate.RestartPolicy.Delay != nil {
			if *svc.TaskTemplate.RestartPolicy.Delay > 0 {
				cmd.Add(fmt.Sprintf("--restart-delay %s", svc.TaskTemplate.RestartPolicy.Delay))
			}
		}

		if svc.TaskTemplate.RestartPolicy.MaxAttempts != nil {
			if *svc.TaskTemplate.RestartPolicy.MaxAttempts > 0 {
				cmd.Add(fmt.Sprintf("--restart-max-attempts %s", svc.TaskTemplate.RestartPolicy.MaxAttempts))
			}
		}

		if svc.TaskTemplate.RestartPolicy.Window != nil {
			if *svc.TaskTemplate.RestartPolicy.Window > 0 {
				cmd.Add(fmt.Sprintf("--restart-window %s", svc.TaskTemplate.RestartPolicy.Window))
			}
		}
	}

	if svc.TaskTemplate.ContainerSpec != nil {
		if svc.TaskTemplate.ContainerSpec.StopGracePeriod != nil {
			cmd.Add(fmt.Sprintf("--stop-grace-period %s", svc.TaskTemplate.ContainerSpec.StopGracePeriod))
		}
	}

	if svc.UpdateConfig != nil {
		if svc.UpdateConfig.Delay > 0 {
			cmd.Add(fmt.Sprintf("--update-delay %s", svc.UpdateConfig.Delay))
		}

		if len(svc.UpdateConfig.FailureAction) > 0 {
			cmd.Add(fmt.Sprintf("--update-failure-action %s", svc.UpdateConfig.FailureAction))
		}

		if svc.UpdateConfig.Parallelism > 0 {
			cmd.Add(fmt.Sprintf("--update-parallelism %s", svc.UpdateConfig.Parallelism))
		}
	}

	if svc.TaskTemplate.ContainerSpec != nil {
		if len(svc.TaskTemplate.ContainerSpec.User) > 0 {
			cmd.Add(fmt.Sprintf("--user %s", svc.TaskTemplate.ContainerSpec.User))
		}

		if len(svc.TaskTemplate.ContainerSpec.Dir) > 0 {
			cmd.Add(fmt.Sprintf("--workdir %s", svc.TaskTemplate.ContainerSpec.Dir))
		}
	}

	if _, ok := req.RequestHeaders["X-Registry-Auth"]; ok {
		cmd.Add("--with-registry-auth")
	}

	if svc.TaskTemplate.ContainerSpec != nil {
		if len(svc.TaskTemplate.ContainerSpec.Image) > 0 {
			cmd.Add(svc.TaskTemplate.ContainerSpec.Image)
		}

		if len(svc.TaskTemplate.ContainerSpec.Command) > 0 {
			for _, command := range svc.TaskTemplate.ContainerSpec.Command {
				cmd.Add(command)
			}
		}

		if len(svc.TaskTemplate.ContainerSpec.Args) > 0 {
			for _, arg := range svc.TaskTemplate.ContainerSpec.Args {
				cmd.Add(arg)
			}
		}
	}

	return cmd.String()
}

func ServiceRemove(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("service")
	cmd.Add("rm")

	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func ServiceInspect(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("service")
	cmd.Add("inspect")

	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func ServiceUpdate(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("service")
	cmd.Add("update")

	svc := &swarm.ServiceSpec{}

	if req.RequestBody != nil {
		if err := json.NewDecoder(bytes.NewReader(req.RequestBody)).Decode(svc); err != nil {
			panic(err)
		}
	}

	if svc.TaskTemplate.Placement != nil {
		if len(svc.TaskTemplate.Placement.Constraints) > 0 {
			for _, constraint := range svc.TaskTemplate.Placement.Constraints {
				cmd.Add(fmt.Sprintf("--constraint %s", constraint))
			}
		}
	}

	if len(svc.Annotations.Labels) > 0 {
		for k, v := range svc.Annotations.Labels {
			cmd.Add(fmt.Sprintf("--label \"%s=%s\"", k, v))
		}
	}

	if svc.TaskTemplate.ContainerSpec != nil {
		if len(svc.TaskTemplate.ContainerSpec.Env) > 0 {
			for _, env := range svc.TaskTemplate.ContainerSpec.Env {
				cmd.Add(fmt.Sprintf("--env %s", env))
			}
		}

		if len(svc.TaskTemplate.ContainerSpec.Labels) > 0 {
			for k, v := range svc.TaskTemplate.ContainerSpec.Labels {
				cmd.Add(fmt.Sprintf("--container-label \"%s=%s\"", k, v))
			}
		}
	}

	if svc.TaskTemplate.Resources != nil {
		if svc.TaskTemplate.Resources.Limits != nil {
			if svc.TaskTemplate.Resources.Limits.NanoCPUs > 0 {
				cmd.Add(fmt.Sprintf("--limit-cpu %s", svc.TaskTemplate.Resources.Limits.NanoCPUs))
			}

			if svc.TaskTemplate.Resources.Limits.MemoryBytes > 0 {
				cmd.Add(fmt.Sprintf("--limit-memory %s", svc.TaskTemplate.Resources.Limits.MemoryBytes))
			}
		}

		if svc.TaskTemplate.Resources.Reservations != nil {
			if svc.TaskTemplate.Resources.Reservations.NanoCPUs > 0 {
				cmd.Add(fmt.Sprintf("--reserve-cpu %s", svc.TaskTemplate.Resources.Reservations.NanoCPUs))
			}

			if svc.TaskTemplate.Resources.Reservations.MemoryBytes > 0 {
				cmd.Add(fmt.Sprintf("--reserve-memory %s", svc.TaskTemplate.Resources.Reservations.MemoryBytes))
			}
		}
	}

	if svc.TaskTemplate.LogDriver != nil {
		if len(svc.TaskTemplate.LogDriver.Name) > 0 {
			cmd.Add(fmt.Sprintf("--log-driver %s", svc.TaskTemplate.LogDriver.Name))
		}

		if len(svc.TaskTemplate.LogDriver.Options) > 0 {
			for k, v := range svc.TaskTemplate.LogDriver.Options {
				cmd.Add(fmt.Sprintf("--log-opt \"%s=%s\"", k, v))
			}
		}
	}

	if svc.TaskTemplate.ContainerSpec != nil {
		if len(svc.TaskTemplate.ContainerSpec.Mounts) > 0 {
			for _, mount := range svc.TaskTemplate.ContainerSpec.Mounts {
				str := []string{}

				if mount.Type == "bind" {
					str = append(str, "type=bind")
				}

				if len(mount.Source) > 0 {
					str = append(str, fmt.Sprintf("source=%s", mount.Source))
				}

				if len(mount.Target) > 0 {
					if mount.ReadOnly {
						str = append(str, fmt.Sprintf("destination=%s:ro", mount.Target))
					} else {
						str = append(str, fmt.Sprintf("destination=%s", mount.Target))
					}
				}

				if mount.BindOptions != nil {
					if len(mount.BindOptions.Propagation) > 0 {
						str = append(str, fmt.Sprintf("bind-propagation=%s", mount.BindOptions.Propagation))
					}
				}

				if mount.VolumeOptions != nil {
					if !mount.VolumeOptions.NoCopy {
						str = append(str, "volume-nocopy=false")
					}

					if len(mount.VolumeOptions.Labels) > 0 {
						for k, v := range mount.VolumeOptions.Labels {
							str = append(str, fmt.Sprintf("volume-label=\"%s=%s\"", k, v))
						}
					}

					if len(mount.VolumeOptions.DriverConfig.Name) > 0 {
						str = append(str, fmt.Sprintf("volume-driver=%s", mount.VolumeOptions.DriverConfig.Name))
					}
					if mount.VolumeOptions.DriverConfig != nil {
						if len(mount.VolumeOptions.DriverConfig.Options) > 0 {
							for k, v := range mount.VolumeOptions.DriverConfig.Options {
								str = append(str, fmt.Sprintf("volume-opt=\"%s=%s\"", k, v))
							}
						}
					}

					cmd.Add(strings.Join(str, ","))
				}
			}
		}
	}

	if len(svc.Annotations.Name) > 0 {
		cmd.Add(fmt.Sprintf("--name %s", svc.Annotations.Name))
	}

	if svc.EndpointSpec != nil {
		if len(svc.EndpointSpec.Mode) > 0 {
			cmd.Add(fmt.Sprintf("--endpoint-mode %s", svc.EndpointSpec.Mode))
		}

		if len(svc.EndpointSpec.Ports) > 0 {
			for _, port := range svc.EndpointSpec.Ports {
				pc := fmt.Sprintf("%s:%s", port.TargetPort, port.PublishedPort)
				if len(port.Protocol) > 0 {
					pc = fmt.Sprintf("%s/%s", port.Protocol, pc)
				}

				cmd.Add(fmt.Sprintf("--publish %s", pc))
			}
		}
	}

	if svc.Mode.Replicated != nil {
		if svc.Mode.Replicated.Replicas != nil {
			cmd.Add(fmt.Sprintf("--replicas %d", *svc.Mode.Replicated.Replicas))
		}
	}

	if svc.TaskTemplate.RestartPolicy != nil {
		if len(svc.TaskTemplate.RestartPolicy.Condition) > 0 {
			cmd.Add(fmt.Sprintf("--restart-condition %s", svc.TaskTemplate.RestartPolicy.Condition))
		}

		if svc.TaskTemplate.RestartPolicy.Delay != nil {
			if *svc.TaskTemplate.RestartPolicy.Delay > 0 {
				cmd.Add(fmt.Sprintf("--restart-delay %s", svc.TaskTemplate.RestartPolicy.Delay))
			}
		}

		if svc.TaskTemplate.RestartPolicy.MaxAttempts != nil {
			if *svc.TaskTemplate.RestartPolicy.MaxAttempts > 0 {
				cmd.Add(fmt.Sprintf("--restart-max-attempts %s", svc.TaskTemplate.RestartPolicy.MaxAttempts))
			}
		}

		if svc.TaskTemplate.RestartPolicy.Window != nil {
			if *svc.TaskTemplate.RestartPolicy.Window > 0 {
				cmd.Add(fmt.Sprintf("--restart-window %s", svc.TaskTemplate.RestartPolicy.Window))
			}
		}
	}

	if svc.TaskTemplate.ContainerSpec != nil {
		if svc.TaskTemplate.ContainerSpec.StopGracePeriod != nil {
			cmd.Add(fmt.Sprintf("--stop-grace-period %s", svc.TaskTemplate.ContainerSpec.StopGracePeriod))
		}
	}

	if svc.UpdateConfig != nil {
		if svc.UpdateConfig.Delay > 0 {
			cmd.Add(fmt.Sprintf("--update-delay %s", svc.UpdateConfig.Delay))
		}

		if len(svc.UpdateConfig.FailureAction) > 0 {
			cmd.Add(fmt.Sprintf("--update-failure-action %s", svc.UpdateConfig.FailureAction))
		}

		if svc.UpdateConfig.Parallelism > 0 {
			cmd.Add(fmt.Sprintf("--update-parallelism %s", svc.UpdateConfig.Parallelism))
		}
	}

	if svc.TaskTemplate.ContainerSpec != nil {
		if len(svc.TaskTemplate.ContainerSpec.User) > 0 {
			cmd.Add(fmt.Sprintf("--user %s", svc.TaskTemplate.ContainerSpec.User))
		}

		if len(svc.TaskTemplate.ContainerSpec.Dir) > 0 {
			cmd.Add(fmt.Sprintf("--workdir %s", svc.TaskTemplate.ContainerSpec.Dir))
		}
	}

	if _, ok := req.RequestHeaders["X-Registry-Auth"]; ok {
		cmd.Add("--with-registry-auth")
	}

	if svc.TaskTemplate.ContainerSpec != nil {
		if len(svc.TaskTemplate.ContainerSpec.Command) > 0 {
			for _, command := range svc.TaskTemplate.ContainerSpec.Command {
				cmd.Add(command)
			}
		}

		if len(svc.TaskTemplate.ContainerSpec.Args) > 0 {
			for _, arg := range svc.TaskTemplate.ContainerSpec.Args {
				cmd.Add(arg)
			}
		}
	}

	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func ServiceLogs(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("service")
	cmd.Add("logs")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		cmd.GetParamAndAdd("details", "--details", true)
		cmd.GetParamAndAdd("follow", "--follow", true)
		cmd.GetParamAndAdd("stdout", "--stdout", true)
		cmd.GetParamAndAdd("stderr", "--stderr", true)
		cmd.GetParamAndAdd("since", "--since", false)
		cmd.GetParamAndAdd("timestamps", "--timestamps", true)
		cmd.GetParamAndAdd("tail", "--tail", false)
	}

	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}
