package dcb

import (
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/pkg/cmdbuilder"
)

func ServiceList(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("service")
	cmd.Add("ls")

	return cmd.String()
}

func ServiceCreate(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("service")
	cmd.Add("create")

	return cmd.String()
}

func ServiceRemove(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("service")
	cmd.Add("rm")

	return cmd.String()
}

func ServiceInspect(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("service")
	cmd.Add("inspect")

	return cmd.String()
}

func ServiceUpdate(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("service")
	cmd.Add("update")

	return cmd.String()
}
