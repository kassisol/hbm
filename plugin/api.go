package plugin

import (
	"fmt"
	"os"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/juliengk/go-log"
	"github.com/juliengk/go-log/driver"
	"github.com/kassisol/hbm/allow"
	"github.com/kassisol/hbm/allow/types"
	"github.com/kassisol/hbm/docker/endpoint"
	"github.com/kassisol/hbm/pkg/uri"
	"github.com/kassisol/hbm/storage"
	"github.com/kassisol/hbm/version"
)

var SupportedVersion = "v1.24"

type Api struct {
	Uris    *uri.URIs
	AppPath string
}

func NewApi(version, appPath string) (*Api, error) {
	if version != SupportedVersion {
		return &Api{}, fmt.Errorf("This version of HBM does not support Docker API version %s. Supported version is %s", version, SupportedVersion)
	}

	uris := endpoint.GetUris()

	return &Api{Uris: uris, AppPath: appPath}, nil
}

func (a *Api) Allow(req authorization.Request) *types.AllowResult {
	l, _ := log.NewDriver("standard", nil)

	uriinfo, err := uri.GetURIInfo(req)
	if err != nil {
		// Log event
		l.Warning(err)

		return &types.AllowResult{Allow: false, Error: err.Error()}
	}

	user := req.User
	if user == "" {
		user = "root"
	}

	host, err := os.Hostname()
	if err != nil {
		host = "localhost"
	}

	config := types.Config{AppPath: a.AppPath, Username: user, Hostname: host}

	u, err := a.Uris.GetURI(req.RequestMethod, uriinfo.Path)
	if err != nil {
		msg := fmt.Sprintf("%s is not implemented", uriinfo.Path)

		// Log event
		l.Warning(msg)

		return &types.AllowResult{Allow: false, Error: msg}
	}

	r := allow.AllowTrue(req, &config)

	s, err := storage.NewDriver("sqlite", a.AppPath)
	if err != nil {
		l.WithFields(driver.Fields{
			"storagedriver": "sqlite",
			"logdriver":     "standard",
			"version":       version.Version,
		}).Fatal(err)
	}
	defer s.End()

	if s.FindConfig("authorization") {
		// Validate Docker command is allowed
		r = allow.AllowAction(&config, u.Action, u.CmdName)
		if r.Allow {
			r = u.AllowFunc(req, &config)
		}
	}

	// Build Docker command from data sent to Docker daemon
	lmsg := u.DCBFunc(req, uriinfo.Path, u.Re)

	// Log event to syslog
	l.WithFields(driver.Fields{
		"user":    user,
		"host":    host,
		"allowed": r.Allow,
	}).Info(lmsg)

	// If Docker command is not allowed, return
	if !r.Allow {
		return r
	}

	return &types.AllowResult{Allow: true}
}
