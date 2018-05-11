package dcb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/pkg/cmdbuilder"
)

func ContainerUpdate(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("container")
	cmd.Add("update")

	uc := &container.UpdateConfig{}

	if req.RequestBody != nil {
		if err := json.NewDecoder(bytes.NewReader(req.RequestBody)).Decode(uc); err != nil {
			panic(err)
		}
	}

	if uc.BlkioWeight > 0 {
		cmd.Add(fmt.Sprintf("--blkio-weight %d", uc.BlkioWeight))
	}

	if uc.CPUPeriod > 0 {
		cmd.Add(fmt.Sprintf("--cpu-period %d", uc.CPUPeriod))
	}

	if uc.CPUQuota > 0 {
		cmd.Add(fmt.Sprintf("--cpu-quota %d", uc.CPUQuota))
	}

	if uc.CPURealtimePeriod > 0 {
		cmd.Add(fmt.Sprintf("--cpu-rt-period %d", uc.CPURealtimePeriod))
	}

	if uc.CPURealtimeRuntime > 0 {
		cmd.Add(fmt.Sprintf("--cpu-rt-runtime %d", uc.CPURealtimeRuntime))
	}

	if uc.CPUShares > 0 {
		cmd.Add(fmt.Sprintf("--cpu-shares %d", uc.CPUShares))
	}

	if uc.NanoCPUs > 0 {
		cmd.Add(fmt.Sprintf("--cpus %d", uc.NanoCPUs))
	}

	if len(uc.CpusetCpus) > 0 {
		cmd.Add(fmt.Sprintf("--cpuset-cpus %s", uc.CpusetCpus))
	}

	if len(uc.CpusetMems) > 0 {
		cmd.Add(fmt.Sprintf("--cpuset-mems %s", uc.CpusetMems))
	}

	if uc.KernelMemory > 0 {
		cmd.Add(fmt.Sprintf("--kernel-memory %d", uc.KernelMemory))
	}

	if uc.Memory > 0 {
		cmd.Add(fmt.Sprintf("--memory %d", uc.Memory))
	}

	if uc.MemoryReservation > 0 {
		cmd.Add(fmt.Sprintf("--memory-reservation %d", uc.MemoryReservation))
	}

	if uc.MemorySwap > 0 {
		cmd.Add(fmt.Sprintf("--memory-swap %d", uc.MemorySwap))
	}

	if uc.MemorySwappiness != nil {
		if *uc.MemorySwappiness > 0 {
			cmd.Add(fmt.Sprintf("--memory-swappiness %d", uc.MemorySwappiness))
		}
	}

	if len(uc.RestartPolicy.Name) > 0 {
		cmd.Add(fmt.Sprintf("--restart %s:%d", uc.RestartPolicy.Name, uc.RestartPolicy.MaximumRetryCount))
	}

	return cmd.String()
}
