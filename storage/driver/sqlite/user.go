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

	c.DB.Table("groups").Joins("JOIN group_users ON group_users.group_id = groups.id").Joins("JOIN users ON users.id = group_users.user_id").Where("users.name = ?", name).Count(&count)

	if count > 0 {
		return true
	}

	return false
}
