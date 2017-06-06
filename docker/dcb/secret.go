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

func SecretList(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("secret")
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

func SecretCreate(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("secret")
	cmd.Add("create")

	s := &swarm.Secret{}

	if req.RequestBody != nil {
		if err := json.NewDecoder(bytes.NewReader(req.RequestBody)).Decode(s); err != nil {
			panic(err)
		}
	}

	if len(s.Spec.Labels) > 0 {
		for k, v := range s.Spec.Labels {
			cmd.Add(fmt.Sprintf("--label=\"%s=%s\"", k, v))
		}
	}

	cmd.Add(s.Spec.Name)

	return cmd.String()
}

func SecretInspect(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("secret")
	cmd.Add("inspect")

	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func SecretRemove(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("secret")
	cmd.Add("rm")

	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func SecretUpdate(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("secret")
	cmd.Add("update")

	s := &swarm.Secret{}

	if req.RequestBody != nil {
		if err := json.NewDecoder(bytes.NewReader(req.RequestBody)).Decode(s); err != nil {
			panic(err)
		}
	}

	if len(s.Spec.Labels) > 0 {
		for k, v := range s.Spec.Labels {
			cmd.Add(fmt.Sprintf("--label=\"%s=%s\"", k, v))
		}
	}

	cmd.Add(s.Spec.Name)

	return cmd.String()
}
