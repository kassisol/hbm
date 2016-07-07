package dcb

import (
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/harbourmaster/hbm/pkg/cmdbuilder"
)

func SwarmInit(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("init")

	return cmd.String()
}

func SwarmJoin(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("join")

	return cmd.String()
}

func SwarmLeave(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("leave")

	return cmd.String()
}

func SwarmUpdate(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("update")

	return cmd.String()
}
