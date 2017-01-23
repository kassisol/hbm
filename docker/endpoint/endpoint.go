package endpoint

import (
	"github.com/kassisol/hbm/allow"
	"github.com/kassisol/hbm/docker/dcb"
	"github.com/kassisol/hbm/pkg/uri"
)

func GetUris() *uri.URIs {
	uris := uri.New()

	uris.Register("GET", `^/containers/json`, allow.AllowTrue, dcb.ContainerList, "container_list", "ps", "List containers")
	uris.Register("POST", `^/containers/create`, allow.AllowContainerCreate, dcb.ContainerCreate, "container_create", "create", "Create a new container")
	uris.Register("GET", `^/containers/(.+)/json`, allow.AllowTrue, dcb.ContainerInspect, "container_inspect", "inspect", "Return low-level information on a container or image")
	uris.Register("GET", `^/containers/(.+)/top`, allow.AllowTrue, dcb.ContainerTop, "container_top", "top", "Display the running processes of a container")
	uris.Register("GET", `^/containers/(.+)/logs`, allow.AllowTrue, dcb.ContainerLogs, "container_logs", "logs", "Fetch the logs of a container")
	uris.Register("GET", `^/containers/(.+)/changes`, allow.AllowTrue, dcb.ContainerChanges, "container_changes", "events", "Get real time events from the server")
	uris.Register("GET", `^/containers/(.+)/export`, allow.AllowTrue, dcb.ContainerExport, "container_export", "export", "Export a container's filesystem as a tar archive")
	uris.Register("GET", `^/containers/(.+)/stats`, allow.AllowTrue, dcb.ContainerStats, "container_stats", "stats", "Display a live stream of container(s) resource usage statistics")
	uris.Register("POST", `^/containers/(.+)/resize`, allow.AllowTrue, dcb.ContainerResize, "container_resize", "resize", "Resize a container TTY")
	uris.Register("POST", `^/containers/(.+)/start`, allow.AllowTrue, dcb.ContainerStart, "container_start", "start", "Start one or more stopped containers")
	uris.Register("POST", `^/containers/(.+)/stop`, allow.AllowTrue, dcb.ContainerStop, "container_stop", "stop", "Stop a running container")
	uris.Register("POST", `^/containers/(.+)/restart`, allow.AllowTrue, dcb.ContainerRestart, "container_restart", "restart", "Restart a container")
	uris.Register("POST", `^/containers/(.+)/kill`, allow.AllowTrue, dcb.ContainerKill, "container_kill", "kill", "Kill a running container")
	uris.Register("POST", `^/containers/(.+)/update`, allow.AllowTrue, dcb.ContainerUpdate, "container_update", "update", "Update configuration of one or more containers")
	uris.Register("POST", `^/containers/(.+)/rename`, allow.AllowTrue, dcb.ContainerRename, "container_rename", "rename", "Rename a container")
	uris.Register("POST", `^/containers/(.+)/pause`, allow.AllowTrue, dcb.ContainerPause, "container_pause", "pause", "Pause all processes within a container")
	uris.Register("POST", `^/containers/(.+)/unpause`, allow.AllowTrue, dcb.ContainerUnpause, "container_unpause", "unpause", "Unpause all processes within a container")
	uris.Register("POST", `^/containers/(.+)/attach`, allow.AllowTrue, dcb.ContainerAttach, "container_attach", "attach", "Attach to a running container")
	uris.Register("GET", `^/containers/(.+)/attach/ws`, allow.AllowTrue, dcb.ContainerAttachWS, "container_attach_ws", "attach_ws", "Attach to a running container (websocket)")
	uris.Register("POST", `^/containers/(.+)/wait`, allow.AllowTrue, dcb.ContainerWait, "container_wait", "wait", "Block until a container stops, then print its exit code")
	uris.Register("DELETE", `^/containers/(.+)`, allow.AllowTrue, dcb.ContainerRemove, "container_remove", "rm", "Remove one or more containers")
	uris.Register("POST", `^/containers/(.+)/copy`, allow.AllowTrue, dcb.ContainerCopy, "container_copy", "cp", "Copy files/folders between a container and the local filesystem")
	uris.Register("HEAD", `^/containers/(.+)/archive`, allow.AllowTrue, dcb.ContainerArchiveInfo, "container_archive_info", "archive", "Retrieving information about files and folders in a container")
	uris.Register("GET", `^/containers/(.+)/archive`, allow.AllowTrue, dcb.ContainerArchive, "container_archive", "archive", "Get an archive of a filesystem resource in a container")
	uris.Register("PUT", `^/containers/(.+)/archive`, allow.AllowTrue, dcb.ContainerArchiveExtract, "container_archive_extract", "archive", "Extract an archive of files or folders to a directory in a container")
	uris.Register("POST", `^/containers/(.+)/exec`, allow.AllowTrue, dcb.ContainerExecCreate, "container_exec_create", "exec", "Run a command in a running container")

	uris.Register("POST", `^/exec/(.+)/start`, allow.AllowTrue, dcb.ExecStart, "exec_start", "exec", "Exec Start")
	uris.Register("POST", `^/exec/(.+)/resize`, allow.AllowTrue, dcb.ExecResize, "exec_resize", "exec", "Exec Resize")
	uris.Register("GET", `^/exec/(.+)/json`, allow.AllowTrue, dcb.ExecInspect, "exec_inspect", "exec", "Exec Inspect")

	uris.Register("GET", `^/images/json`, allow.AllowTrue, dcb.ImageList, "image_list", "images", "List images")
	uris.Register("POST", `^/build`, allow.AllowTrue, dcb.ImageBuild, "image_build", "build", "Build an image from a Dockerfile")
	uris.Register("POST", `^/images/create`, allow.AllowImageCreate, dcb.ImageCreate, "image_create", "pull", "Pull an image or a repository from a registry")
	uris.Register("GET", `^/images/(.+)/json`, allow.AllowTrue, dcb.ImageInspect, "image_inspect", "inspect", "Return low-level information on a container or image")
	uris.Register("GET", `^/images/(.+)/history`, allow.AllowTrue, dcb.ImageHistory, "image_history", "history", "Show the history of an image")
	uris.Register("POST", `^/images/(.+)/push`, allow.AllowTrue, dcb.ImagePush, "image_push", "push", "Push an image or a repository to a registry")
	uris.Register("POST", `^/images/(.+)/tag`, allow.AllowTrue, dcb.ImageTag, "image_tag", "tag", "Tag an image into a repository")
	uris.Register("DELETE", `^/images/(.+)`, allow.AllowTrue, dcb.ImageRemove, "image_remove", "rmi", "Remove one or more images")
	uris.Register("GET", `^/images/search`, allow.AllowTrue, dcb.ImageSearch, "image_search", "search", "Search the Docker Hub for images")
	uris.Register("GET", `^/images/(.+)/get`, allow.AllowTrue, dcb.ImageSaveImage, "image_save_image", "save", "Save one or more images to a tar archive")
	uris.Register("GET", `^/images/get`, allow.AllowTrue, dcb.ImageSaveImages, "image_save_images", "save", "Save one or more images to a tar archive")
	uris.Register("POST", `^/images/load`, allow.AllowTrue, dcb.ImageLoad, "image_load", "load", "Load an image from a tar archive or STDIN")

	uris.Register("OPTIONS", `^/(.*)`, allow.AllowTrue, dcb.Anyroute, "anyroute_options", "", "Anyroute OPTIONS")

	uris.Register("POST", `^/auth`, allow.AllowTrue, dcb.Auth, "auth", "login", "Log in to a Docker registry")
	uris.Register("GET", `^/info`, allow.AllowTrue, dcb.Info, "info", "info", "Display system-wide information")
	uris.Register("GET", `^/version`, allow.AllowTrue, dcb.Version, "version", "version", "Show the Docker version information")
	uris.Register("GET", `^/_ping`, allow.AllowTrue, dcb.Ping, "ping", "", "Ping the docker server")
	uris.Register("POST", `^/commit`, allow.AllowTrue, dcb.Commit, "commit", "commit", "Create a new image from a container's changes")
	uris.Register("GET", `^/events`, allow.AllowTrue, dcb.Events, "events", "events", "Monitor Docker’s events")

	uris.Register("GET", `^/volumes$`, allow.AllowTrue, dcb.VolumeList, "volume_list", "volume ls", "List volumes")
	uris.Register("POST", `^/volumes/create`, allow.AllowTrue, dcb.VolumeCreate, "volume_create", "volume create", "Create a volume")
	uris.Register("GET", `^/volumes/(.+)`, allow.AllowTrue, dcb.VolumeInspect, "volume_inspect", "volume inspect", "Return low-level information on a volume")
	uris.Register("DELETE", `^/volumes/(.+)`, allow.AllowTrue, dcb.VolumeRemove, "volume_remove", "volume rm", "Remove a volume")

	uris.Register("GET", `^/networks$`, allow.AllowTrue, dcb.NetworkList, "network_list", "network ls", "List all networks")
	uris.Register("GET", `^/networks/(.+)`, allow.AllowTrue, dcb.NetworkInspect, "network_inspect", "network inspect", "Display detailed network information")
	uris.Register("POST", `^/networks/create`, allow.AllowTrue, dcb.NetworkCreate, "network_create", "network create", "Create a network")
	uris.Register("POST", `^/networks/(.+)/connect`, allow.AllowTrue, dcb.NetworkConnect, "network_connect", "network connect", "Connect container to a network")
	uris.Register("POST", `^/networks/(.+)/disconnect`, allow.AllowTrue, dcb.NetworkDisconnect, "network_disconnect", "network disconnect", "Disconnect container from a network")
	uris.Register("DELETE", `^/networks/(.+)`, allow.AllowTrue, dcb.NetworkRemove, "network_remove", "network rm", "Remove a network")

	return uris
}
