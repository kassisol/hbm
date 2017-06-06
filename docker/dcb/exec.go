package dcb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/pkg/cmdbuilder"
	"github.com/moby/moby/api/types"
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
