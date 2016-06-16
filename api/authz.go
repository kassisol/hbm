package api

import (
	"log"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/harbourmaster/hbm/pkg/utils"
)

type plugin struct {
	appPath	string
}

func NewPlugin(appPath string) (*plugin, error) {
	return &plugin{appPath: appPath}, nil
}

func (p *plugin) AuthZReq(req authorization.Request) authorization.Response {
	apiver, _ := utils.GetURIInfo(req)

	a, err := NewApi(apiver, p.appPath)
	if err != nil {
		log.Fatal(err)
	}

	allow, e, msg := a.Allow(req)
	if e != "" {
		return authorization.Response{Err: e}
	}
	if ! allow {
		return authorization.Response{Msg: msg}
	}

	return authorization.Response{Allow: true}
}

func (p *plugin) AuthZRes(req authorization.Request) authorization.Response {
	return authorization.Response{Allow: true}
}
