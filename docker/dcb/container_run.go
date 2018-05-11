package dcb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/pkg/cmdbuilder"
)

type ContainerCreateConfig struct {
	container.Config
	HostConfig       container.HostConfig
	NetworkingConfig network.NetworkingConfig
}

func ContainerCreate(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("container")
	cmd.Add("run")

	cc := &ContainerCreateConfig{}

	if req.RequestBody != nil {
		if err := json.NewDecoder(bytes.NewReader(req.RequestBody)).Decode(cc); err != nil {
			panic(err)
		}
	}

	if len(cc.HostConfig.ExtraHosts) > 0 {
		for _, eh := range cc.HostConfig.ExtraHosts {
			cmd.Add(fmt.Sprintf("--add-host %s", eh))
		}
	}

	if cc.AttachStdin {
		cmd.Add("-a stdin")
	}

	if cc.AttachStdout {
		cmd.Add("-a stdout")
	}

	if cc.AttachStderr {
		cmd.Add("-a stderr")
	}

	if cc.HostConfig.BlkioWeight > 0 {
		cmd.Add(fmt.Sprintf("--blkio-weight %d", cc.HostConfig.BlkioWeight))
	}

	if len(cc.HostConfig.BlkioWeightDevice) > 0 {
		cmd.Add(fmt.Sprintf("--blkio-weight-device %s", cc.HostConfig.BlkioWeightDevice))
	}

	if len(cc.HostConfig.CapAdd) > 0 {
		for _, ca := range cc.HostConfig.CapAdd {
			cmd.Add(fmt.Sprintf("--cap-add %s", ca))
		}
	}

	if len(cc.HostConfig.CapDrop) > 0 {
		for _, cd := range cc.HostConfig.CapDrop {
			cmd.Add(fmt.Sprintf("--cap-drop %s", cd))
		}
	}

	if len(cc.HostConfig.CgroupParent) > 0 {
		cmd.Add(fmt.Sprintf("--cgroup-parent %s", cc.HostConfig.CgroupParent))
	}

	// --cidfile

	if cc.HostConfig.CPUPeriod > 0 {
		cmd.Add(fmt.Sprintf("--cpu-period %d", cc.HostConfig.CPUPeriod))
	}

	if cc.HostConfig.CPUQuota > 0 {
		cmd.Add(fmt.Sprintf("--cpu-quota %d", cc.HostConfig.CPUQuota))
	}

	if cc.HostConfig.CPURealtimePeriod > 0 {
		cmd.Add(fmt.Sprintf("--cpu-rt-period %d", cc.HostConfig.CPURealtimePeriod))
	}

	if cc.HostConfig.CPURealtimeRuntime > 0 {
		cmd.Add(fmt.Sprintf("--cpu-rt-runtime %d", cc.HostConfig.CPURealtimeRuntime))
	}

	if cc.HostConfig.CPUShares > 0 {
		cmd.Add(fmt.Sprintf("--cpu-shares %d", cc.HostConfig.CPUShares))
	}

	if cc.HostConfig.NanoCPUs > 0 {
		cmd.Add(fmt.Sprintf("--cpus %d", cc.HostConfig.NanoCPUs))
	}

	if len(cc.HostConfig.CpusetCpus) > 0 {
		cmd.Add(fmt.Sprintf("--cpuset-cpus %s", cc.HostConfig.CpusetCpus))
	}

	if len(cc.HostConfig.CpusetMems) > 0 {
		cmd.Add(fmt.Sprintf("--cpuset-mems %s", cc.HostConfig.CpusetMems))
	}

	// --detach
	// --detach-keys

	if len(cc.HostConfig.Devices) > 0 {
		for _, d := range cc.HostConfig.Devices {
			cmd.Add(fmt.Sprintf("--device %s", d))
		}
	}

	if len(cc.HostConfig.BlkioDeviceReadBps) > 0 {
		for _, drb := range cc.HostConfig.BlkioDeviceReadBps {
			cmd.Add(fmt.Sprintf("--device-read-bps %s:%s", drb.Path, drb.Rate))
		}
	}

	if len(cc.HostConfig.BlkioDeviceWriteBps) > 0 {
		for _, dwb := range cc.HostConfig.BlkioDeviceWriteBps {
			cmd.Add(fmt.Sprintf("--device-write-bps %s:%s", dwb.Path, dwb.Rate))
		}
	}

	if len(cc.HostConfig.BlkioDeviceReadIOps) > 0 {
		for _, dri := range cc.HostConfig.BlkioDeviceReadIOps {
			cmd.Add(fmt.Sprintf("--device-read-iops %s:%s", dri.Path, dri.Rate))
		}
	}

	if len(cc.HostConfig.BlkioDeviceWriteIOps) > 0 {
		for _, dwi := range cc.HostConfig.BlkioDeviceReadIOps {
			cmd.Add(fmt.Sprintf("--device-write-iops %s:%s", dwi.Path, dwi.Rate))
		}
	}

	// --disable-content-trust

	if len(cc.HostConfig.DNS) > 0 {
		for _, d := range cc.HostConfig.DNS {
			cmd.Add(fmt.Sprintf("--dns %s", d))
		}
	}

	if len(cc.HostConfig.DNSOptions) > 0 {
		for _, do := range cc.HostConfig.DNSOptions {
			cmd.Add(fmt.Sprintf("--dns-opt %s", do))
		}
	}

	if len(cc.HostConfig.DNSSearch) > 0 {
		for _, ds := range cc.HostConfig.DNSSearch {
			cmd.Add(fmt.Sprintf("--dns-search %s", ds))
		}
	}

	if len(cc.Entrypoint) > 0 {
		cmd.Add(fmt.Sprintf("--entrypoint %s", cc.Entrypoint))
	}

	if len(cc.Env) > 0 {
		for _, e := range cc.Env {
			cmd.Add(fmt.Sprintf("--env \"%s\"", e))
		}
	}

	if len(cc.HostConfig.GroupAdd) > 0 {
		for _, ga := range cc.HostConfig.GroupAdd {
			cmd.Add(fmt.Sprintf("--group-add %s", ga))
		}
	}

	if cc.Healthcheck != nil {
		if len(cc.Healthcheck.Test) > 0 {
			for _, c := range cc.Healthcheck.Test {
				cmd.Add(fmt.Sprintf("--health-cmd %s", c))
			}
		}

		if cc.Healthcheck.Interval > 0 {
			cmd.Add(fmt.Sprintf("--health-interval %d", cc.Healthcheck.Interval))
		}

		if cc.Healthcheck.Retries > 0 {
			cmd.Add(fmt.Sprintf("--health-retries %d", cc.Healthcheck.Retries))
		}

		if cc.Healthcheck.StartPeriod > 0 {
			cmd.Add(fmt.Sprintf("--health-start-period %d", cc.Healthcheck.StartPeriod))
		}

		if cc.Healthcheck.Timeout > 0 {
			cmd.Add(fmt.Sprintf("--health-timeout %d", cc.Healthcheck.Timeout))
		}
	}

	if len(cc.Hostname) > 0 {
		cmd.Add(fmt.Sprintf("--hostname %s", cc.Hostname))
	}

	if *cc.HostConfig.Init {
		cmd.Add("--init")
	}

	if cc.OpenStdin {
		cmd.Add("--interactive")
	}

	// --ip
	// --ip6

	if len(cc.HostConfig.IpcMode) > 0 {
		cmd.Add(fmt.Sprintf("--ipc %s", cc.HostConfig.IpcMode))
	}

	if len(cc.HostConfig.Isolation) > 0 {
		cmd.Add(fmt.Sprintf("--isolation %s", cc.HostConfig.Isolation))
	}

	if cc.HostConfig.KernelMemory > 0 {
		cmd.Add(fmt.Sprintf("--kernel-memory %d", cc.HostConfig.KernelMemory))
	}

	if len(cc.Labels) > 0 {
		for k, v := range cc.Labels {
			cmd.Add(fmt.Sprintf("--label \"%s=%s\"", k, v))
		}
	}

	if len(cc.HostConfig.Links) > 0 {
		for _, l := range cc.HostConfig.Links {
			cmd.Add(fmt.Sprintf("--link %s", l))
		}
	}

	// --link-local-ip

	if len(cc.HostConfig.LogConfig.Type) > 0 {
		cmd.Add(fmt.Sprintf("--log-driver %s", cc.HostConfig.LogConfig.Type))
	}

	if len(cc.HostConfig.LogConfig.Config) > 0 {
		for k, v := range cc.HostConfig.LogConfig.Config {
			cmd.Add(fmt.Sprintf("--log-opt \"%s=%s\"", k, v))
		}
	}

	if len(cc.MacAddress) > 0 {
		cmd.Add(fmt.Sprintf("--mac-address %s", cc.MacAddress))
	}

	if cc.HostConfig.Memory > 0 {
		cmd.Add(fmt.Sprintf("--memory %d", cc.HostConfig.Memory))
	}

	if cc.HostConfig.MemoryReservation > 0 {
		cmd.Add(fmt.Sprintf("--memory-reservation %d", cc.HostConfig.MemoryReservation))
	}

	if cc.HostConfig.MemorySwap > 0 {
		cmd.Add(fmt.Sprintf("--memory-swap %d", cc.HostConfig.MemorySwap))
	}

	if cc.HostConfig.MemorySwappiness != nil {
		if *cc.HostConfig.MemorySwappiness > 0 {
			cmd.Add(fmt.Sprintf("--memory-swappiness %d", *cc.HostConfig.MemorySwappiness))
		}
	}

	cmd.GetParams(req.RequestURI)
	if len(cmd.Params) > 0 {
		cmd.GetParamAndAdd("name", "--name", false)
	}

	if len(cc.HostConfig.NetworkMode) > 0 {
		cmd.Add(fmt.Sprintf("--net %s", cc.HostConfig.NetworkMode))
	}

	// --net-alias
	// --no-healthcheck

	if cc.HostConfig.OomKillDisable != nil {
		if *cc.HostConfig.OomKillDisable {
			cmd.Add("--oom-kill-disable")
		}
	}

	if cc.HostConfig.OomScoreAdj > 0 {
		cmd.Add(fmt.Sprintf("--oom-score-adj %d", cc.HostConfig.OomScoreAdj))
	}

	if len(cc.HostConfig.PidMode) > 0 {
		cmd.Add(fmt.Sprintf("--pid %s", cc.HostConfig.PidMode))
	}

	if cc.HostConfig.PidsLimit > 0 {
		cmd.Add(fmt.Sprintf("--pids-limit %d", cc.HostConfig.PidsLimit))
	}

	if cc.HostConfig.Privileged {
		cmd.Add("--privileged")
	}

	if cc.HostConfig.PublishAllPorts {
		cmd.Add("--publish-all")
	}

	if cc.HostConfig.ReadonlyRootfs {
		cmd.Add("--read-only")
	}

	if len(cc.HostConfig.RestartPolicy.Name) > 0 {
		cmd.Add(fmt.Sprintf("--restart %s", cc.HostConfig.RestartPolicy.Name))
	}

	if len(cc.HostConfig.Runtime) > 0 {
		cmd.Add(fmt.Sprintf("--runtime %s", cc.HostConfig.Runtime))
	}

	if len(cc.HostConfig.SecurityOpt) > 0 {
		for _, so := range cc.HostConfig.SecurityOpt {
			cmd.Add(fmt.Sprintf("--security-opt %s", so))
		}
	}

	if cc.HostConfig.ShmSize > 0 {
		cmd.Add(fmt.Sprintf("--shm-size %d", cc.HostConfig.ShmSize))
	}

	// --sig-proxy

	if len(cc.StopSignal) > 0 {
		cmd.Add(fmt.Sprintf("--stop-signal %s", cc.StopSignal))
	}

	if *cc.StopTimeout > 0 {
		cmd.Add(fmt.Sprintf("--stop-timeout %d", cc.StopSignal))
	}

	if len(cc.HostConfig.Sysctls) > 0 {
		for k, v := range cc.HostConfig.Sysctls {
			cmd.Add(fmt.Sprintf("--sysctl %s=%s", k, v))
		}
	}

	if len(cc.HostConfig.Tmpfs) > 0 {
		for k, v := range cc.HostConfig.Tmpfs {
			cmd.Add(fmt.Sprintf("--tmpfs %s=%s", k, v))
		}
	}

	if cc.Tty {
		cmd.Add("--tty")
	}

	if len(cc.HostConfig.Ulimits) > 0 {
		for _, u := range cc.HostConfig.Ulimits {
			cmd.Add(fmt.Sprintf("--ulimit %s", u))
		}
	}

	if len(cc.User) > 0 {
		cmd.Add(fmt.Sprintf("--user %s", cc.User))
	}

	if len(cc.HostConfig.UsernsMode) > 0 {
		cmd.Add(fmt.Sprintf("--userns %s", cc.HostConfig.UsernsMode))
	}

	if len(cc.HostConfig.UTSMode) > 0 {
		cmd.Add(fmt.Sprintf("--uts %s", cc.HostConfig.UTSMode))
	}

	if len(cc.HostConfig.Binds) > 0 {
		for _, v := range cc.HostConfig.Binds {
			cmd.Add(fmt.Sprintf("--volume %s", v))
		}
	}

	if len(cc.HostConfig.VolumeDriver) > 0 {
		cmd.Add(fmt.Sprintf("--volume-driver %s", cc.HostConfig.VolumeDriver))
	}

	if len(cc.HostConfig.VolumesFrom) > 0 {
		for _, vf := range cc.HostConfig.VolumesFrom {
			cmd.Add(fmt.Sprintf("--volumes-from %s", vf))
		}
	}

	if len(cc.WorkingDir) > 0 {
		cmd.Add(fmt.Sprintf("--workdir %s", cc.WorkingDir))
	}

	if len(cc.Image) > 0 {
		cmd.Add(cc.Image)
	}

	if len(cc.Cmd) > 0 {
		cmd.Add(strings.Join(cc.Cmd, " "))
	}

	return cmd.String()
}
