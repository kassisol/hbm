package dcb

import (
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/harbourmaster/hbm/pkg/cmdbuilder"
)

func NodesList(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("ls")

	return cmd.String()
}

func NodesInspect(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("inspect")

	return cmd.String()
}

func NodesDelete(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("rm")

	return cmd.String()
}

func NodesUpdate(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("update")

	return cmd.String()
}
