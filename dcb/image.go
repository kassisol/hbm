package dcb

import (
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/harbourmaster/hbm/pkg/cmdbuilder"
)

func ImageList(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("images")

	return cmd.String()
}

func ImageBuild(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("build")

	return cmd.String()
}

func ImageCreate(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("pull")

	return cmd.String()
}

func ImageInspect(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("inspect")

	return cmd.String()
}

func ImageHistory(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("history")

	return cmd.String()
}

func ImagePush(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("push")

	return cmd.String()
}

func ImageTag(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("tag")

	return cmd.String()
}

func ImageRemove(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("rmi")

	return cmd.String()
}

func ImageSearch(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("search")

	return cmd.String()
}

func ImageSaveImage(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("images")

	return cmd.String()
}

func ImageSaveImages(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("images")

	return cmd.String()
}

func ImageLoad(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("images")

	return cmd.String()
}
