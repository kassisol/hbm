package dcb

import (
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/pkg/cmdbuilder"
)

func ContainerRename(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("container")
	cmd.Add("rename")

	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	cmd.GetParams(req.RequestURI)

	if v, ok := cmd.Params["name"]; ok {
		cmd.Add(v[0])
	}

	return cmd.String()
}
