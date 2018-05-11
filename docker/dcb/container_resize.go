package dcb

import (
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/pkg/cmdbuilder"
)

func ContainerResize(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("*resize*")

	return cmd.String()
}
