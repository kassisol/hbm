package dcb

import (
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/harbourmaster/hbm/pkg/cmdbuilder"
)

func ServicesList(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("ls")

	return cmd.String()
}

func ServicesInspect(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("inspect")

	return cmd.String()
}

func ServicesDelete(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("rm")

	return cmd.String()
}

func ServicesCreate(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("create")

	return cmd.String()
}

func ServicesUpdate(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("update")

	return cmd.String()
}
