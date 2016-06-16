package api

import (
	"fmt"
	"log"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/harbourmaster/hbm/api/types"
	"github.com/harbourmaster/hbm/pkg/db"
	"github.com/harbourmaster/hbm/pkg/uri"
	"github.com/harbourmaster/hbm/pkg/utils"
)

var SupportedVersion = "v1.23"

type Api struct {
	Uris	*uri.URIs
	AppPath	string
}

func NewApi(version, appPath string) (*Api, error) {
	if version != SupportedVersion {
		return &Api{}, fmt.Errorf("This version of Harbourmaster does not support Docker API version %s. Supported version is %s", version, SupportedVersion)
	}

	uris := uri.New()

	uris.Register("GET", `^/containers/json`, AllowTrue, "container_list", "ps")
	uris.Register("POST", `^/containers/create`, AllowContainerCreate, "container_create", "run")
	uris.Register("GET", `^/containers/(.*)/json`, AllowTrue, "container_inspect", "inspect")
	uris.Register("GET", `^/containers/(.*)/top`, AllowTrue, "container_top", "top")
	uris.Register("GET", `^/containers/(.*)/logs`, AllowTrue, "container_logs", "logs")
	uris.Register("GET", `^/containers/(.*)/changes`, AllowTrue, "container_changes", "events")
	uris.Register("GET", `^/containers/(.*)/export`, AllowTrue, "container_export", "export")
	uris.Register("GET", `^/containers/(.*)/stats`, AllowTrue, "container_stats", "stats")
	uris.Register("GET", `^/containers/(.*)/resize`, AllowTrue, "container_resize", "resize")
	uris.Register("GET", `^/containers/(.*)/start`, AllowTrue, "container_start", "start")
	uris.Register("GET", `^/containers/(.*)/stop`, AllowTrue, "container_stop", "stop")
	uris.Register("GET", `^/containers/(.*)/restart`, AllowTrue, "container_restart", "restart")
	uris.Register("GET", `^/containers/(.*)/kill`, AllowTrue, "container_kill", "kill")
	uris.Register("GET", `^/containers/(.*)/update`, AllowTrue, "container_update", "update")
	uris.Register("GET", `^/containers/(.*)/rename`, AllowTrue, "container_rename", "rename")
	uris.Register("GET", `^/containers/(.*)/pause`, AllowTrue, "container_pause", "pause")
	uris.Register("GET", `^/containers/(.*)/unpause`, AllowTrue, "container_unpause", "unpause")
	uris.Register("GET", `^/containers/(.*)/attach/ws`, AllowTrue, "container_attach_ws", "attach_ws")
	uris.Register("GET", `^/containers/(.*)/attach`, AllowTrue, "container_attach", "attach")
	uris.Register("GET", `^/containers/(.*)/wait`, AllowTrue, "container_wait", "wait")
	uris.Register("DELETE", `^/containers/(.*)`, AllowTrue, "container_remove", "rm")
	uris.Register("POST", `^/containers/(.*)/copy`, AllowTrue, "container_copy", "cp")
	uris.Register("HEAD", `^/containers/(.*)/archive`, AllowTrue, "container_archive", "archive")
	uris.Register("GET", `^/containers/(.*)/archive`, AllowTrue, "container_archive", "archive")
	uris.Register("PUT", `^/containers/(.*)/archive`, AllowTrue, "container_archive", "archive")
	uris.Register("POST", `^/containers/(.*)/exec`, AllowTrue, "container_exec_create", "exec")

	uris.Register("POST", `^/exec/(.*)/start`, AllowTrue, "exec_start", "cp")
	uris.Register("POST", `^/exec/(.*)/resize`, AllowTrue, "exec_resize", "cp")
	uris.Register("GET", `^/exec/(.*)/json`, AllowTrue, "exec_inspect", "cp")

	uris.Register("GET", `^/images/json`, AllowTrue, "image_list", "images")
	uris.Register("POST", `^/build`, AllowTrue, "image_build", "build")
	uris.Register("POST", `^/images/create`, AllowImageCreate, "image_create", "pull")
	uris.Register("GET", `^/images/(.*)/json`, AllowTrue, "image_inspect", "inspect")
	uris.Register("GET", `^/images/(.*)/history`, AllowTrue, "image_history", "history")
	uris.Register("POST", `^/images/(.*)/push`, AllowTrue, "image_push", "push")
	uris.Register("POST", `^/images/(.*)/tag`, AllowTrue, "image_tag", "tag")
	uris.Register("DELETE", `^/images/(.*)`, AllowTrue, "image_remove", "rmi")
	uris.Register("GET", `^/images/search`, AllowTrue, "image_search", "search")
	uris.Register("GET", `^/images/(.*)/get`, AllowTrue, "image_save_image", "save")
	uris.Register("GET", `^/images/get`, AllowTrue, "image_save_images", "save")
	uris.Register("POST", `^/images/load`, AllowTrue, "image_load", "load")

	uris.Register("POST", `^/auth`, AllowTrue, "auth", "login")
	uris.Register("GET", `^/info`, AllowTrue, "info", "info")
	uris.Register("GET", `^/version`, AllowTrue, "version", "version")
	uris.Register("GET", `^/_ping`, AllowTrue, "ping", "ping")
	uris.Register("POST", `^/commit`, AllowTrue, "commit", "commit")
	uris.Register("GET", `^/events`, AllowTrue, "events", "events")

	uris.Register("GET", `^/volumes`, AllowTrue, "volume_list", "volume ls")
	uris.Register("POST", `^/volumes/create`, AllowTrue, "volume_create", "volume create")
	uris.Register("GET", `^/volumes/(.*)`, AllowTrue, "volume_inspect", "volume inspect")
	uris.Register("DELETE", `^/volumes/(.*)`, AllowTrue, "volume_remove", "volume rm")

	uris.Register("GET", `^/networks`, AllowTrue, "network_list", "network ls")
	uris.Register("GET", `^/networks/(.*)`, AllowTrue, "network_inspect", "network inspect")
	uris.Register("POST", `^/networks/create`, AllowTrue, "network_create", "network create")
	uris.Register("POST", `^/networks/(.*)/connect`, AllowTrue, "network_connect", "network connect")
	uris.Register("POST", `^/networks/(.*)/disconnect`, AllowTrue, "network_disconnect", "network disconnect")
	uris.Register("DELETE", `^/networks/(.*)`, AllowTrue, "network_remove", "network rm")

	return &Api{Uris: uris, AppPath: appPath}, nil
}

func (a *Api) Allow(req authorization.Request) (bool, string, string) {
	_, urlPath := utils.GetURIInfo(req)

	defer db.RecoverFunc()

	d, err := db.NewDB(a.AppPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, u := range *a.Uris {
		if req.RequestMethod == u.Method {
			re := u.Re
			if re.MatchString(urlPath) {
				if ! d.KeyExists("action", u.Action) {
					return false, "", fmt.Sprintf("%s is not allowed", u.CmdName)
				}
				d.Conn.Close()

				config := types.Config{AppPath: a.AppPath}

				msg, err := u.Func(req, &config)
				if err != "" {
					return false, err, ""
				}
				if msg != "" {
					return false, "", msg
				}
			}
		}
	}

	return true, "", ""
}
