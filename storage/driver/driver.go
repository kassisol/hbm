package driver

type Storager interface {
	AddConfig(name string)
	RemoveConfig(name string) error
	ListConfigs() []string
	FindConfig(name string) bool

	AddUser(name string)
	RemoveUser(name string) error
	ListUsers(filter map[string]string) map[string][]string
	FindUser(name string) bool
	CountUser() int
	AddUserToGroup(group, user string)
	RemoveUserFromGroup(group, user string)

	AddGroup(name string)
	RemoveGroup(name string) error
	ListGroups(filter map[string]string) map[string][]string
	FindGroup(name string) bool
	CountGroup() int

	AddHost(name string)
	RemoveHost(name string) error
	ListHosts(filter map[string]string) map[string][]string
	FindHost(name string) bool
	CountHost() int
	AddHostToCluster(cluster, host string)
	RemoveHostFromCluster(cluster, host string)

	AddCluster(name string)
	RemoveCluster(name string) error
	ListClusters(filter map[string]string) map[string][]string
	FindCluster(name string) bool
	CountCluster() int

	AddResource(name, rtype, value, options string)
	RemoveResource(name string) error
	ListResources(filter map[string]string) map[ResourceResult][]string
	FindResource(name string) bool
	CountResource(rtype string) int
	AddResourceToCollection(col, res string)
	RemoveResourceFromCollection(col, res string)

	AddCollection(name string)
	RemoveCollection(name string) error
	ListCollections(filter map[string]string) map[string][]string
	FindCollection(name string) bool
	CountCollection() int

	AddPolicy(name, group, cluster, collection string)
	RemovePolicy(name string)
	ListPolicies(filter map[string]string) []PolicyResult
	FindPolicy(name string) bool
	CountPolicy() int

	ValidatePolicy(user, host, res_type, res_value, option string) bool

	End()
}
