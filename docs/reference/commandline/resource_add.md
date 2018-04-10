---
title: "hbm resource add"
description: "The resource add command description and usage"
keywords: [ "resource", "add" ]
date: "2017-01-27"
menu:
  docs:
    parent: "hbm_cli_resource"
github_edit: "https://github.com/kassisol/hbm/edit/master/docs/reference/commandline/resource_add.md"
toc: true
---

```markdown
Add resource to the whitelist

Usage:
  hbm resource add [name] [flags]

Flags:
  -o, --option value   Specify options (default [])
  -t, --type string    Set resource type (action|capability|config|device|dns|image|logdriver|logopt|plugin|port|registry|volume) (default "action")
  -v, --value string   Set resource value
```


## Resources
### Action
#### Type
`action`

#### Value
| Value                     | Command Name       | Description								|
|:--------------------------|:-------------------|:---------------------------------------------------------------------|
| container_archive         | archive            | Get an archive of a filesystem resource in a container               |
| container_attach_ws       | attach_ws          | Attach to a running container (websocket)                            |
| container_changes         | events             | Get real time events from the server                                 |
| container_export          | export             | Export a container's filesystem as a tar archive                     |
| container_inspect         | container inspect  | Display detailed information on one or more containers               |
| container_logs            | container logs     | Fetch the logs of a container                                        |
| container_stats           | container stats    | Display a live stream of container(s) resource usage statistics      |
| container_top             | container top      | Display the running processes of a container                         |
| container_list            | container ls       | List containers                                                      |
| container_attach          | container attach   | Attach to a running container                                        |
| container_kill            | container kill     | Kill a running container                                             |
| container_pause           | container pause    | Pause all processes within a container                               |
| container_rename          | container rename   | Rename a container                                                   |
| container_resize          | resize             | Resize a container TTY                                               |
| container_restart         | container restart  | Restart a container                                                  |
| container_start           | container start    | Start one or more stopped containers                                 |
| container_stop            | container stop     | Stop a running container                                             |
| container_unpause         | container unpause  | Unpause all processes within a container                             |
| container_update          | container update   | Update configuration of one or more containers                       |
| container_wait            | container wait     | Block until a container stops, then print its exit code              |
| container_create          | container create   | Create a new container                                               |
| container_prune           | container prune    | Remove all stopped containers                                        |
| container_archive_extract | archive            | Extract an archive of files or folders to a directory in a container |
| container_remove          | container rm       | Remove one or more containers                                        |
| container_archive_info    | archive            | Retrieving information about files and folders in a container        |
| image_save_image          | image save         | Save one or more images to a tar archive                             |
| image_history             | image history      | Show the history of an image                                         |
| image_inspect             | image inspect      | Return low-level information on a container or image                 |
| image_save_images         | image save         | Save one or more images to a tar archive                             |
| image_list                | image ls           | List images                                                          |
| image_search              | search             | Search the Docker Hub for images                                     |
| image_build               | image build        | Build an image from a Dockerfile                                     |
| commit                    | commit             | Create a new image from a container's changes                        |
| image_push                | image push         | Push an image or a repository to a registry                          |
| image_tag                 | image tag          | Tag an image into a repository                                       |
| image_create              | image pull         | Pull an image or a repository from a registry                        |
| image_load                | image load         | Load an image from a tar archive or STDIN                            |
| image_prune               | image prune        | Delete unused images                                                 |
| image_remove              | image rm           | Remove one or more images                                            |
| network_list              | network ls         | List all networks                                                    |
| network_inspect           | network inspect    | Display detailed network information                                 |
| network_connect           | network connect    | Connect container to a network                                       |
| network_disconnect        | network disconnect | Disconnect container from a network                                  |
| network_create            | network create     | Create a network                                                     |
| network_prune             | network prune      | Delete unused networks                                               |
| network_remove            | network rm         | Remove a network                                                     |
| volume_list               | volume ls          | List volumes                                                         |
| volume_inspect            | volume inspect     | Return low-level information on a volume                             |
| volume_create             | volume create      | Create a volume                                                      |
| volume_prune              | volume prune       | Delete unused volumes                                                |
| volume_remove             | volume rm          | Remove a volume                                                      |
| exec_inspect              | exec               | Exec Inspect                                                         |
| container_exec_create     | exec               | Run a command in a running container                                 |
| exec_resize               | exec               | Exec Resize                                                          |
| exec_start                | exec               | Exec Start                                                           |
| swarm_unlock_key          | swarm unlock       | Get the unlock key                                                   |
| swarm_inspect             | swarm info         | Get swarm info                                                       |
| swarm_init                | swarm init         | Initialize a new swarm                                               |
| swarm_join                | swarm join         | Join an existing swarm                                               |
| swarm_leave               | swarm leave        | Leave a swarm                                                        |
| swarm_unlock              | swarm unlock       | Unlock a locked manager                                              |
| swarm_update              | swarm update       | Update a swarm                                                       |
| node_inspect              | node inspect       | Return low-level information on the node id                          |
| node_list                 | node ls            | List nodes                                                           |
| node_update               | node update        | Update the node id                                                   |
| node_remove               | node rm            | Remove a node [id] from the swarm                                    |
| service_logs              | service logs       | Get service logs                                                     |
| service_inspect           | service inspect    | Return information on the service id                                 |
| service_list              | service ls         | List services                                                        |
| service_update            | service update     | Update a service                                                     |
| service_create            | service create     | Create a service                                                     |
| service_remove            | service rm         | Remove a service                                                     |
| task_logs                 | task logs          | Get task logs                                                        |
| task_inspect              | stask tasks        | Get details on a task                                                |
| task_list                 | stask services     | List tasks                                                           |
| secret_inspect            | secret inspect     | Inspect a secret                                                     |
| secret_list               | secret ls          | List secrets                                                         |
| secret_update             | secret update      | Update a secret                                                      |
| secret_create             | secret create      | Create a secret                                                      |
| secret_remove             | secret rm          | Delete a secret                                                      |
| plugin_inspect            | plugin inspect     | Inspect a plugin                                                     |
| plugin_privileges         | plugin ls          | Get plugin privileges                                                |
| plugin_list               | plugin ls          | List plugins                                                         |
| plugin_disable            | plugin disable     | Disable a plugin                                                     |
| plugin_enable             | plugin enable      | Enable a plugin                                                      |
| plugin_push               | plugin push        | Push a plugin                                                        |
| plugin_set                | plugin set         | Configure a plugin                                                   |
| plugin_upgrade            | plugin upgrade     | Upgrade a plugin                                                     |
| plugin_create             | plugin create      | Create a plugin                                                      |
| plugin_pull               | plugin pull        | Install a plugin                                                     |
| plugin_remove             | plugin rm          | Remove a plugin                                                      |
| auth                      | login              | Log in to a Docker registry                                          |
| info                      | info               | Display system-wide information                                      |
| version                   | version            | Show the Docker version information                                  |
| events                    | events             | Monitor Docker's events                                              |
| system_df                 | system df          | Get data usage information                                           |
| config_inspect            | config inspect     | Inspect a config                                                     |
| config_list               | config ls          | List configs                                                         |
| config_update             | config update      | Update a config                                                      |
| config_create             | config create      | Create a config                                                      |
| config_remove             | config rm          | Delete a config                                                      |

#### Option
No option available

#### Examples

```bash
# hbm resource add --type action --value container_list resource1
# hbm resource ls -f "type=action"
NAME                TYPE                VALUE               OPTION              COLLECTIONS
resource1           action              container_list
```

---
### Capability
#### Type
`cap`

#### Value
By default Docker drops all capabilities except [those needed](https://github.com/moby/moby/blob/master/oci/defaults.go#L14-L30). You can see a full list of available capabilities in Linux [manpages](http://man7.org/linux/man-pages/man7/capabilities.7.html).

#### Option
No option available

#### Examples

```bash
# hbm resource add --type cap --value SYS_ADMIN resource1
# hbm resource ls -f "type=cap"
NAME                TYPE                VALUE               OPTION              COLLECTIONS
resource1           cap                 SYS_ADMIN
```

---
### Config
#### Type
`config`

#### Value
| Value                               | Description             |
|:------------------------------------|:------------------------|
| container_create_param_privileged   | --privileged param      |
| container_create_param_ipc_host     | --ipc=\"host\" param    |
| container_create_param_net_host     | --net=\"host\" param    |
| container_create_param_pid_host     | --pid=\"host\" param    |
| container_create_param_userns_host  | --userns=\"host\" param |
| container_create_param_uts_host     | --uts=\"host\" param    |
| container_create_param_user_root    | --user=\"root\" param   |
| container_create_param_publish_all  | --publish-all param     |
| container_create_param_security_opt | --security-opt param    |
| container_create_param_sysctl       | --sysctl param          |
| image_create_official               | Pull of Official image 	|

#### Option
No option available

#### Examples

```bash
# hbm resource add --type config --value container_create_param_privileged resource1
# hbm resource ls -f "type=config"
NAME                TYPE                VALUE                                      OPTION              COLLECTIONS
resource1           config              container_create_param_privileged
```

---
### Device
#### Type
`device`

#### Value

#### Option
No option available

#### Examples

```bash
# hbm resource add --type device --value /dev/snd resource1
# hbm resource ls -f "type=device"
NAME                TYPE                VALUE               OPTION              COLLECTIONS
resource1           device              /dev/snd
```

---
### DNS Server
#### Type
`dns`

#### Value

#### Option
No option available

#### Examples

```bash
# hbm resource add --type dns --value 1.1.1.1 resource1
# hbm resource ls -f "type=dns"
NAME                TYPE                VALUE               OPTION              COLLECTIONS
resource1           dns                 1.1.1.1
```

---
### Image
#### Type
`image`

#### Value

#### Option
No option available

#### Examples

```bash
# hbm resource add --type image --value kassisol/hbm resource1
# hbm resource ls -f "type=image"
NAME                TYPE                VALUE               OPTION              COLLECTIONS
resource1           image               kassisol/hbm
```

---
### Log Driver
#### Type
`logdriver`

#### Value

#### Option
No option available

---
### Log Option
#### Type
`logopt`

#### Value

#### Option
No option available

---
### Plugin
#### Type
`plugin`

#### Value

#### Option
No option available

#### Examples

```bash
# hbm resource add --type plugin --value kassisol/gitvol resource1
# hbm resource ls -f "type=plugin"
NAME                TYPE                VALUE               OPTION              COLLECTIONS
resource1           plugin              kassisol/gitvol
```

---
### Port
#### Type
`port`

#### Value
Port can be specify individually or a range of ports is also possible. Range of ports should be separated with a dash character (eg: 8080-8085).

#### Option
No option available

#### Examples

```bash
# hbm resource add --type port --value 80 resource1
# hbm resource add --type port --value 10000-10010 resource2
# hbm resource ls -f "type=port"
NAME                TYPE                VALUE               OPTION              COLLECTIONS
resource1           port                80
resource2           port                10000-10010
```

---
### Registry
#### Type
`registry`

#### Value

#### Option
No option available

#### Examples

```bash
# hbm resource add --type registry --value registry.example.com resource1
# hbm resource ls -f "type=plugin"
NAME                TYPE                VALUE                         OPTION              COLLECTIONS
resource1           registry            registry.example.com
```

---
### Volume
#### Type
`volume`

#### Value
Any path

#### Option
| Key       | Value   | Description |
|:----------|:--------|-------------|
| recursive | boolean |             |
| suid      | boolean |             |

#### Examples

```bash
# hbm resource add --type volume --value /path/to/dir1 resource1
# hbm resource add --type volume --value /path/to --option "recursive=true" resource2
# hbm resource ls -f "type=volume"
NAME                TYPE                VALUE                  OPTION              COLLECTIONS
resource1           volume              /path/to/dir1
```

## Related information

* [resource_find](resource_find.md)
* [resource_ls](resource_ls.md)
* [resource_member](resource_member.md)
* [resource_rm](resource_rm.md)
