package dcb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/docker/engine-api/types"
	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/pkg/cmdbuilder"
)

// Anyroute function
func Anyroute(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	return ""
}

// Auth function
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

// Info function
func Info(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("info")

	return cmd.String()
}

// Version function
func Version(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("version")

	return cmd.String()
}

// Ping function
func Ping(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	return ""
}

// Commit function
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

// Events function
func Events(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("events")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		cmd.GetParamAndAdd("since", "--since", false)
		cmd.GetParamAndAdd("until", "--until", false)

		// Filters
		if _, ok := cmd.Params["filters"]; ok {
			var v map[string][]string

			err := json.Unmarshal([]byte(cmd.Params["filters"][0]), &v)
			if err != nil {
				panic(err)
			}

			for k, val := range v {
				for _, f := range val {
					cmd.Add(fmt.Sprintf("--filter \"%s=%s\"", k, f))
				}
			}
		}
	}

	return cmd.String()
}
