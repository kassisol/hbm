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
		return &Api{}, fmt.Errorf("This version of HBM does not support Docker API version %s. Supported version is %s", version, SupportedVersion)
	}

	uris := uri.New()

	uris.Register("GET", `^/containers/json`, AllowTrue, "container_list", "ps", "List containers")
	uris.Register("POST", `^/containers/create`, AllowContainerCreate, "container_create", "create", "Create a new container")
	uris.Register("GET", `^/containers/(.+)/json`, AllowTrue, "container_inspect", "inspect", "Return low-level information on a container or image")
	uris.Register("GET", `^/containers/(.+)/top`, AllowTrue, "container_top", "top", "Display the running processes of a container")
	uris.Register("GET", `^/containers/(.+)/logs`, AllowTrue, "container_logs", "logs", "Fetch the logs of a container")
	uris.Register("GET", `^/containers/(.+)/changes`, AllowTrue, "container_changes", "events", "Get real time events from the server")
	uris.Register("GET", `^/containers/(.+)/export`, AllowTrue, "container_export", "export", "Export a container's filesystem as a tar archive")
	uris.Register("GET", `^/containers/(.+)/stats`, AllowTrue, "container_stats", "stats", "Display a live stream of container(s) resource usage statistics")
	uris.Register("POST", `^/containers/(.+)/resize`, AllowTrue, "container_resize", "resize", "Resize a container TTY")
	uris.Register("POST", `^/containers/(.+)/start`, AllowTrue, "container_start", "start", "Start one or more stopped containers")
	uris.Register("POST", `^/containers/(.+)/stop`, AllowTrue, "container_stop", "stop", "Stop a running container")
	uris.Register("POST", `^/containers/(.+)/restart`, AllowTrue, "container_restart", "restart", "Restart a container")
	uris.Register("POST", `^/containers/(.+)/kill`, AllowTrue, "container_kill", "kill", "Kill a running container")
	uris.Register("POST", `^/containers/(.+)/update`, AllowTrue, "container_update", "update", "Update configuration of one or more containers")
	uris.Register("POST", `^/containers/(.+)/rename`, AllowTrue, "container_rename", "rename", "Rename a container")
	uris.Register("POST", `^/containers/(.+)/pause`, AllowTrue, "container_pause", "pause", "Pause all processes within a container")
	uris.Register("POST", `^/containers/(.+)/unpause`, AllowTrue, "container_unpause", "unpause", "Unpause all processes within a container")
	uris.Register("POST", `^/containers/(.+)/attach`, AllowTrue, "container_attach", "attach", "Attach to a running container")
	uris.Register("GET", `^/containers/(.+)/attach/ws`, AllowTrue, "container_attach_ws", "attach_ws", "Attach to a running container (websocket)")
	uris.Register("POST", `^/containers/(.+)/wait`, AllowTrue, "container_wait", "wait", "Block until a container stops, then print its exit code")
	uris.Register("DELETE", `^/containers/(.+)`, AllowTrue, "container_remove", "rm", "Remove one or more containers")
	uris.Register("POST", `^/containers/(.+)/copy`, AllowTrue, "container_copy", "cp", "Copy files/folders between a container and the local filesystem")
	uris.Register("HEAD", `^/containers/(.+)/archive`, AllowTrue, "container_archive_info", "archive", "Retrieving information about files and folders in a container")
	uris.Register("GET", `^/containers/(.+)/archive`, AllowTrue, "container_archive", "archive", "Get an archive of a filesystem resource in a container")
	uris.Register("PUT", `^/containers/(.+)/archive`, AllowTrue, "container_archive_extract", "archive", "Extract an archive of files or folders to a directory in a container")
	uris.Register("POST", `^/containers/(.+)/exec`, AllowTrue, "container_exec_create", "exec", "Run a command in a running container")

	uris.Register("POST", `^/exec/(.+)/start`, AllowTrue, "exec_start", "exec", "Exec Start")
	uris.Register("POST", `^/exec/(.+)/resize`, AllowTrue, "exec_resize", "exec", "Exec Resize")
	uris.Register("GET", `^/exec/(.+)/json`, AllowTrue, "exec_inspect", "exec", "Exec Inspect")

	uris.Register("GET", `^/images/json`, AllowTrue, "image_list", "images", "List images")
	uris.Register("POST", `^/build`, AllowTrue, "image_build", "build", "Build an image from a Dockerfile")
	uris.Register("POST", `^/images/create`, AllowImageCreate, "image_create", "pull", "Pull an image or a repository from a registry")
	uris.Register("GET", `^/images/(.+)/json`, AllowTrue, "image_inspect", "inspect", "Return low-level information on a container or image")
	uris.Register("GET", `^/images/(.+)/history`, AllowTrue, "image_history", "history", "Show the history of an image")
	uris.Register("POST", `^/images/(.+)/push`, AllowTrue, "image_push", "push", "Push an image or a repository to a registry")
	uris.Register("POST", `^/images/(.+)/tag`, AllowTrue, "image_tag", "tag", "Tag an image into a repository")
	uris.Register("DELETE", `^/images/(.+)`, AllowTrue, "image_remove", "rmi", "Remove one or more images")
	uris.Register("GET", `^/images/search`, AllowTrue, "image_search", "search", "Search the Docker Hub for images")
	uris.Register("GET", `^/images/(.+)/get`, AllowTrue, "image_save_image", "save", "Save one or more images to a tar archive")
	uris.Register("GET", `^/images/get`, AllowTrue, "image_save_images", "save", "Save one or more images to a tar archive")
	uris.Register("POST", `^/images/load`, AllowTrue, "image_load", "load", "Load an image from a tar archive or STDIN")

	uris.Register("POST", `^/auth`, AllowTrue, "auth", "login", "Log in to a Docker registry")
	uris.Register("GET", `^/info`, AllowTrue, "info", "info", "Display system-wide information")
	uris.Register("GET", `^/version`, AllowTrue, "version", "version", "Show the Docker version information")
	uris.Register("GET", `^/_ping`, AllowTrue, "ping", "", "Ping the docker server")
	uris.Register("POST", `^/commit`, AllowTrue, "commit", "commit", "Create a new image from a container's changes")
	uris.Register("GET", `^/events`, AllowTrue, "events", "events", "Monitor Docker’s events")

	uris.Register("GET", `^/volumes`, AllowTrue, "volume_list", "volume ls", "List volumes")
	uris.Register("POST", `^/volumes/create`, AllowTrue, "volume_create", "volume create", "Create a volume")
	uris.Register("GET", `^/volumes/(.+)`, AllowTrue, "volume_inspect", "volume inspect", "Return low-level information on a volume")
	uris.Register("DELETE", `^/volumes/(.+)`, AllowTrue, "volume_remove", "volume rm", "Remove a volume")

	uris.Register("GET", `^/networks`, AllowTrue, "network_list", "network ls", "List all networks")
	uris.Register("GET", `^/networks/(.+)`, AllowTrue, "network_inspect", "network inspect", "Display detailed network information")
	uris.Register("POST", `^/networks/create`, AllowTrue, "network_create", "network create", "Create a network")
	uris.Register("POST", `^/networks/(.+)/connect`, AllowTrue, "network_connect", "network connect", "Connect container to a network")
	uris.Register("POST", `^/networks/(.+)/disconnect`, AllowTrue, "network_disconnect", "network disconnect", "Disconnect container from a network")
	uris.Register("DELETE", `^/networks/(.+)`, AllowTrue, "network_remove", "network rm", "Remove a network")

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
					d.Conn.Close()

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
