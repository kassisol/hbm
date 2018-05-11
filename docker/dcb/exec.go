package dcb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/docker/docker/api/types"
	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/pkg/cmdbuilder"
)

func ContainerExecCreate(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("exec")

	ec := &types.ExecConfig{}

	if req.RequestBody != nil {
		if err := json.NewDecoder(bytes.NewReader(req.RequestBody)).Decode(ec); err != nil {
			panic(err)
		}
	}

	if ec.Detach {
		cmd.Add("--detach")
	}

	if len(ec.DetachKeys) > 0 {
		cmd.Add(fmt.Sprintf("--detach-keys=%s", ec.DetachKeys))
	}

	if len(ec.Env) > 0 {
		for _, e := range ec.Env {
			cmd.Add(fmt.Sprintf("--env %s", e))
		}
	}

	if ec.AttachStdin {
		cmd.Add("--interactive")
	}

	if ec.Privileged {
		cmd.Add("--privileged")
	}

	if ec.Tty {
		cmd.Add("--tty")
	}

	if len(ec.User) > 0 {
		cmd.Add(fmt.Sprintf("--user %s", ec.User))
	}

	if len(ec.WorkingDir) > 0 {
		cmd.Add("--workdir")
	}

	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	if len(ec.Cmd) > 0 {
		for _, c := range ec.Cmd {
			cmd.Add(c)
		}
	}

	return cmd.String()
}

func ExecStart(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("exec")

	return cmd.String()
}

func ExecResize(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("exec")

	return cmd.String()
}

func ExecInspect(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("exec")

	return cmd.String()
}
