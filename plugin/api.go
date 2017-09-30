package plugin

import (
	"fmt"
	"strings"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/juliengk/go-log"
	"github.com/juliengk/go-log/driver"
	"github.com/juliengk/go-utils"
	"github.com/kassisol/hbm/allow"
	"github.com/kassisol/hbm/allow/types"
	"github.com/kassisol/hbm/docker/endpoint"
	"github.com/kassisol/hbm/pkg/uri"
	"github.com/kassisol/hbm/storage"
	"github.com/kassisol/hbm/version"
)

var SupportedDockerAPIVersions = []string{
	"v1.30",
	"v1.32",
}

type Api struct {
	URIInfo *uri.URIInfo
	Uris    *uri.URIs
	AppPath string
}

func NewApi(uriinfo *uri.URIInfo, appPath string) (*Api, error) {
	if !utils.StringInSlice(uriinfo.Version, SupportedDockerAPIVersions, false) {
		return &Api{}, fmt.Errorf("This version of HBM does not support Docker API version %s. Supported version are %s.", uriinfo.Version, strings.Join(SupportedDockerAPIVersions, ", "))
	}

	uris := endpoint.GetUris()

	return &Api{URIInfo: uriinfo, Uris: uris, AppPath: appPath}, nil
}

func (a *Api) Allow(req authorization.Request) *types.AllowResult {
	l, _ := log.NewDriver("standard", nil)

	// Authentication
	username := req.User
	if len(username) == 0 {
		username = "root"
	}

	// Authorization
	u, err := a.Uris.GetURI(req.RequestMethod, a.URIInfo.Path)
	if err != nil {
		return &types.AllowResult{Allow: false, Error: err.Error()}
	}

	// Validate Docker command is allowed
	config := types.Config{AppPath: a.AppPath, Username: username}
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
		r = allow.AllowAction(&config, u.Action, u.CmdName)
		if r.Allow {
			r = u.AllowFunc(req, &config)
		}
	}

	// Accounting
	// Build Docker command from data sent to Docker daemon
	lmsg := u.DCBFunc(req, a.URIInfo.Path, u.Re)

	// Log event to syslog
	if len(lmsg) > 0 {
		fields := driver.Fields{
			"user":    username,
			"allowed": r.Allow,
		}

		if !r.Allow {
			fields["msg"] = r.Msg
		}

		l.WithFields(fields).Info(lmsg)
	}

	// If Docker command is not allowed, return
	if !r.Allow {
		return r
	}

	return &types.AllowResult{Allow: true}
}
