package dcb

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/pkg/cmdbuilder"
	"github.com/kassisol/hbm/pkg/utils"
)

func ImageList(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("images")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		cmd.GetParamAndAdd("all", "-a", true)

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

		if v, ok := cmd.Params["filter"]; ok {
			cmd.Add(v[0])
		}
	}

	return cmd.String()
}

func ImageBuild(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("build")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		cmd.GetParamAndAdd("t", "--tag", false)
		cmd.GetParamAndAdd("q", "-q", true)
		cmd.GetParamAndAdd("nocache", "--no-cache", true)
		cmd.GetParamAndAdd("pull", "--pull", true)
		cmd.GetParamAndAdd("rm", "--rm", true)
		cmd.GetParamAndAdd("forcerm", "--force-rm", true)
		cmd.GetParamAndAdd("memory", "--memory", false)
		cmd.GetParamAndAdd("memswap", "--memory-swap", false)
		cmd.GetParamAndAdd("cpushares", "--cpu-shares", true)
		cmd.GetParamAndAdd("cpusetcpus", "--cpuset-cpus", false)
		cmd.GetParamAndAdd("cpusetmems", "--cpuset-mems", false)
		cmd.GetParamAndAdd("cpuperiod", "--cpu-period", false)
		cmd.GetParamAndAdd("cpuquota", "--cpu-quota", false)

		// TODO: buildargs

		cmd.GetParamAndAdd("shmsize", "--shm-size", false)

		// TODO: labels

		if v, ok := cmd.Params["remote"]; ok {
			cmd.Add(v[0])
		} else {
			cmd.GetParamAndAdd("dockerfile", "--file", false)
		}
	}

	return cmd.String()
}

func ImageCreate(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("pull")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		if v, ok := cmd.Params["fromImage"]; ok {
			cmd.Add(v[0])
		}
	}

	return cmd.String()
}

func ImageInspect(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("inspect")

	_, urlPath := utils.GetURIInfo(req)
	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func ImageHistory(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("history")

	_, urlPath := utils.GetURIInfo(req)
	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func ImagePush(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("push")

	_, urlPath := utils.GetURIInfo(req)
	image := re.FindStringSubmatch(urlPath)[1]

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		if v, ok := cmd.Params["tag"]; ok {
			cmd.Add(fmt.Sprintf("%s:%s", image, v[0]))
		}
	} else {
		cmd.Add(image)
	}

	return cmd.String()
}

func ImageTag(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("tag")

	_, urlPath := utils.GetURIInfo(req)
	image := re.FindStringSubmatch(urlPath)[1]

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		if v, ok := cmd.Params["tag"]; ok {
			cmd.Add(fmt.Sprintf("%s:%s", image, v[0]))
		}
	}

	return cmd.String()
}

func ImageRemove(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("rmi")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		cmd.GetParamAndAdd("force", "-f", true)
		cmd.GetParamAndAdd("noprune", "--no-prune", true)
	}

	_, urlPath := utils.GetURIInfo(req)
	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func ImageSearch(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("search")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		if v, ok := cmd.Params["term"]; ok {
			cmd.Add(v[0])
		}
	}

	return cmd.String()
}

func ImageSaveImage(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("save")

	_, urlPath := utils.GetURIInfo(req)
	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func ImageSaveImages(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("save")

	return cmd.String()
}

func ImageLoad(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("load")

	return cmd.String()
}
