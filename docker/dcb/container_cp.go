package dcb

import (
	"fmt"
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/pkg/cmdbuilder"
)

func ContainerArchiveInfo(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("*archive head*")

	return cmd.String()
}

func ContainerArchive(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("container")
	cmd.Add("cp")

	image := re.FindStringSubmatch(urlPath)[1]

	cmd.GetParams(req.RequestURI)

	// --archive
	// --follow-link

	if v, ok := cmd.Params["path"]; ok {
		cmd.Add(fmt.Sprintf("%s:%s <file>", image, v))
	}

	return cmd.String()
}

func ContainerArchiveExtract(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("container")
	cmd.Add("cp")

	image := re.FindStringSubmatch(urlPath)[1]

	cmd.GetParams(req.RequestURI)

	if v, ok := cmd.Params["path"]; ok {
		cmd.Add(fmt.Sprintf("<file> %s:%s", image, v))
	}

	return cmd.String()
}
