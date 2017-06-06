package dcb

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/pkg/cmdbuilder"
)

func PluginList(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("plugin")
	cmd.Add("ls")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
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

		if v, ok := cmd.Params["filter"]; ok {
			cmd.Add(v[0])
		}
	}

	return cmd.String()
}

func PluginPrivileges(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("plugin")
	cmd.Add("install")

	// TODO

	return cmd.String()
}

func PluginPull(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("plugin")
	cmd.Add("pull")

	// TODO

	return cmd.String()
}

func PluginInspect(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("plugin")
	cmd.Add("inspect")

	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func PluginRemove(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("plugin")
	cmd.Add("rm")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		cmd.GetParamAndAdd("force", "-f", true)
	}

	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func PluginEnable(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("plugin")
	cmd.Add("enable")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		cmd.GetParamAndAdd("timeout", "--timeout", false)
	}

	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func PluginDisable(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("plugin")
	cmd.Add("disable")

	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func PluginUpgrade(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("plugin")
	cmd.Add("upgrade")

	// TODO

	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func PluginCreate(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("plugin")
	cmd.Add("create")

	// TODO

	return cmd.String()
}

func PluginPush(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("plugin")
	cmd.Add("push")

	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func PluginSet(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("plugin")
	cmd.Add("set")

	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	// TODO

	return cmd.String()
}
