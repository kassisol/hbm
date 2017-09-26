package dcb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/pkg/cmdbuilder"
	"github.com/moby/moby/api/types/swarm"
)

func ConfigList(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("config")
	cmd.Add("ls")

	cmd.GetParams(req.RequestURI)

	cmd.AddFilters()

	return cmd.String()
}

func ConfigCreate(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("config")
	cmd.Add("create")

	c := &swarm.Config{}

	if req.RequestBody != nil {
		if err := json.NewDecoder(bytes.NewReader(req.RequestBody)).Decode(c); err != nil {
			panic(err)
		}
	}

	if len(c.Spec.Labels) > 0 {
		for k, v := range c.Spec.Labels {
			cmd.Add(fmt.Sprintf("--label=\"%s=%s\"", k, v))
		}
	}

	cmd.Add(c.Spec.Name)

	return cmd.String()
}

func ConfigInspect(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("config")
	cmd.Add("inspect")

	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func ConfigRemove(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("config")
	cmd.Add("rm")

	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func ConfigUpdate(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("config")
	cmd.Add("update")

	c := &swarm.Config{}

	if req.RequestBody != nil {
		if err := json.NewDecoder(bytes.NewReader(req.RequestBody)).Decode(c); err != nil {
			panic(err)
		}
	}

	if len(c.Spec.Labels) > 0 {
		for k, v := range c.Spec.Labels {
			cmd.Add(fmt.Sprintf("--label=\"%s=%s\"", k, v))
		}
	}

	cmd.Add(c.Spec.Name)

	return cmd.String()
}
