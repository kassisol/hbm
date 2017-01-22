package sqlite

import (
	"fmt"
)

func (c *Config) AddGroup(name string) {
	c.DB.Create(&Group{Name: name})
}

func (c *Config) RemoveGroup(name string) error {
	if c.groupUsedInPolicy(name) {
		return fmt.Errorf("group \"%s\" cannot be removed. It is being used by a policy", name)
	}

	c.DB.Where("name = ?", name).Delete(Group{})

	return nil
}

func (c *Config) ListGroups() map[string][]string {
	var groups []Group
	var users []User

	result := make(map[string][]string)

	c.DB.Find(&groups)

	for _, group := range groups {
		c.DB.Model(group).Related(&users, "Users")

		result[group.Name] = []string{}

		if len(users) > 0 {
			for _, user := range users {
				result[group.Name] = append(result[group.Name], user.Name)
			}
		}
	}

	return result
}

func (c *Config) FindGroup(name string) bool {
	var count int64

	c.DB.Model(&Group{}).Where("name = ?", name).Count(&count)

	if count == 1 {
		return true
	}

	return false
}

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
