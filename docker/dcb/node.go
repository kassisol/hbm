package dcb

import (
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/pkg/cmdbuilder"
)

func NodeList(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("node")
	cmd.Add("ls")

	return cmd.String()
}

func NodeInspect(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("node")
	cmd.Add("inspect")

	return cmd.String()
}

func NodeRemove(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("node")
	cmd.Add("rm")

	return cmd.String()
}

func NodeUpdate(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("node")
	cmd.Add("update")

	return cmd.String()
}
