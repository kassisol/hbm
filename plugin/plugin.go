package plugin

import (
	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/pkg/uri"
)

// Plugin structure
type Plugin struct {
	appPath string
}

// NewPlugin function
func NewPlugin(appPath string) (*Plugin, error) {
	return &Plugin{appPath: appPath}, nil
}

// AuthZReq function
func (p *Plugin) AuthZReq(req authorization.Request) authorization.Response {
	uriinfo, err := uri.GetURIInfo(SupportedVersion, req)
	if err != nil {
		return authorization.Response{Err: err.Error()}
	}

	a, err := NewAPI(uriinfo.Version, p.appPath)
	if err != nil {
		return authorization.Response{Err: err.Error()}
	}

	r := a.Allow(req)
	if r.Error != "" {
		return authorization.Response{Err: r.Error}
	}
	if !r.Allow {
		return authorization.Response{Msg: r.Msg}
	}

	return authorization.Response{Allow: true}
}

// AuthZRes function
func (p *Plugin) AuthZRes(req authorization.Request) authorization.Response {
	return authorization.Response{Allow: true}
}
