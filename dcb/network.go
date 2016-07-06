package dcb

import (
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/harbourmaster/hbm/pkg/cmdbuilder"
)

func NetworkList(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("network")
	cmd.Add("ls")

	return cmd.String()
}

func NetworkInspect(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("network")
	cmd.Add("inspect")

	return cmd.String()
}

func NetworkCreate(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("network")
	cmd.Add("create")

	return cmd.String()
}

func NetworkConnect(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("network")
	cmd.Add("connect")

	return cmd.String()
}

func NetworkDisconnect(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("network")
	cmd.Add("disconnect")

	return cmd.String()
}

func NetworkRemove(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("network")
	cmd.Add("rm")

	return cmd.String()
}
