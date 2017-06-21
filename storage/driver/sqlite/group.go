package sqlite

import (
	"fmt"
)

// AddGroup function
func (c *Config) AddGroup(name string) {
	c.DB.Create(&Group{Name: name})
}

// RemoveGroup function
func (c *Config) RemoveGroup(name string) error {
	if c.groupUsedInPolicy(name) {
		return fmt.Errorf("group \"%s\" cannot be removed. It is being used by a policy", name)
	}

	c.DB.Where("name = ?", name).Delete(Group{})

	return nil
}

// ListGroups function
func (c *Config) ListGroups(filter map[string]string) map[string][]string {
	result := make(map[string][]string)

	sql := c.DB.Table("groups").Select("groups.name, users.name").Joins("LEFT JOIN group_users ON group_users.group_id = groups.id").Joins("LEFT JOIN users ON users.id = group_users.user_id")

	if v, ok := filter["name"]; ok {
		sql = sql.Where("groups.name = ?", v)
	}

	if v, ok := filter["elem"]; ok {
		sql = sql.Where("users.name = ?", v)
	}

	rows, _ := sql.Rows()
	defer rows.Close()

	for rows.Next() {
		var group string
		var user string

		rows.Scan(&group, &user)

		result[group] = append(result[group], user)
	}

	return result
}

// FindGroup function
func (c *Config) FindGroup(name string) bool {
	var count int64

	c.DB.Model(&Group{}).Where("name = ?", name).Count(&count)

	if count == 1 {
		return true
	}

	return false
}

// CountGroup function
func (c *Config) CountGroup() int {
	var count int64

	c.DB.Model(&Group{}).Count(&count)

	return int(count)
}

func (c *Config) groupUsedInPolicy(name string) bool {
	var count int64

	c.DB.Table("policies").Joins("JOIN groups ON groups.id = policies.group_id").Where("groups.name = ?", name).Count(&count)

	if count > 0 {
		return true
	}

	return false
}
