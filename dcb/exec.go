package dcb

import (
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/harbourmaster/hbm/pkg/cmdbuilder"
)

func ExecStart(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("exec")

	return cmd.String()
}

func ExecResize(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("exec")

	return cmd.String()
}

func ExecInspect(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("exec")

	return cmd.String()
}
