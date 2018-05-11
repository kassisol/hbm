package dcb

import (
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/pkg/cmdbuilder"
)

func ServiceList(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("service")
	cmd.Add("ls")

	cmd.GetParams(req.RequestURI)

	cmd.AddFilters()

	return cmd.String()
}
