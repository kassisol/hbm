package dcb

import (
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/pkg/cmdbuilder"
)

func ContainerRemove(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("container")
	cmd.Add("rm")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		cmd.GetParamAndAdd("force", "--force", true)
		cmd.GetParamAndAdd("link", "--link", true)
		cmd.GetParamAndAdd("v", "--volumes", true)
	}

	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}
