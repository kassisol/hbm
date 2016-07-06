package dcb

import (
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/harbourmaster/hbm/pkg/cmdbuilder"
)

func ContainerList(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("ps")

	return cmd.String()
}

func ContainerCreate(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("run")

	return cmd.String()
}

func ContainerInspect(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("inspect")

	return cmd.String()
}

func ContainerTop(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("top")

	return cmd.String()
}

func ContainerLogs(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("logs")

	return cmd.String()
}

func ContainerChanges(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("change")

	return cmd.String()
}

func ContainerExport(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("export")

	return cmd.String()
}

func ContainerStats(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("stats")

	return cmd.String()
}

func ContainerResize(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("resize")

	return cmd.String()
}

func ContainerStart(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("start")

	return cmd.String()
}

func ContainerStop(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("stop")

	return cmd.String()
}

func ContainerRestart(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("restart")

	return cmd.String()
}

func ContainerKill(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("kill")

	return cmd.String()
}

func ContainerUpdate(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("update")

	return cmd.String()
}

func ContainerRename(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("rename")

	return cmd.String()
}

func ContainerPause(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("pause")

	return cmd.String()
}

func ContainerUnpause(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("unpause")

	return cmd.String()
}

func ContainerAttachWS(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("attach")

	return cmd.String()
}

func ContainerAttach(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("attach")

	return cmd.String()
}

func ContainerWait(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("wait")

	return cmd.String()
}

func ContainerRemove(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("rm")

	return cmd.String()
}

func ContainerCopy(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("cp")

	return cmd.String()
}

func ContainerArchiveInfo(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("archive")

	return cmd.String()
}

func ContainerArchive(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("archive")

	return cmd.String()
}

func ContainerArchiveExtract(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("archive")

	return cmd.String()
}

func ContainerExecCreate(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("exec")

	return cmd.String()
}
