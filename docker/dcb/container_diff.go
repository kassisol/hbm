package dcb

import (
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/pkg/cmdbuilder"
)

func ContainerChanges(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("container")
	cmd.Add("diff")

	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}
