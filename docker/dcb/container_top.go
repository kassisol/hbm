package dcb

import (
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/pkg/cmdbuilder"
)

func ContainerTop(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("container")
	cmd.Add("top")

	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		if v, ok := cmd.Params["ps_args"]; ok {
			cmd.Add(v[0])
		}
	}

	return cmd.String()
}
