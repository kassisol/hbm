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

func Anyroute(req authorization.Request, re *regexp.Regexp) string {
	return ""
}

func Auth(req authorization.Request, re *regexp.Regexp) string {
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

func Info(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("info")

	return cmd.String()
}

func Version(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("version")

	return cmd.String()
}

func Ping(req authorization.Request, re *regexp.Regexp) string {
	return ""
}

func Commit(req authorization.Request, re *regexp.Regexp) string {
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

func Events(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("events")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		cmd.GetParamAndAdd("since", "--since", false)
		cmd.GetParamAndAdd("until", "--until", false)

		// Filters
		if _, ok := cmd.Params["filters"]; ok {
			var v map[string]map[string]bool

			err := json.Unmarshal([]byte(cmd.Params["filters"][0]), &v)
			if err != nil {
				panic(err)
			}

			var r []string

			for k, val := range v {
				r = append(r, k)

				for ka, _ := range val {
					r = append(r, ka)
				}
			}

			cmd.Add(fmt.Sprintf("--filter \"%s=%s\"", r[0], r[1]))
		}
	}

	return cmd.String()
}
