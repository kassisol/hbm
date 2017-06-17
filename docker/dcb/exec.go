package dcb

import (
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/pkg/cmdbuilder"
)

// ExecStart function
func ExecStart(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("exec")

	return cmd.String()
}

// ExecResize function
func ExecResize(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("exec")

	return cmd.String()
}

// ExecInspect function
func ExecInspect(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("exec")

	return cmd.String()
}
