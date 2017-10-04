package dcb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/docker/docker/api/types"
	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/pkg/cmdbuilder"
)

func Auth(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("login")

	ac := &types.AuthConfig{}

	if req.RequestBody != nil {
		if err := json.NewDecoder(bytes.NewReader(req.RequestBody)).Decode(ac); err != nil {
			panic(err)
		}
	}

	if len(ac.Username) > 0 {
		cmd.Add(fmt.Sprintf("-u %s", ac.Username))
	}

	if len(ac.ServerAddress) > 0 {
		cmd.Add(ac.ServerAddress)
	}

	return cmd.String()
}

func Info(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("system")
	cmd.Add("info")

	return cmd.String()
}

func Version(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("version")

	return cmd.String()
}

func Ping(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("*ping*")

	return cmd.String()
}

func Events(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("system")
	cmd.Add("events")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		cmd.GetParamAndAdd("since", "--since", false)
		cmd.GetParamAndAdd("until", "--until", false)

		cmd.AddFilters()
	}

	return cmd.String()
}

func SystemDF(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("system")
	cmd.Add("df")

	return cmd.String()
}

func Anyroute(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	return ""
}
