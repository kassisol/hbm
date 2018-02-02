package sqlite

import (
	"github.com/kassisol/hbm/object/types"
)

func (c *Config) AddPolicy(name, group, collection string) {
	g := Group{}
	col := Collection{}

	c.DB.Where("name = ?", group).First(&g)
	c.DB.Where("name = ?", collection).First(&col)

	c.DB.Create(&Policy{Name: name, GroupID: g.ID, CollectionID: col.ID})
}

func (c *Config) RemovePolicy(name string) {
	c.DB.Where("name = ?", name).Delete(Policy{})
}

func (c *Config) ListPolicies(filter map[string]string) []types.Policy {
	var policies []types.Policy

	sql := c.DB.Table("policies").Select("policies.name, groups.name, collections.name").Joins("JOIN groups ON groups.id = policies.group_id").Joins("JOIN collections ON collections.id = policies.collection_id")

	if v, ok := filter["user"]; ok {
		sql = sql.Joins("JOIN group_users ON group_users.group_id = groups.id").Joins("JOIN users ON users.id = group_users.user_id").Where("users.name = ?", v)
	}
	if v, ok := filter["group"]; ok {
		sql = sql.Where("groups.name = ?", v)
	}
	rt, rtok := filter["resource-type"]
	rv, rvok := filter["resource-value"]
	if rtok && rvok {
		sql = sql.Joins("JOIN collection_resources ON collection_resources.collection_id = collections.id").Joins("JOIN resources ON resources.id = collection_resources.resource_id").Where("resources.type = ? AND resources.value = ?", rt, rv)
	}
	if v, ok := filter["resource-options"]; ok {
		sql = sql.Where("resources.option = ?", v)
	}
	if v, ok := filter["collection"]; ok {
		sql = sql.Where("collections.name = ?", v)
	}

	rows, _ := sql.Rows()
	defer rows.Close()

	for rows.Next() {
		var policy string
		var group string
		var collection string

		rows.Scan(&policy, &group, &collection)

		policies = append(policies, types.Policy{Name: policy, Group: group, Collection: collection})
	}

	return policies
}

func (c *Config) GetResourceValues(username, rType string) []types.Resource {
	var result []types.Resource

	sql := c.DB.Table("policies").Select("resources.name, resources.type, resources.value, resources.option").Joins("JOIN groups ON groups.id = policies.group_id").Joins("JOIN collections ON collections.id = policies.collection_id")

	sql = sql.Joins("JOIN group_users ON group_users.group_id = groups.id").Joins("JOIN users ON users.id = group_users.user_id").Where("users.name = ?", username)

	sql = sql.Joins("JOIN collection_resources ON collection_resources.collection_id = collections.id").Joins("JOIN resources ON resources.id = collection_resources.resource_id").Where("resources.type = ?", rType)

	rows, _ := sql.Rows()
	defer rows.Close()

	for rows.Next() {
		var rName string
		var rType string
		var rValue string
		var rOption string

		rows.Scan(&rName, &rType, &rValue, &rOption)

		result = append(result, types.Resource{Name: rName, Type: rType, Value: rValue, Option: rOption})
	}

	return result
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
