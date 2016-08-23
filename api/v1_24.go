package api

import (
	"github.com/harbourmaster/hbm/dcb"
	"github.com/harbourmaster/hbm/pkg/uri"
)

func register1_24(uris *uri.URIs) {
	// Deprecated container copy
	uris.UnRegister("POST", `^/containers/(.+)/copy`, AllowTrue, dcb.ContainerCopy, "container_copy", "cp", "Copy files/folders between a container and the local filesystem")

	// New swarm endpoints
	uris.Register("GET", `^/nodes`, AllowTrue, dcb.NodesList, "nodes_list", "node ls", "List all nodes")
	uris.Register("GET", `^/nodes/(.+)`, AllowTrue, dcb.NodesInspect, "nodes_inspect", "node inspect", "Inspect node")
	uris.Register("DELETE", `^/nodes/(.+)`, AllowTrue, dcb.NodesDelete, "nodes_delete", "node rm", "Delete node")
	uris.Register("POST", `^/nodes/(.+)/update`, AllowTrue, dcb.NodesUpdate, "nodes_update", "node accept", "Update node")

	uris.Register("POST", `^/swarm`, AllowTrue, dcb.SwarmInfo, "swarm_info", "swarm info", "Get swarm info")
	uris.Register("POST", `^/swarm/init`, AllowTrue, dcb.SwarmInit, "swarm_init", "swarm init", "Initialize new swarm")
	uris.Register("POST", `^/swarm/join`, AllowTrue, dcb.SwarmJoin, "swarm_join", "swarm join", "Join swarm")
	uris.Register("POST", `^/swarm/leave`, AllowTrue, dcb.SwarmLeave, "swarm_leave", "swarm leave", "Leave swarm")
	uris.Register("POST", `^/swarm/update`, AllowTrue, dcb.SwarmUpdate, "swarm_update", "swarm update", "Update swarm")

	uris.Register("GET", `^/services`, AllowTrue, dcb.ServicesList, "services_list", "service ls", "List all services")
	uris.Register("GET", `^/services/(.+)`, AllowTrue, dcb.ServicesInspect, "services_inspect", "service inspect", "Inspect service")
	uris.Register("DELETE", `^/services/(.+)`, AllowTrue, dcb.ServicesDelete, "services_delete", "service rm", "Delete service")
	uris.Register("POST", `^/services/create`, AllowTrue, dcb.ServicesCreate, "services_create", "service create", "Create service")
	uris.Register("POST", `^/services/(.+)/update`, AllowTrue, dcb.ServicesUpdate, "services_update", "service update", "Update service")

	uris.Register("GET", `^/tasks`, AllowTrue, dcb.TasksLists, "tasks_list", "service tasks", "List all tasks")
	uris.Register("GET", `^/tasks/(.+)`, AllowTrue, dcb.TasksInspect, "tasks_inspect", "service tasks", "Inspect task")
}
