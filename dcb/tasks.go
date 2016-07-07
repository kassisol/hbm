package dcb

import (
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/harbourmaster/hbm/pkg/cmdbuilder"
)

func TasksLists(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("ls")

	return cmd.String()
}

func TasksInspect(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("inspect")

	return cmd.String()
}
