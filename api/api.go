package api

import (
	"fmt"
	"log/syslog"

	log "github.com/Sirupsen/logrus"
	"github.com/davecgh/go-spew/spew"
	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/harbourmaster/hbm/api/types"
	"github.com/harbourmaster/hbm/pkg/db"
	"github.com/harbourmaster/hbm/pkg/uri"
	"github.com/harbourmaster/hbm/pkg/utils"
)

var DefaultVersion = "v1.23"
var SupportedVersions = []string{"v1.23", "v1.24"}

type Api struct {
	Uris    *uri.URIs
	AppPath string
}

func NewApi(version, appPath string) (*Api, error) {
	if !utils.StringInSlice(SupportedVersions, version) {
		return &Api{}, fmt.Errorf("This version of HBM does not support Docker API version %s. Supported versions are %s", version, SupportedVersions)
	}

	uris := uri.New()

	// Development was started on API vesion 1.23 and these are the defaults for that version
	register1_23(uris)

	if version == "v1.24" {
		register1_24(uris)
	}

	// Maybe like this?
	//	if version == "v1.25" {
	//		register1_24(uris)
	//		register1_25(uris)
	//	}

	return &Api{Uris: uris, AppPath: appPath}, nil
}

func initUserBucket(name string, user string, appPath string) {
	bucket := name + "_" + user

	defer db.RecoverFunc()
	d, err := db.NewDB(appPath)
	if err != nil {
		log.Fatal(err)
	}
	d.InitBucket(bucket)
}

func (a *Api) Allow(req authorization.Request) *types.AllowResult {
	_, urlPath := utils.GetURIInfo(req)
	log.Debug(spew.Sdump(req))

	if req.User != "" {
		initUserBucket("action", req.User, a.AppPath)
	}

	defer db.RecoverFunc()
	d, err := db.NewDB(a.AppPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, u := range *a.Uris {
		if req.RequestMethod == u.Method {
			re := u.Re
			if re.MatchString(urlPath) {
				r := &types.AllowResult{Allow: false, Error: fmt.Sprintf("%s is not allowed for anonymous user", u.CmdName)}

				// Validate Docker command is allowed
				if req.User != "" {
					if d.KeyExists("action_"+req.User, u.Action) || d.KeyExists("action", u.Action) {
						log.Debug("Passing on allowed action")
						r = &types.AllowResult{Allow: true}
					} else {
						r = &types.AllowResult{Allow: false, Error: fmt.Sprintf("%s is not allowed for user %s", u.CmdName, req.User)}
					}
				} else {
					if d.KeyExists("action", u.Action) {
						log.Debug("Passing on allowed action")
						r = &types.AllowResult{Allow: true}
					}
				}

				d.Conn.Close()

				if r.Allow {
					config := types.Config{AppPath: a.AppPath}

					r = u.AllowFunc(req, &config)
				}

				// Build Docker command from data sent to Docker daemon
				lmsg := u.DCBFunc(req, u.Re)

				// Log event to syslog
				w, e := syslog.New(syslog.LOG_LOCAL3, "hbm")
				if e != nil {
					log.Fatal(e)
				}
				msg := fmt.Sprintf("%s ; %t", lmsg, r.Allow)
				w.Info(msg)
				w.Close()

				// If Docker command is not allowed, return
				if !r.Allow {
					return r
				}

				break
			}
		}
	}

	return &types.AllowResult{Allow: true}
}
