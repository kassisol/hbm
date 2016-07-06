package dcb

import (
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/harbourmaster/hbm/pkg/cmdbuilder"
)

func Anyroute(req authorization.Request, re *regexp.Regexp) string {
	return ""
}

func Auth(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("login")

	return cmd.String()
}

func Info(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("info")

	return cmd.String()
}

func Version(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("version")

	return cmd.String()
}

func Ping(req authorization.Request, re *regexp.Regexp) string {
	return ""
}

func Commit(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("commit")

	return cmd.String()
}

func Events(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("events")

	return cmd.String()
}
