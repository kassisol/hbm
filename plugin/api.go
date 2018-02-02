package plugin

import (
	"runtime/debug"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/juliengk/go-log"
	"github.com/juliengk/go-log/driver"
	"github.com/kassisol/hbm/docker/allow"
	"github.com/kassisol/hbm/docker/allow/types"
	"github.com/kassisol/hbm/docker/endpoint"
	configobj "github.com/kassisol/hbm/object/config"
	groupobj "github.com/kassisol/hbm/object/group"
	"github.com/kassisol/hbm/pkg/uri"
	"github.com/kassisol/hbm/version"
)

type Api struct {
	URIInfo *uri.URIInfo
	Uris    *uri.URIs
	AppPath string
}

func NewApi(uriinfo *uri.URIInfo, appPath string) (*Api, error) {
	uris := endpoint.GetUris()

	return &Api{URIInfo: uriinfo, Uris: uris, AppPath: appPath}, nil
}

func (a *Api) Allow(req authorization.Request) (ar *types.AllowResult) {
	l, _ := log.NewDriver("standard", nil)

	s, err := configobj.New("sqlite", a.AppPath)
	if err != nil {
		l.WithFields(driver.Fields{
			"storagedriver": "sqlite",
			"logdriver":     "standard",
			"version":       version.Version,
		}).Fatal(err)
	}
	defer s.End()

	g, err := groupobj.New("sqlite", a.AppPath)
	if err != nil {
		l.WithFields(driver.Fields{
			"storagedriver": "sqlite",
			"logdriver":     "standard",
			"version":       version.Version,
		}).Fatal(err)
	}
	defer g.End()

	defer func() {
		if r := recover(); r != nil {
			l.Warn("Recovered panic: ", r)
			l.Warnf("%s", debug.Stack())

			allow, _ := s.Get("default-allow-action-error")
			err := "an error occurred; contact your system administrator"

			result := types.AllowResult{Allow: allow}
			if !allow {
				result.Error = err
			}

			ar = &result
		}
	}()

	// Authentication
	username := req.User
	if len(username) == 0 {
		username = "root"
	}

	// Authorization
	isAdmin := false

	filters := map[string]string{
		"name": "administrators",
		"elem": username,
	}
	groups, _ := g.List(filters)
	if len(groups) > 0 {
		isAdmin = true
	}

	u, err := a.Uris.GetURI(req.RequestMethod, a.URIInfo.Path)
	if err != nil {
		return &types.AllowResult{Allow: false, Error: err.Error()}
	}

	// Validate Docker command is allowed
	config := types.Config{AppPath: a.AppPath, Username: username}
	r := allow.AllowTrue(req, &config)

	if !isAdmin {
		aR, _ := s.Get("authorization")
		if aR {
			r = allow.AllowAction(&config, u.Action, u.CmdName)
			if r.Allow {
				r = u.AllowFunc(req, &config)
			}
		}
	}

	// Accounting
	// Build Docker command from data sent to Docker daemon
	lmsg := u.DCBFunc(req, a.URIInfo.Path, u.Re)

	// Log event to syslog
	if len(lmsg) > 0 {
		fields := driver.Fields{
			"user":    username,
			"admin":   isAdmin,
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
