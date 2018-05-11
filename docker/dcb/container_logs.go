package dcb

import (
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/pkg/cmdbuilder"
)

func ContainerLogs(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("container")
	cmd.Add("logs")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		cmd.GetParamAndAdd("follow", "--follow", false)
		cmd.GetParamAndAdd("since", "--since", false)
		cmd.GetParamAndAdd("tail", "--tail", false)
		cmd.GetParamAndAdd("timestamps", "--timestamps", false)
		cmd.GetParamAndAdd("until", "--until", false)
	}

	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}
