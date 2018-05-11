package dcb

import (
	"fmt"
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/pkg/cmdbuilder"
)

func ImageBuild(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("image")
	cmd.Add("build")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		cmd.GetParamAndAdd("extrahosts", "--add-host", false)
		// --build-arg
		cmd.GetParamAndAdd("cachefrom", "--cache-from", true)
		cmd.GetParamAndAdd("cpuperiod", "--cpu-period", false)
		cmd.GetParamAndAdd("cpuquota", "--cpu-quota", false)
		cmd.GetParamAndAdd("cpushares", "--cpu-shares", true)
		cmd.GetParamAndAdd("cpusetcpus", "--cpuset-cpus", false)
		cmd.GetParamAndAdd("cpusetmems", "--cpuset-mems", false)
		cmd.GetParamAndAdd("forcerm", "--force-rm", true)
		// --label
		cmd.GetParamAndAdd("memory", "--memory", false)
		cmd.GetParamAndAdd("memswap", "--memory-swap", false)
		cmd.GetParamAndAdd("networkmode", "--network", false)
		cmd.GetParamAndAdd("nocache", "--no-cache", true)
		cmd.GetParamAndAdd("pull", "--pull", true)
		cmd.GetParamAndAdd("q", "--quiet", true)
		cmd.GetParamAndAdd("rm", "--rm", true)
		cmd.GetParamAndAdd("shmsize", "--shm-size", false)
		cmd.GetParamAndAdd("t", "--tag", false)

		if v, ok := cmd.Params["remote"]; ok {
			cmd.Add(v[0])
		} else {
			cmd.GetParamAndAdd("dockerfile", "--file", false)
		}
	}

	return cmd.String()
}

func ImageHistory(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("image")
	cmd.Add("history")

	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func ImageInspect(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("image")
	cmd.Add("inspect")

	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func ImageLoad(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("image")
	cmd.Add("load")

	return cmd.String()
}

func ImageList(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("image")
	cmd.Add("ls")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		cmd.GetParamAndAdd("all", "--all", true)
		cmd.GetParamAndAdd("digests", "--digests", true)

		cmd.AddFilters()
	}

	return cmd.String()
}

func ImagePrune(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("image")
	cmd.Add("prune")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		cmd.AddFilters()
	}

	return cmd.String()
}

func ImageCreate(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("image")
	cmd.Add("pull")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		if v, ok := cmd.Params["tag"]; ok {
			if len(v) == 0 {
				cmd.GetParamAndAdd("tag", "--all-tags", false)
			}
		}

		if v, ok := cmd.Params["fromImage"]; ok {
			cmd.Add(v[0])
		}
	}

	return cmd.String()
}

func ImagePush(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("image")
	cmd.Add("push")

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

func ImageRemove(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("image")
	cmd.Add("rm")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		cmd.GetParamAndAdd("force", "--force", true)
		cmd.GetParamAndAdd("noprune", "--no-prune", true)
	}

	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func ImageSaveImage(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("image")
	cmd.Add("save")

	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func ImageSaveImages(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("image")
	cmd.Add("save")

	return cmd.String()
}

func ImageTag(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("image")
	cmd.Add("tag")

	image := re.FindStringSubmatch(urlPath)[1]

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		if v, ok := cmd.Params["tag"]; ok {
			cmd.Add(fmt.Sprintf("%s:%s", image, v[0]))
		}
	}

	return cmd.String()
}

func ImageSearch(req authorization.Request, urlPath string, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("search")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		if v, ok := cmd.Params["term"]; ok {
			cmd.Add(v[0])
		}

		cmd.AddFilters()
	}

	return cmd.String()
}
