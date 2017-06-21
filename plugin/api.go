package plugin

import (
	"fmt"

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

// SupportedVersion is the supported Docker API version
var SupportedVersion = "v1.24"

// API structure
type API struct {
	Uris    *uri.URIs
	AppPath string
}

// NewAPI function
func NewAPI(version, appPath string) (*API, error) {
	if version != SupportedVersion {
		return &API{}, fmt.Errorf("This version of HBM does not support Docker API version %s. Supported version is %s", version, SupportedVersion)
	}

	uris := endpoint.GetUris()

	return &API{Uris: uris, AppPath: appPath}, nil
}

// Allow function
func (a *API) Allow(req authorization.Request) *types.AllowResult {
	l, _ := log.NewDriver("standard", nil)

	uriinfo, err := uri.GetURIInfo(SupportedVersion, req)
	if err != nil {
		// Log event
		l.Warning(err)

		return &types.AllowResult{Allow: false, Error: err.Error()}
	}

	user := req.User
	if user == "" {
		user = "root"
	}

	config := types.Config{AppPath: a.AppPath, Username: user}

	u, err := a.Uris.GetURI(req.RequestMethod, uriinfo.Path)
	if err != nil {
		msg := fmt.Sprintf("%s is not implemented", uriinfo.Path)

		// Log event
		l.Warning(msg)

		return &types.AllowResult{Allow: false, Error: msg}
	}

	r := allow.True(req, &config)

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
		r = allow.Action(&config, u.Action, u.CmdName)
		if r.Allow {
			r = u.AllowFunc(req, &config)
		}
	}

	// Build Docker command from data sent to Docker daemon
	lmsg := u.DCBFunc(req, uriinfo.Path, u.Re)

	// Log event to syslog
	if len(lmsg) > 0 {
		l.WithFields(driver.Fields{
			"user":    user,
			"allowed": r.Allow,
		}).Info(lmsg)
	}

	// If Docker command is not allowed, return
	if !r.Allow {
		return r
	}

	return &types.AllowResult{Allow: true}
}
