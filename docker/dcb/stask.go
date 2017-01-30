package dcb

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/pkg/cmdbuilder"
	"github.com/kassisol/hbm/pkg/utils"
)

func TaskList(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("stask")
	cmd.Add("services")

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

func TaskInspect(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("stask")
	cmd.Add("tasks")

	_, urlPath := utils.GetURIInfo(req)
	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}
