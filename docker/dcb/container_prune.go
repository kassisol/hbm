package dcb

import (
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/pkg/cmdbuilder"
)

func ContainerPrune(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("container")
	cmd.Add("prune")

	cmd.GetParams(req.RequestURI)
	cmd.AddFilters()

	return cmd.String()
}
