package dcb

import (
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/pkg/cmdbuilder"
)

func ContainerAttach(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("container")
	cmd.Add("attach")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		cmd.GetParamAndAdd("detachKeys", "--detach-keys", false)

		if v, ok := cmd.Params["stdin"]; ok {
			if v[0] == "0" || v[0] == "false" || v[0] == "False" {
				cmd.Add("--no-stdin")
			}
		}

		// --sig-proxy
	}

	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}
