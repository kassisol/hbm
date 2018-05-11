package dcb

import (
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/pkg/cmdbuilder"
)

func ServiceLogs(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("service")
	cmd.Add("logs")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		cmd.GetParamAndAdd("details", "--details", true)
		cmd.GetParamAndAdd("follow", "--follow", true)
		cmd.GetParamAndAdd("stdout", "--stdout", true)
		cmd.GetParamAndAdd("stderr", "--stderr", true)
		cmd.GetParamAndAdd("since", "--since", false)
		cmd.GetParamAndAdd("timestamps", "--timestamps", true)
		cmd.GetParamAndAdd("tail", "--tail", false)
	}

	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}
