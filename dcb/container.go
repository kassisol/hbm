package dcb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/container"
	"github.com/docker/engine-api/types/network"
	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/pkg/cmdbuilder"
	"github.com/kassisol/hbm/pkg/utils"
)

type ContainerCreateConfig struct {
	container.Config
	HostConfig       container.HostConfig
	NetworkingConfig network.NetworkingConfig
}

func ContainerList(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("ps")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		cmd.GetParamAndAdd("all", "-a", true)
		cmd.GetParamAndAdd("limit", "-n", false)
		cmd.GetParamAndAdd("size", "-s", true)

		// Filters
		var v map[string]map[string]bool

		err := json.Unmarshal([]byte(cmd.Params["filters"][0]), &v)
		if err != nil {
			panic(err)
		}

		var r []string

		for k, val := range v {
			r = append(r, k)

			for ka, _ := range val {
				r = append(r, ka)
			}
		}

		cmd.Add(fmt.Sprintf("--filter \"%s=%s\"", r[0], r[1]))
	}

	return cmd.String()
}

func ContainerCreate(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("run")

	cc := &ContainerCreateConfig{}

	if req.RequestBody != nil {
		if err := json.NewDecoder(bytes.NewReader(req.RequestBody)).Decode(cc); err != nil {
			panic(err)
		}
	}

	if cc.Tty {
		cmd.Add("-t")
	}

	if cc.OpenStdin {
		cmd.Add("-i")
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

	if len(cc.Env) > 0 {
		for _, e := range cc.Env {
			cmd.Add(fmt.Sprintf("-e \"%s\"", e))
		}
	}

	if len(cc.Labels) > 0 {
		for _, l := range cc.Labels {
			cmd.Add(fmt.Sprintf("-l \"%s\"", l))
		}
	}

	if len(cc.User) > 0 {
		cmd.Add(fmt.Sprintf("--user=%s", cc.User))
	}

	if len(cc.WorkingDir) > 0 {
		cmd.Add(fmt.Sprintf("--workdir=%s", cc.WorkingDir))
	}

	if len(cc.Hostname) > 0 {
		cmd.Add(fmt.Sprintf("-h %s", cc.Hostname))
	}

	if len(cc.StopSignal) > 0 {
		cmd.Add(fmt.Sprintf("--stop-signal=%s", cc.StopSignal))
	}

	if len(cc.HostConfig.Binds) > 0 {
		for _, v := range cc.HostConfig.Binds {
			cmd.Add(fmt.Sprintf("-v %s", v))
		}
	}

	if len(cc.HostConfig.Links) > 0 {
		for _, l := range cc.HostConfig.Links {
			cmd.Add(fmt.Sprintf("--link %s", l))
		}
	}

	if cc.HostConfig.Memory > 0 {
		cmd.Add(fmt.Sprintf("-m %s", cc.HostConfig.Memory))
	}

	if cc.HostConfig.MemoryReservation > 0 {
		cmd.Add(fmt.Sprintf("--memory-reservation=%s", cc.HostConfig.MemoryReservation))
	}

	if cc.HostConfig.MemorySwap > 0 {
		cmd.Add(fmt.Sprintf("--memory-swap=%s", cc.HostConfig.MemorySwap))
	}

	if *cc.HostConfig.MemorySwappiness > 0 {
		cmd.Add(fmt.Sprintf("--memory-swappiness=%s", cc.HostConfig.MemorySwappiness))
	}

	if cc.HostConfig.KernelMemory > 0 {
		cmd.Add(fmt.Sprintf("--kernel-memory=%s", cc.HostConfig.KernelMemory))
	}

	if cc.HostConfig.CPUShares > 0 {
		cmd.Add(fmt.Sprintf("--cpu-shares=%s", cc.HostConfig.CPUShares))
	}

	if cc.HostConfig.CPUPeriod > 0 {
		cmd.Add(fmt.Sprintf("--cpu-period=%s", cc.HostConfig.CPUPeriod))
	}

	if cc.HostConfig.CPUQuota > 0 {
		cmd.Add(fmt.Sprintf("--cpu-quota=%s", cc.HostConfig.CPUQuota))
	}

	if len(cc.HostConfig.CpusetCpus) > 0 {
		cmd.Add(fmt.Sprintf("--cpuset-cpus=%s", cc.HostConfig.CpusetCpus))
	}

	if len(cc.HostConfig.CpusetMems) > 0 {
		cmd.Add(fmt.Sprintf("--cpuset-mems=%s", cc.HostConfig.CpusetMems))
	}

	if cc.HostConfig.BlkioWeight > 0 {
		cmd.Add(fmt.Sprintf("--blkio-weight=%s", cc.HostConfig.BlkioWeight))
	}

	if len(cc.HostConfig.BlkioWeightDevice) > 0 {
		cmd.Add(fmt.Sprintf("--blkio-weight-device=%s", cc.HostConfig.BlkioWeightDevice))
	}

	if len(cc.HostConfig.BlkioDeviceReadBps) > 0 {
		for _, drb := range cc.HostConfig.BlkioDeviceReadBps {
			cmd.Add(fmt.Sprintf("--device-read-bps=%s:%s", drb.Path, drb.Rate))
		}
	}

	if len(cc.HostConfig.BlkioDeviceWriteBps) > 0 {
		for _, dwb := range cc.HostConfig.BlkioDeviceWriteBps {
			cmd.Add(fmt.Sprintf("--device-write-bps=%s:%s", dwb.Path, dwb.Rate))
		}
	}

	if len(cc.HostConfig.BlkioDeviceReadIOps) > 0 {
		for _, dri := range cc.HostConfig.BlkioDeviceReadIOps {
			cmd.Add(fmt.Sprintf("--device-read-iops=%s:%s", dri.Path, dri.Rate))
		}
	}

	if len(cc.HostConfig.BlkioDeviceWriteIOps) > 0 {
		for _, dwi := range cc.HostConfig.BlkioDeviceReadIOps {
			cmd.Add(fmt.Sprintf("--device-write-iops=%s:%s", dwi.Path, dwi.Rate))
		}
	}

	if *cc.HostConfig.OomKillDisable {
		cmd.Add("--oom-kill-disable")
	}

	if cc.HostConfig.OomScoreAdj > 0 {
		cmd.Add(fmt.Sprintf("--oom-score-adj=%s", cc.HostConfig.OomScoreAdj))
	}

	if cc.HostConfig.PidsLimit > 0 {
		cmd.Add(fmt.Sprintf("--pids-limit=%s", cc.HostConfig.PidsLimit))
	}

	if cc.HostConfig.Privileged {
		cmd.Add("--privileged")
	}

	if cc.HostConfig.ReadonlyRootfs {
		cmd.Add("--read-only")
	}

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

	if len(cc.HostConfig.ExtraHosts) > 0 {
		for _, eh := range cc.HostConfig.ExtraHosts {
			cmd.Add(fmt.Sprintf("--add-host %s", eh))
		}
	}

	if len(cc.HostConfig.VolumesFrom) > 0 {
		for _, vf := range cc.HostConfig.VolumesFrom {
			cmd.Add(fmt.Sprintf("--volumes-from %s", vf))
		}
	}

	//PortBindings - A map of exposed container ports and the host port they should map to. A JSON object in the form { <port>/<protocol>: [{ "HostPort": "<port>" }] } Take note that port is specified as a string and not an integer value.

	if cc.HostConfig.PublishAllPorts {
		cmd.Add("--publish-all")
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

	if len(cc.HostConfig.GroupAdd) > 0 {
		for _, ga := range cc.HostConfig.GroupAdd {
			cmd.Add(fmt.Sprintf("--group-add %s", ga))
		}
	}

	if len(cc.HostConfig.RestartPolicy.Name) > 0 {
		cmd.Add(fmt.Sprintf("--restart=%s", cc.HostConfig.RestartPolicy.Name))
	}

	if len(cc.HostConfig.UsernsMode) > 0 {
		cmd.Add(fmt.Sprintf("--userns=%s", cc.HostConfig.UsernsMode))
	}

	if len(cc.HostConfig.NetworkMode) > 0 {
		cmd.Add(fmt.Sprintf("--net=%s", cc.HostConfig.NetworkMode))
	}

	if len(cc.HostConfig.Devices) > 0 {
		for _, d := range cc.HostConfig.Devices {
			cmd.Add(fmt.Sprintf("--device=%s", d))
		}
	}

	if len(cc.HostConfig.Ulimits) > 0 {
		for _, u := range cc.HostConfig.Ulimits {
			cmd.Add(fmt.Sprintf("--ulimit=%s", u))
		}
	}

	if len(cc.HostConfig.SecurityOpt) > 0 {
		for _, so := range cc.HostConfig.SecurityOpt {
			cmd.Add(fmt.Sprintf("--security-opt=%s", so))
		}
	}

	//LogConfig - Log configuration for the container, specified as a JSON object in the form { "Type": "<driver_name>", "Config": {"key1": "val1"}}. Available types: json-file, syslog, journald, gelf, fluentd, awslogs, splunk, etwlogs, none. json-file logging driver.

	if len(cc.HostConfig.CgroupParent) > 0 {
		cmd.Add(fmt.Sprintf("--cgroup-parent=%s", cc.HostConfig.CgroupParent))
	}

	if len(cc.HostConfig.VolumeDriver) > 0 {
		cmd.Add(fmt.Sprintf("--volume-driver=%s", cc.HostConfig.VolumeDriver))
	}

	if cc.HostConfig.ShmSize > 0 {
		cmd.Add(fmt.Sprintf("--shm-size=%s", cc.HostConfig.ShmSize))
	}

	if len(cc.Entrypoint) > 0 {
		cmd.Add(fmt.Sprintf("--entrypoint=%s", cc.Entrypoint))
	}

	cmd.GetParams(req.RequestURI)
	if len(cmd.Params) > 0 {
		cmd.GetParamAndAdd("name", "--name", false)
	}

	if len(cc.Image) > 0 {
		cmd.Add(cc.Image)
	}

	if len(cc.Cmd) > 0 {
		cmd.Add(strings.Join(cc.Cmd, " "))
	}

	return cmd.String()
}

func ContainerInspect(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("inspect")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		cmd.GetParamAndAdd("size", "-s", true)
	}

	_, urlPath := utils.GetURIInfo(req)
	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func ContainerTop(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("top")

	_, urlPath := utils.GetURIInfo(req)
	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	cmd.GetParams(req.RequestURI)
	if len(cmd.Params) > 0 {
		if v, ok := cmd.Params["ps_args"]; ok {
			cmd.Add(v[0])
		}
	}

	return cmd.String()
}

func ContainerLogs(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("logs")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		cmd.GetParamAndAdd("follow", "--follow", false)
		cmd.GetParamAndAdd("stdout", "--stdout", false)
		cmd.GetParamAndAdd("stderr", "--stderr", false)
		cmd.GetParamAndAdd("since", "--since", false)
		cmd.GetParamAndAdd("timestamps", "--timestamps", false)
		cmd.GetParamAndAdd("tail", "--tail", false)
	}

	_, urlPath := utils.GetURIInfo(req)
	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func ContainerChanges(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("diff")

	_, urlPath := utils.GetURIInfo(req)
	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func ContainerExport(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("export")

	_, urlPath := utils.GetURIInfo(req)
	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func ContainerStats(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("stats")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		cmd.GetParamAndAdd("stream", "--no-stream", true)
	}

	_, urlPath := utils.GetURIInfo(req)
	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func ContainerResize(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("resize")

	return cmd.String()
}

func ContainerStart(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("start")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		cmd.GetParamAndAdd("detachKeys", "--detach-keys", false)
	}

	_, urlPath := utils.GetURIInfo(req)
	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func ContainerStop(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("stop")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		cmd.GetParamAndAdd("t", "--time", false)
	}

	_, urlPath := utils.GetURIInfo(req)
	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func ContainerRestart(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("restart")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		cmd.GetParamAndAdd("t", "--time", false)
	}

	_, urlPath := utils.GetURIInfo(req)
	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func ContainerKill(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("kill")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		cmd.GetParamAndAdd("signal", "--signal", false)
	}

	_, urlPath := utils.GetURIInfo(req)
	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func ContainerUpdate(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("update")

	uc := &container.UpdateConfig{}

	if req.RequestBody != nil {
		if err := json.NewDecoder(bytes.NewReader(req.RequestBody)).Decode(uc); err != nil {
			panic(err)
		}
	}

	if uc.BlkioWeight > 0 {
		cmd.Add(fmt.Sprintf("--blkio-weight=%s", uc.BlkioWeight))
	}

	if uc.CPUShares > 0 {
		cmd.Add(fmt.Sprintf("--cpu-shares=%s", uc.CPUShares))
	}

	if uc.CPUPeriod > 0 {
		cmd.Add(fmt.Sprintf("--cpu-period=%s", uc.CPUPeriod))
	}

	if uc.CPUQuota > 0 {
		cmd.Add(fmt.Sprintf("--cpu-quota=%s", uc.CPUQuota))
	}

	if len(uc.CpusetCpus) > 0 {
		cmd.Add(fmt.Sprintf("--cpuset-cpus=%s", uc.CpusetCpus))
	}

	if len(uc.CpusetMems) > 0 {
		cmd.Add(fmt.Sprintf("--cpuset-mems=%s", uc.CpusetMems))
	}

	if uc.Memory > 0 {
		cmd.Add(fmt.Sprintf("-m %s", uc.Memory))
	}

	if uc.MemoryReservation > 0 {
		cmd.Add(fmt.Sprintf("--memory-reservation=%s", uc.MemoryReservation))
	}

	if uc.MemorySwap > 0 {
		cmd.Add(fmt.Sprintf("--memory-swap=%s", uc.MemorySwap))
	}

	if *uc.MemorySwappiness > 0 {
		cmd.Add(fmt.Sprintf("--memory-swappiness=%s", uc.MemorySwappiness))
	}

	if uc.KernelMemory > 0 {
		cmd.Add(fmt.Sprintf("--kernel-memory=%s", uc.KernelMemory))
	}

	// TODO: Restart

	return cmd.String()
}

func ContainerRename(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("rename")

	_, urlPath := utils.GetURIInfo(req)
	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	cmd.GetParams(req.RequestURI)

	if v, ok := cmd.Params["name"]; ok {
		cmd.Add(v[0])
	}

	return cmd.String()
}

func ContainerPause(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("pause")

	_, urlPath := utils.GetURIInfo(req)
	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func ContainerUnpause(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("unpause")

	_, urlPath := utils.GetURIInfo(req)
	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func ContainerAttach(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("attach")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		cmd.GetParamAndAdd("detachKeys", "--detach-keys", false)

		if v, ok := cmd.Params["stdin"]; ok {
			if v[0] == "0" || v[0] == "false" || v[0] == "False" {
				cmd.Add("--no-stdin")
			}
		}

	}

	_, urlPath := utils.GetURIInfo(req)
	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func ContainerAttachWS(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("attach")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		cmd.GetParamAndAdd("detachKeys", "--detach-keys", false)

		if v, ok := cmd.Params["stdin"]; ok {
			if v[0] == "0" || v[0] == "false" || v[0] == "False" {
				cmd.Add("--no-stdin")
			}
		}

	}

	_, urlPath := utils.GetURIInfo(req)
	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func ContainerWait(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("wait")

	_, urlPath := utils.GetURIInfo(req)
	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func ContainerRemove(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("rm")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		cmd.GetParamAndAdd("v", "-v", true)
		cmd.GetParamAndAdd("f", "-f", true)
	}

	_, urlPath := utils.GetURIInfo(req)
	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func ContainerCopy(req authorization.Request, re *regexp.Regexp) string {
	return ""
}

func ContainerArchiveInfo(req authorization.Request, re *regexp.Regexp) string {
	return ""
}

func ContainerArchive(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("cp")

	_, urlPath := utils.GetURIInfo(req)
	image := re.FindStringSubmatch(urlPath)[1]

	cmd.GetParams(req.RequestURI)

	cmd.Add(fmt.Sprintf("%s:%s <file>", image, cmd.Params["path"]))

	return cmd.String()
}

func ContainerArchiveExtract(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("cp")

	_, urlPath := utils.GetURIInfo(req)
	image := re.FindStringSubmatch(urlPath)[1]

	cmd.GetParams(req.RequestURI)

	cmd.Add(fmt.Sprintf("<file> %s:%s", image, cmd.Params["path"]))

	return cmd.String()
}

func ContainerExecCreate(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("exec")

	ec := &types.ExecConfig{}

	if req.RequestBody != nil {
		if err := json.NewDecoder(bytes.NewReader(req.RequestBody)).Decode(ec); err != nil {
			panic(err)
		}
	}

	if ec.Detach {
		cmd.Add("-d")
	}

	if len(ec.DetachKeys) > 0 {
		cmd.Add(fmt.Sprintf("--detach-keys=%s", ec.DetachKeys))
	}

	if ec.Tty {
		cmd.Add("-t")
	}

	if ec.AttachStdin {
		cmd.Add("-i")
	}

	if ec.Privileged {
		cmd.Add("--privileged")
	}

	if len(ec.User) > 0 {
		cmd.Add(fmt.Sprintf("-u %s", ec.User))
	}

	_, urlPath := utils.GetURIInfo(req)
	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	if len(ec.Cmd) > 0 {
		for _, c := range ec.Cmd {
			cmd.Add(c)
		}
	}

	return cmd.String()
}
