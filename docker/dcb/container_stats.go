package dcb

import (
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/pkg/cmdbuilder"
)

func ContainerStats(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("container")
	cmd.Add("stats")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		cmd.GetParamAndAdd("stream", "--no-stream", true)
	}

	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}
