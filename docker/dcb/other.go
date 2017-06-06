package dcb

import (
	"fmt"
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/pkg/cmdbuilder"
)

func Commit(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("commit")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		cmd.GetParamAndAdd("author", "--author", false)
		cmd.GetParamAndAdd("comment", "--message", false)
		cmd.GetParamAndAdd("pause", "-p", true)

		// TODO: changes ; --change=[]

		if v, ok := cmd.Params["container"]; ok {
			cmd.Add(v[0])
		}

		if v, ok := cmd.Params["repo"]; ok {
			cmd.Add(v[0])
		}
		if v, ok := cmd.Params["tag"]; ok {
			cmd.Add(fmt.Sprintf(":%s", v[0]))
		}
	}

	return cmd.String()
}

func TaskLogs(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("task")
	cmd.Add("logs")

	return cmd.String()
}
