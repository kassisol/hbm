package dcb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/docker/engine-api/types/swarm"
	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/pkg/cmdbuilder"
)

func SwarmInspect(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("swarm")
	cmd.Add("info")

	return cmd.String()
}

func SwarmInit(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("swarm")
	cmd.Add("init")

	ir := &swarm.InitRequest{}

	if req.RequestBody != nil {
		if err := json.NewDecoder(bytes.NewReader(req.RequestBody)).Decode(ir); err != nil {
			panic(err)
		}
	}

	if len(ir.ListenAddr) > 0 {
		cmd.Add(fmt.Sprintf("--listen-addr=%s", ir.ListenAddr))
	}

	if len(ir.AdvertiseAddr) > 0 {
		cmd.Add(fmt.Sprintf("--advertise-addr=%s", ir.AdvertiseAddr))
	}

	if ir.ForceNewCluster {
		cmd.Add("--force-new-cluster")
	}

	if ir.Spec.CAConfig.NodeCertExpiry > 0 {
		cmd.Add(fmt.Sprintf("--cert-expiry=%d", ir.Spec.CAConfig.NodeCertExpiry))
	}

	if len(ir.Spec.CAConfig.ExternalCAs) > 0 {
		cmd.Add(fmt.Sprintf("--external-ca=%d", ir.Spec.CAConfig.ExternalCAs))
	}

	if ir.Spec.Dispatcher.HeartbeatPeriod > 0 {
		cmd.Add(fmt.Sprintf("--dispatcher-heartbeat=%d", ir.Spec.Dispatcher.HeartbeatPeriod))
	}

	if ir.Spec.Orchestration.TaskHistoryRetentionLimit > 0 {
		cmd.Add(fmt.Sprintf("--task-history-limit=%d", ir.Spec.Orchestration.TaskHistoryRetentionLimit))
	}

	return cmd.String()
}

func SwarmJoin(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("swarm")
	cmd.Add("join")

	jr := &swarm.JoinRequest{}

	if req.RequestBody != nil {
		if err := json.NewDecoder(bytes.NewReader(req.RequestBody)).Decode(jr); err != nil {
			panic(err)
		}
	}

	if len(jr.ListenAddr) > 0 {
		cmd.Add(fmt.Sprintf("--listen-addr=%s", jr.ListenAddr))
	}

	if len(jr.AdvertiseAddr) > 0 {
		cmd.Add(fmt.Sprintf("--advertise-addr=%s", jr.AdvertiseAddr))
	}

	if len(jr.JoinToken) > 0 {
		cmd.Add(fmt.Sprintf("--token %s", jr.JoinToken))
	}

	if len(jr.RemoteAddrs) > 0 {
		cmd.Add(jr.RemoteAddrs[0])
	}

	return cmd.String()
}

func SwarmLeave(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("swarm")
	cmd.Add("leave")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		cmd.GetParamAndAdd("force", "-f", true)
	}

	return cmd.String()
}

func SwarmUpdate(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("swarm")
	cmd.Add("update")

	spec := &swarm.Spec{}

	if req.RequestBody != nil {
		if err := json.NewDecoder(bytes.NewReader(req.RequestBody)).Decode(spec); err != nil {
			panic(err)
		}
	}

	if spec.CAConfig.NodeCertExpiry > 0 {
		cmd.Add(fmt.Sprintf("--cert-expiry=%d", spec.CAConfig.NodeCertExpiry))
	}

	if len(spec.CAConfig.ExternalCAs) > 0 {
		cmd.Add(fmt.Sprintf("--external-ca=%d", spec.CAConfig.ExternalCAs))
	}

	if spec.Dispatcher.HeartbeatPeriod > 0 {
		cmd.Add(fmt.Sprintf("--dispatcher-heartbeat=%d", spec.Dispatcher.HeartbeatPeriod))
	}

	if spec.Orchestration.TaskHistoryRetentionLimit > 0 {
		cmd.Add(fmt.Sprintf("--task-history-limit=%d", spec.Orchestration.TaskHistoryRetentionLimit))
	}

	return cmd.String()
}
