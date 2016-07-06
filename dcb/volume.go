package dcb

import (
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/harbourmaster/hbm/pkg/cmdbuilder"
)

func VolumeList(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("volume")
	cmd.Add("ls")

	return cmd.String()
}

func VolumeCreate(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("volume")
	cmd.Add("create")

	return cmd.String()
}

func VolumeInspect(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("volume")
	cmd.Add("inspect")

	return cmd.String()
}

func VolumeRemove(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("volume")
	cmd.Add("rm")

	return cmd.String()
}
