package dcb

import (
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/pkg/cmdbuilder"
)

func ServiceInspect(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("service")
	cmd.Add("inspect")

	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}
