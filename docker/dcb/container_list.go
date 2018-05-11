package dcb

import (
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/pkg/cmdbuilder"
)

func ContainerList(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("container")
	cmd.Add("ls")

	cmd.GetParams(req.RequestURI)

	cmd.GetParamAndAdd("all", "-a", true)
	cmd.GetParamAndAdd("last", "-n", false)
	cmd.GetParamAndAdd("size", "-s", true)

	cmd.AddFilters()

	return cmd.String()
}

