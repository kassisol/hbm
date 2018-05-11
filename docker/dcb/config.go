package dcb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/pkg/cmdbuilder"
)

func ConfigCreate(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("config")
	cmd.Add("create")

	c := &swarm.ConfigSpec{}

	if req.RequestBody != nil {
		if err := json.NewDecoder(bytes.NewReader(req.RequestBody)).Decode(c); err != nil {
			panic(err)
		}
	}

	if len(c.Annotations.Labels) > 0 {
		for k, v := range c.Annotations.Labels {
			cmd.Add(fmt.Sprintf("--label \"%s=%s\"", k, v))
		}
	}

	// --template-driver

	cmd.Add(c.Annotations.Name)

	return cmd.String()
}

func ConfigInspect(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("config")
	cmd.Add("inspect")

	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func ConfigList(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("config")
	cmd.Add("ls")

	cmd.GetParams(req.RequestURI)

	cmd.AddFilters()

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

	c := &swarm.ConfigSpec{}

	if req.RequestBody != nil {
		if err := json.NewDecoder(bytes.NewReader(req.RequestBody)).Decode(c); err != nil {
			panic(err)
		}
	}

	if len(c.Annotations.Labels) > 0 {
		for k, v := range c.Annotations.Labels {
			cmd.Add(fmt.Sprintf("--label \"%s=%s\"", k, v))
		}
	}

	cmd.Add(c.Annotations.Name)

	return cmd.String()
}
