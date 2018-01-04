package driver

import (
	"github.com/kassisol/hbm/object/types"
)

type Storager interface {
	SetConfig(name string, value bool)
	ListConfigs(filter map[string]string) []types.Config
	GetConfig(name string) bool

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

	AddResource(name, rtype, value, options string)
	RemoveResource(name string) error
	ListResources(filter map[string]string) map[types.Resource][]string
	FindResource(name string) bool
	CountResource(rtype string) int
	AddResourceToCollection(col, res string)
	RemoveResourceFromCollection(col, res string)

	AddCollection(name string)
	RemoveCollection(name string) error
	ListCollections(filter map[string]string) map[string][]string
	FindCollection(name string) bool
	CountCollection() int

	AddPolicy(name, group, collection string)
	RemovePolicy(name string)
	ListPolicies(filter map[string]string) []types.Policy
	FindPolicy(name string) bool
	CountPolicy() int

	End()
}
