package dcb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/pkg/cmdbuilder"
	"github.com/moby/moby/api/types/swarm"
)

func ServiceList(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("service")
	cmd.Add("ls")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		if _, ok := cmd.Params["filters"]; ok {
			var v map[string]map[string]bool

			err := json.Unmarshal([]byte(cmd.Params["filters"][0]), &v)
			if err != nil {
				panic(err)
			}

			for k, val := range v {
				for ka, _ := range val {
					cmd.Add(fmt.Sprintf("--filter \"%s=%s\"", k, ka))
				}
			}
		}

		if v, ok := cmd.Params["filter"]; ok {
			cmd.Add(v[0])
		}
	}

	return cmd.String()
}

func ServiceCreate(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("service")
	cmd.Add("create")

	svc := &swarm.Service{}

	if req.RequestBody != nil {
		if err := json.NewDecoder(bytes.NewReader(req.RequestBody)).Decode(svc); err != nil {
			panic(err)
		}
	}

	if svc.Spec.TaskTemplate.Placement != nil {
		if len(svc.Spec.TaskTemplate.Placement.Constraints) > 0 {
			for _, constraint := range svc.Spec.TaskTemplate.Placement.Constraints {
				cmd.Add(fmt.Sprintf("--constraint=%s", constraint))
			}
		}
	}

	if len(svc.Spec.Labels) > 0 {
		for k, v := range svc.Spec.Labels {
			cmd.Add(fmt.Sprintf("--label=\"%s=%s\"", k, v))
		}
	}

	if len(svc.Spec.TaskTemplate.ContainerSpec.Env) > 0 {
		for _, env := range svc.Spec.TaskTemplate.ContainerSpec.Env {
			cmd.Add(fmt.Sprintf("--env=%s", env))
		}
	}

	if len(svc.Spec.TaskTemplate.ContainerSpec.Labels) > 0 {
		for k, v := range svc.Spec.TaskTemplate.ContainerSpec.Labels {
			cmd.Add(fmt.Sprintf("--container-label=\"%s=%s\"", k, v))
		}
	}

	if svc.Spec.TaskTemplate.Resources != nil {
		if svc.Spec.TaskTemplate.Resources.Limits.NanoCPUs > 0 {
			cmd.Add(fmt.Sprintf("--limit-cpu=%s", svc.Spec.TaskTemplate.Resources.Limits.NanoCPUs))
		}

		if svc.Spec.TaskTemplate.Resources.Limits.MemoryBytes > 0 {
			cmd.Add(fmt.Sprintf("--limit-memory=%s", svc.Spec.TaskTemplate.Resources.Limits.MemoryBytes))
		}

		if svc.Spec.TaskTemplate.Resources.Reservations.NanoCPUs > 0 {
			cmd.Add(fmt.Sprintf("--reserve-cpu=%s", svc.Spec.TaskTemplate.Resources.Reservations.NanoCPUs))
		}

		if svc.Spec.TaskTemplate.Resources.Reservations.MemoryBytes > 0 {
			cmd.Add(fmt.Sprintf("--reserve-memory=%s", svc.Spec.TaskTemplate.Resources.Reservations.MemoryBytes))
		}
	}

	if svc.Spec.TaskTemplate.LogDriver != nil {
		if len(svc.Spec.TaskTemplate.LogDriver.Name) > 0 {
			cmd.Add(fmt.Sprintf("--log-driver=%s", svc.Spec.TaskTemplate.LogDriver.Name))
		}

		if len(svc.Spec.TaskTemplate.LogDriver.Options) > 0 {
			for k, v := range svc.Spec.TaskTemplate.LogDriver.Options {
				cmd.Add(fmt.Sprintf("--log-opt \"%s=%s\"", k, v))
			}
		}
	}

	if len(svc.Spec.TaskTemplate.ContainerSpec.Mounts) > 0 {
		for _, mount := range svc.Spec.TaskTemplate.ContainerSpec.Mounts {
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

			if len(mount.BindOptions.Propagation) > 0 {
				str = append(str, fmt.Sprintf("bind-propagation=%s", mount.BindOptions.Propagation))
			}

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
			if len(mount.VolumeOptions.DriverConfig.Options) > 0 {
				for k, v := range mount.VolumeOptions.DriverConfig.Options {
					str = append(str, fmt.Sprintf("volume-opt=\"%s=%s\"", k, v))
				}
			}

			cmd.Add(strings.Join(str, ","))
		}
	}

	if len(svc.ID) > 0 {
		cmd.Add(fmt.Sprintf("--name=%s", svc.ID))
	}

	if len(svc.Spec.Networks) > 0 {
		nt := []string{}

		for _, network := range svc.Spec.Networks {
			nt = append(nt, network.Target)
		}

		cmd.Add(strings.Join(nt, ","))
	}

	if svc.Spec.EndpointSpec != nil {
		if len(svc.Spec.EndpointSpec.Mode) > 0 {
			cmd.Add(fmt.Sprintf("--endpoint-mode=%s", svc.Spec.EndpointSpec.Mode))
		}

		if len(svc.Spec.EndpointSpec.Ports) > 0 {
			for _, port := range svc.Spec.EndpointSpec.Ports {
				pc := fmt.Sprintf("%s:%s", port.TargetPort, port.PublishedPort)
				if len(port.Protocol) > 0 {
					pc = fmt.Sprintf("%s/%s", port.Protocol, pc)
				}

				cmd.Add(fmt.Sprintf("--publish %s", pc))
			}
		}
	}

	if svc.Spec.Mode.Replicated != nil {
		if *svc.Spec.Mode.Replicated.Replicas > 0 {
			cmd.Add("--mode replicated")
		}

		if *svc.Spec.Mode.Replicated.Replicas > 0 {
			cmd.Add(fmt.Sprintf("--replicas=%s", svc.Spec.Mode.Replicated.Replicas))
		}
	}

	if svc.Spec.TaskTemplate.RestartPolicy != nil {
		if len(svc.Spec.TaskTemplate.RestartPolicy.Condition) > 0 {
			cmd.Add(fmt.Sprintf("--restart-condition=%s", svc.Spec.TaskTemplate.RestartPolicy.Condition))
		}

		if *svc.Spec.TaskTemplate.RestartPolicy.Delay > 0 {
			cmd.Add(fmt.Sprintf("--restart-delay=%s", svc.Spec.TaskTemplate.RestartPolicy.Delay))
		}

		if *svc.Spec.TaskTemplate.RestartPolicy.MaxAttempts > 0 {
			cmd.Add(fmt.Sprintf("--restart-max-attempts=%s", svc.Spec.TaskTemplate.RestartPolicy.MaxAttempts))
		}

		if *svc.Spec.TaskTemplate.RestartPolicy.Window > 0 {
			cmd.Add(fmt.Sprintf("--restart-window=%s", svc.Spec.TaskTemplate.RestartPolicy.Window))
		}
	}

	if svc.Spec.TaskTemplate.ContainerSpec.StopGracePeriod != nil {
		cmd.Add(fmt.Sprintf("--stop-grace-period=%s", svc.Spec.TaskTemplate.ContainerSpec.StopGracePeriod))
	}

	if svc.Spec.UpdateConfig != nil {
		if svc.Spec.UpdateConfig.Delay > 0 {
			cmd.Add(fmt.Sprintf("--update-delay=%s", svc.Spec.UpdateConfig.Delay))
		}

		if len(svc.Spec.UpdateConfig.FailureAction) > 0 {
			cmd.Add(fmt.Sprintf("--update-failure-action=%s", svc.Spec.UpdateConfig.FailureAction))
		}

		if svc.Spec.UpdateConfig.Parallelism > 0 {
			cmd.Add(fmt.Sprintf("--update-parallelism=%s", svc.Spec.UpdateConfig.Parallelism))
		}
	}

	if len(svc.Spec.TaskTemplate.ContainerSpec.User) > 0 {
		cmd.Add(fmt.Sprintf("--user=%s", svc.Spec.TaskTemplate.ContainerSpec.User))
	}

	if len(svc.Spec.TaskTemplate.ContainerSpec.Dir) > 0 {
		cmd.Add(fmt.Sprintf("--workdir=%s", svc.Spec.TaskTemplate.ContainerSpec.Dir))
	}

	if _, ok := req.RequestHeaders["X-Registry-Auth"]; ok {
		cmd.Add("--with-registry-auth=XXXXXXXXXXXX")
	}

	if len(svc.Spec.TaskTemplate.ContainerSpec.Image) > 0 {
		cmd.Add(svc.Spec.TaskTemplate.ContainerSpec.Image)
	}

	if len(svc.Spec.TaskTemplate.ContainerSpec.Command) > 0 {
		for _, command := range svc.Spec.TaskTemplate.ContainerSpec.Command {
			cmd.Add(command)
		}
	}

	if len(svc.Spec.TaskTemplate.ContainerSpec.Args) > 0 {
		for _, arg := range svc.Spec.TaskTemplate.ContainerSpec.Args {
			cmd.Add(arg)
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

	svc := &swarm.Service{}

	if req.RequestBody != nil {
		if err := json.NewDecoder(bytes.NewReader(req.RequestBody)).Decode(svc); err != nil {
			panic(err)
		}
	}

	if len(svc.Spec.TaskTemplate.Placement.Constraints) > 0 {
		for _, constraint := range svc.Spec.TaskTemplate.Placement.Constraints {
			cmd.Add(fmt.Sprintf("--constraint=%s", constraint))
		}
	}

	if len(svc.Spec.Labels) > 0 {
		for k, v := range svc.Spec.Labels {
			cmd.Add(fmt.Sprintf("--label=\"%s=%s\"", k, v))
		}
	}

	if len(svc.Spec.TaskTemplate.ContainerSpec.Env) > 0 {
		for _, env := range svc.Spec.TaskTemplate.ContainerSpec.Env {
			cmd.Add(fmt.Sprintf("--env=%s", env))
		}
	}

	if len(svc.Spec.TaskTemplate.ContainerSpec.Labels) > 0 {
		for k, v := range svc.Spec.TaskTemplate.ContainerSpec.Labels {
			cmd.Add(fmt.Sprintf("--container-label=\"%s=%s\"", k, v))
		}
	}

	if svc.Spec.TaskTemplate.Resources.Limits.NanoCPUs > 0 {
		cmd.Add(fmt.Sprintf("--limit-cpu=%s", svc.Spec.TaskTemplate.Resources.Limits.NanoCPUs))
	}

	if svc.Spec.TaskTemplate.Resources.Limits.MemoryBytes > 0 {
		cmd.Add(fmt.Sprintf("--limit-memory=%s", svc.Spec.TaskTemplate.Resources.Limits.MemoryBytes))
	}

	if svc.Spec.TaskTemplate.Resources.Reservations.NanoCPUs > 0 {
		cmd.Add(fmt.Sprintf("--reserve-cpu=%s", svc.Spec.TaskTemplate.Resources.Reservations.NanoCPUs))
	}

	if svc.Spec.TaskTemplate.Resources.Reservations.MemoryBytes > 0 {
		cmd.Add(fmt.Sprintf("--reserve-memory=%s", svc.Spec.TaskTemplate.Resources.Reservations.MemoryBytes))
	}

	if len(svc.Spec.TaskTemplate.LogDriver.Name) > 0 {
		cmd.Add(fmt.Sprintf("--log-driver=%s", svc.Spec.TaskTemplate.LogDriver.Name))
	}

	if len(svc.Spec.TaskTemplate.LogDriver.Options) > 0 {
		for k, v := range svc.Spec.TaskTemplate.LogDriver.Options {
			cmd.Add(fmt.Sprintf("--log-opt \"%s=%s\"", k, v))
		}
	}

	if len(svc.Spec.TaskTemplate.ContainerSpec.Mounts) > 0 {
		for _, mount := range svc.Spec.TaskTemplate.ContainerSpec.Mounts {
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

			if len(mount.BindOptions.Propagation) > 0 {
				str = append(str, fmt.Sprintf("bind-propagation=%s", mount.BindOptions.Propagation))
			}

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
			if len(mount.VolumeOptions.DriverConfig.Options) > 0 {
				for k, v := range mount.VolumeOptions.DriverConfig.Options {
					str = append(str, fmt.Sprintf("volume-opt=\"%s=%s\"", k, v))
				}
			}

			cmd.Add(strings.Join(str, ","))
		}
	}

	if len(svc.ID) > 0 {
		cmd.Add(fmt.Sprintf("--name=%s", svc.ID))
	}

	if len(svc.Spec.EndpointSpec.Mode) > 0 {
		cmd.Add(fmt.Sprintf("--endpoint-mode=%s", svc.Spec.EndpointSpec.Mode))
	}

	if len(svc.Spec.EndpointSpec.Ports) > 0 {
		for _, port := range svc.Spec.EndpointSpec.Ports {
			pc := fmt.Sprintf("%s:%s", port.TargetPort, port.PublishedPort)
			if len(port.Protocol) > 0 {
				pc = fmt.Sprintf("%s/%s", port.Protocol, pc)
			}

			cmd.Add(fmt.Sprintf("--publish %s", pc))
		}
	}

	if *svc.Spec.Mode.Replicated.Replicas > 0 {
		cmd.Add(fmt.Sprintf("--replicas=%s", svc.Spec.Mode.Replicated.Replicas))
	}

	if len(svc.Spec.TaskTemplate.RestartPolicy.Condition) > 0 {
		cmd.Add(fmt.Sprintf("--restart-condition=%s", svc.Spec.TaskTemplate.RestartPolicy.Condition))
	}

	if *svc.Spec.TaskTemplate.RestartPolicy.Delay > 0 {
		cmd.Add(fmt.Sprintf("--restart-delay=%s", svc.Spec.TaskTemplate.RestartPolicy.Delay))
	}

	if *svc.Spec.TaskTemplate.RestartPolicy.MaxAttempts > 0 {
		cmd.Add(fmt.Sprintf("--restart-max-attempts=%s", svc.Spec.TaskTemplate.RestartPolicy.MaxAttempts))
	}

	if *svc.Spec.TaskTemplate.RestartPolicy.Window > 0 {
		cmd.Add(fmt.Sprintf("--restart-window=%s", svc.Spec.TaskTemplate.RestartPolicy.Window))
	}

	if *svc.Spec.TaskTemplate.ContainerSpec.StopGracePeriod > 0 {
		cmd.Add(fmt.Sprintf("--stop-grace-period=%s", svc.Spec.TaskTemplate.ContainerSpec.StopGracePeriod))
	}

	if svc.Spec.UpdateConfig.Delay > 0 {
		cmd.Add(fmt.Sprintf("--update-delay=%s", svc.Spec.UpdateConfig.Delay))
	}

	if len(svc.Spec.UpdateConfig.FailureAction) > 0 {
		cmd.Add(fmt.Sprintf("--update-failure-action=%s", svc.Spec.UpdateConfig.FailureAction))
	}

	if svc.Spec.UpdateConfig.Parallelism > 0 {
		cmd.Add(fmt.Sprintf("--update-parallelism=%s", svc.Spec.UpdateConfig.Parallelism))
	}

	if len(svc.Spec.TaskTemplate.ContainerSpec.User) > 0 {
		cmd.Add(fmt.Sprintf("--user=%s", svc.Spec.TaskTemplate.ContainerSpec.User))
	}

	if len(svc.Spec.TaskTemplate.ContainerSpec.Dir) > 0 {
		cmd.Add(fmt.Sprintf("--workdir=%s", svc.Spec.TaskTemplate.ContainerSpec.Dir))
	}

	if _, ok := req.RequestHeaders["X-Registry-Auth"]; ok {
		cmd.Add("--with-registry-auth=XXXXXXXXXXXX")
	}

	if len(svc.Spec.TaskTemplate.ContainerSpec.Command) > 0 {
		for _, command := range svc.Spec.TaskTemplate.ContainerSpec.Command {
			cmd.Add(command)
		}
	}

	if len(svc.Spec.TaskTemplate.ContainerSpec.Args) > 0 {
		for _, arg := range svc.Spec.TaskTemplate.ContainerSpec.Args {
			cmd.Add(arg)
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
