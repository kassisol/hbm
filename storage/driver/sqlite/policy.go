package sqlite

import (
	"github.com/kassisol/hbm/storage/driver"
)

func (c *Config) AddPolicy(name, group, cluster, collection string) {
	g := Group{}
	cl := Cluster{}
	col := Collection{}

	c.DB.Where("name = ?", group).First(&g)
	c.DB.Where("name = ?", cluster).First(&cl)
	c.DB.Where("name = ?", collection).First(&col)

	c.DB.Create(&Policy{Name: name, GroupID: g.ID, ClusterID: cl.ID, CollectionID: col.ID})
}

func (c *Config) RemovePolicy(name string) {
	c.DB.Where("name = ?", name).Delete(Policy{})
}

func (c *Config) ListPolicies(filter map[string]string) []driver.PolicyResult {
	var policies []driver.PolicyResult

	sql := c.DB.Table("policies").Select("policies.name, groups.name, clusters.name, collections.name").Joins("JOIN groups ON groups.id = policies.group_id").Joins("JOIN clusters ON clusters.id = policies.cluster_id").Joins("JOIN collections ON collections.id = policies.collection_id")

	if v, ok := filter["user"]; ok {
		sql = sql.Joins("JOIN group_users ON group_users.group_id = groups.id").Joins("JOIN users ON users.id = group_users.user_id").Where("users.name = ?", v)
	}
	if v, ok := filter["group"]; ok {
		sql = sql.Where("groups.name = ?", v)
	}
	if v, ok := filter["host"]; ok {
		sql = sql.Joins("JOIN cluster_hosts ON cluster_hosts.cluster_id = clusters.id").Joins("JOIN hosts ON hosts.id = cluster_hosts.host_id").Where("hosts.name = ?", v)
	}
	if v, ok := filter["cluster"]; ok {
		sql = sql.Where("clusters.name = ?", v)
	}
	rt, rtok := filter["resource-type"]
	rv, rvok := filter["resource-value"]
	if rtok && rvok {
		sql = sql.Joins("JOIN collection_resources ON collection_resources.collection_id = collections.id").Joins("JOIN resources ON resources.id = collection_resources.resource_id").Where("resources.type = ? AND resources.value = ?", rt, rv)
	}
	if v, ok := filter["collection"]; ok {
		sql = sql.Where("collections.name = ?", v)
	}

	rows, _ := sql.Rows()
	defer rows.Close()

	for rows.Next() {
		var policy string
		var group string
		var cluster string
		var collection string

		rows.Scan(&policy, &group, &cluster, &collection)

		policies = append(policies, driver.PolicyResult{Name: policy, Group: group, Cluster: cluster, Collection: collection})
	}

	return policies
}

func (c *Config) FindPolicy(name string) bool {
	var count int64

	c.DB.Model(&Policy{}).Where("name = ?", name).Count(&count)

	if count == 1 {
		return true
	}

	return false
}

func (c *Config) CountPolicy() int {
	var count int64

	c.DB.Model(&Policy{}).Count(&count)

	return int(count)
}

func (c *Config) ValidatePolicy(user, host, res_type, res_value, option string) bool {
	rows, _ := c.DB.Raw("SELECT COUNT(*) FROM users, hosts, resources, groups, clusters, collections, group_users, cluster_hosts, collection_resources, policies WHERE policies.group_id = groups.id AND policies.cluster_id = clusters.id AND policies.collection_id = collections.id AND group_users.group_id = groups.id AND group_users.user_id = users.id AND cluster_hosts.cluster_id = clusters.id AND cluster_hosts.host_id = hosts.id AND collection_resources.collection_id = collections.id AND collection_resources.resource_id = resources.id AND users.name = ? AND hosts.name = ? AND resources.type = ? AND resources.value = ? AND resources.option = ?", user, host, res_type, res_value, option).Rows()
	defer rows.Close()

	for rows.Next() {
		var count int

		rows.Scan(&count)
		if count >= 1 {
			return true
		}
	}

	return false
}
