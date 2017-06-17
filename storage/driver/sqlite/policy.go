package sqlite

import (
	"github.com/kassisol/hbm/storage/driver"
)

// AddPolicy function
func (c *Config) AddPolicy(name, group, collection string) {
	g := Group{}
	col := Collection{}

	c.DB.Where("name = ?", group).First(&g)
	c.DB.Where("name = ?", collection).First(&col)

	c.DB.Create(&Policy{Name: name, GroupID: g.ID, CollectionID: col.ID})
}

// RemovePolicy function
func (c *Config) RemovePolicy(name string) {
	c.DB.Where("name = ?", name).Delete(Policy{})
}

// ListPolicies function
func (c *Config) ListPolicies(filter map[string]string) []driver.PolicyResult {
	var policies []driver.PolicyResult

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

		policies = append(policies, driver.PolicyResult{Name: policy, Group: group, Collection: collection})
	}

	return policies
}

// FindPolicy function
func (c *Config) FindPolicy(name string) bool {
	var count int64

	c.DB.Model(&Policy{}).Where("name = ?", name).Count(&count)

	if count == 1 {
		return true
	}

	return false
}

// CountPolicy function
func (c *Config) CountPolicy() int {
	var count int64

	c.DB.Model(&Policy{}).Count(&count)

	return int(count)
}

// ValidatePolicy function
func (c *Config) ValidatePolicy(user, res_type, res_value, option string) bool {
	rows, _ := c.DB.Raw("SELECT COUNT(*) FROM users, resources, groups, collections, group_users, collection_resources, policies WHERE policies.group_id = groups.id AND policies.collection_id = collections.id AND group_users.group_id = groups.id AND group_users.user_id = users.id AND collection_resources.collection_id = collections.id AND collection_resources.resource_id = resources.id AND users.name = ? AND resources.type = ? AND resources.value = ? AND resources.option = ?", user, res_type, res_value, option).Rows()
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
