package dcb

import (
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/pkg/cmdbuilder"
)

func TaskList(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("stask")
	cmd.Add("services")

	return cmd.String()
}

func TaskInspect(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("stask")
	cmd.Add("tasks")

	return cmd.String()
}
