package sqlite

import (
	"fmt"
)

func (c *Config) AddUser(name string) {
	c.DB.Create(&User{Name: name})
}

func (c *Config) RemoveUser(name string) error {
	if c.memberOfGroup(name) {
		return fmt.Errorf("user \"%s\" cannot be removed. It is being used by a group", name)
	}

	c.DB.Where("name = ?", name).Delete(User{})

	return nil
}

func (c *Config) ListUsers(filter map[string]string) map[string][]string {
	result := make(map[string][]string)

	sql := c.DB.Table("users").Select("users.name, groups.name").Joins("LEFT JOIN group_users ON group_users.user_id = users.id").Joins("LEFT JOIN groups ON groups.id = group_users.group_id")

	if v, ok := filter["name"]; ok {
		sql = sql.Where("users.name = ?", v)
	}

	if v, ok := filter["elem"]; ok {
		sql = sql.Where("groups.name = ?", v)
	}

	rows, _ := sql.Rows()
	defer rows.Close()

	for rows.Next() {
		var user string
		var group string

		rows.Scan(&user, &group)

		result[user] = append(result[user], group)
	}

	return result
}

func (c *Config) FindUser(name string) bool {
	var count int64

	c.DB.Model(&User{}).Where("name = ?", name).Count(&count)

	if count == 1 {
		return true
	}

	return false
}

func (c *Config) CountUser() int {
	var count int64

	c.DB.Model(&User{}).Count(&count)

	return int(count)
}

func (c *Config) AddUserToGroup(group, user string) {
	g := Group{}
	u := User{}

	c.DB.Where("name = ?", user).Find(&u)
	c.DB.Where("name = ?", group).Find(&g)

	c.DB.Model(&g).Association("Users").Append(&u)
}

func (c *Config) RemoveUserFromGroup(group, user string) {
	g := Group{}
	u := User{}

	c.DB.Where("name = ?", user).Find(&u)
	c.DB.Where("name = ?", group).Find(&g)

	c.DB.Model(&g).Association("Users").Delete(&u)
}

func (c *Config) memberOfGroup(name string) bool {
	var count int64

	c.DB.Table("clusters").Joins("JOIN cluster_hosts ON cluster_hosts.cluster_id = clusters.id").Joins("JOIN hosts ON hosts.id = cluster_hosts.host_id").Where("hosts.name = ?", name).Count(&count)

	if count > 0 {
		return true
	}

	return false
}
