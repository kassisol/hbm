---
description: The resource add command description and usage
keywords:
- resource, add
menu:
  main:
    parent: smn_cli
title: resource add
---

# hbm resource add
***

```markdown
Add resource to the whitelist

Usage:
  hbm resource add [name] [flags]

Flags:
  -o, --option value   Specify options (default [])
  -t, --type string    Set resource type (action|cap|config|device|dns|image|logdriver|logopt|port|registry|volume) (default "action")
  -v, --value string   Set resource value
```


## Actions

| Action 	 		| Command Name		| Description								|
|:------------------------------|:----------------------|:----------------------------------------------------------------------|
| container_list		| ps           		| List containers							|
| container_top			| top			| Display the running processes of a container				|
| container_logs		| logs			| Fetch the logs of a container						|
| container_changes		| events		| Get real time events from the server					|
| container_export		| export		| Export a container's filesystem as a tar archive			|
| container_stats		| stats			| Display a live stream of container(s) resource usage statistics	|
| container_resize		| resize		| Resize a container TTY						|
| container_start		| start			| Start one or more stopped containers					|
| container_stop		| stop			| Stop a running container						|
| container_restart		| restart		| Restart a container							|
| container_kill		| kill			| Kill a running container						|
| container_update		| update		| Update configuration of one or more containers			|
| container_rename		| rename		| Rename a container							|
| container_pause		| pause			| Pause all processes within a container				|
| container_unpause		| unpause		| Unpause all processes within a container				|
| container_attach		| attach		| Attach to a running container						|
| container_attach_ws		| attach_ws		| Attach to a running container (websocket)				|
| container_wait		| wait			| Block until a container stops, then print its exit code		|
| container_remove		| rm			| Remove one or more containers						|
| container_archive_info	| archive		| Retrieving information about files and folders in a container		|
| container_archive		| archive		| Get an archive of a filesystem resource in a container		|
| container_archive_extract	| archive		| Extract an archive of files or folders to a directory in a container	|
| container_exec_create		| exec			| Run a command in a running container					|
| exec_start			| exec			| Exec Start								|
| exec_resize			| exec			| Exec Resize								|
| exec_inspect			| exec			| Exec Inspect								|
| image_list			| images		| List images								|
| image_build			| build			| Build an image from a Dockerfile					|
| image_create			| pull			| Pull an image or a repository from a registry				|
| image_inspect			| inspect		| Return low-level information on a container or image			|
| image_history			| history		| Show the history of an image						|
| image_push			| push			| Push an image or a repository to a registry				|
| image_tag			| tag			| Tag an image into a repository					|
| image_remove			| rmi			| Remove one or more images						|
| image_search			| search		| Search the Docker Hub for images					|
| image_save_image		| save			| Save one or more images to a tar archive				|
| image_save_images		| save			| Save one or more images to a tar archive				|
| image_load			| load			| Load an image from a tar archive or STDIN				|
| auth				| login			| Log in to a Docker registry						|
| info				| info			| Display system-wide information					|
| version			| version		| Show the Docker version information					|
| ping				| 			| Ping the docker server						|
| events			| events		| Monitor Docker’s events						|
| volume_list			| volume ls		| List volumes								|
| volume_create			| volume create		| Create a volume							|
| volume_inspect		| volume inspect	| Return low-level information on a volume				|
| volume_remove			| volume rm		| Remove a volume							|
| network_list			| network ls		| List all networks							|
| network_inspect		| network inspect	| Display detailed network information					|
| network_create		| network create	| Create a network							|
| network_connect		| network connect	| Connect container to a network					|
| network_disconnect		| network disconnect	| Disconnect container from a network					|
| network_remove		| network rm		| Remove a network							|
| container_create		| create		| Create a new container						|
| container_inspect		| inspect		| Return low-level information on a container or image			|
| commit			| commit		| Create a new image from a container's changes				|
| node_list			| node ls		| List nodes								|
| node_inspect			| node inspect		| Return low-level information on the node id				|
| node_remove			| node rm		| Remove a node [id] from the swarm					|
| node_update			| node update		| Update the node id							|
| swarm_inspect			| swarm info		| Get swarm info							|
| swarm_init			| swarm init		| Initialize a new swarm						|
| swarm_join			| swarm join		| Join an existing swarm						|
| swarm_leave			| swarm leave		| Leave a swarm								|
| swarm_update			| swarm update		| Update a swarm							|
| service_list			| service ls		| List services								|
| service_create		| service create	| Create a service							|
| service_remove		| service rm		| Remove a service							|
| service_inspect		| service inspect	| Return information on the service id					|
| service_update		| service update	| Update a service							|
| task_list			| stask services	| List tasks								|
| task_inspect			| stask tasks		| Get details on a task							|

## Capability

## Config
| Action 	 		| Description			|
|:------------------------------|:------------------------------|
| container_create_privileged	| --privileged param		|
| container_create_ipc_host	| --ipc=\"host\" param		|
| container_create_net_host	| --net=\"host\" param		|
| container_create_pid_host     | --pid=\"host\" param		|
| container_create_userns_host	| --userns=\"host\" param 	|
| container_create_uts_host	| --uts=\"host\" param 		|
| container_create_user_root	| --user=\"root\" param		|
| image_create_official		| Pull of Official image 	|

## Device

## DNS Server

## Image

## Log Driver

## Log Option

## Port

## Registry

## Volume

## Examples

```bash
# hbm resource add --type action --value container_list resource1
# hbm resource ls
NAME                TYPE                VALUE               OPTION              COLLECTIONS
resource1           action              container_list
```

## Related information

* [resource_find](resource_find.md)
* [resource_ls](resource_ls.md)
* [resource_member](resource_member.md)
* [resource_rm](resource_rm.md)
