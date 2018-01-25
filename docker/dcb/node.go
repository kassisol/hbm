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

func NodeList(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("node")
	cmd.Add("ls")

	cmd.GetParams(req.RequestURI)

	cmd.AddFilters()

	return cmd.String()
}

func NodeInspect(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("node")
	cmd.Add("inspect")

	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func NodeRemove(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("node")
	cmd.Add("rm")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		cmd.GetParamAndAdd("force", "-f", true)
	}

	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func NodeUpdate(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("node")
	cmd.Add("update")

	ns := &swarm.NodeSpec{}

	if req.RequestBody != nil {
		if err := json.NewDecoder(bytes.NewReader(req.RequestBody)).Decode(ns); err != nil {
			panic(err)
		}
	}

	if len(ns.Labels) > 0 {
		for k, v := range ns.Labels {
			cmd.Add(fmt.Sprintf("--label \"%s=%s\"", k, v))
		}
	}

	if len(ns.Availability) > 0 {
		cmd.Add(fmt.Sprintf("--availability %s", ns.Availability))
	}

	if len(ns.Role) > 0 {
		cmd.Add(fmt.Sprintf("--role %s", ns.Role))
	}

	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}
