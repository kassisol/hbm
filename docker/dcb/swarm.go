package dcb

import (
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/pkg/cmdbuilder"
)

func SwarmInspect(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("swarm")
	cmd.Add("info")

	return cmd.String()
}

func SwarmInit(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("swarm")
	cmd.Add("init")

	return cmd.String()
}

func SwarmJoin(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("swarm")
	cmd.Add("join")

	return cmd.String()
}

func SwarmLeave(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("swarm")
	cmd.Add("leave")

	return cmd.String()
}

func SwarmUpdate(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("swarm")
	cmd.Add("update")

	return cmd.String()
}
