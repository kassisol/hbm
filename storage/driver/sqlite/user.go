package sqlite

func (c *Config) AddUser(name string) {
	c.DB.Create(&User{Name: name})
}

func (c *Config) RemoveUser(name string) {
	c.DB.Where("name = ?", name).Delete(User{})
}

func (c *Config) ListUsers() map[string][]string {
	var users []User

	result := make(map[string][]string)

	c.DB.Find(&users)

	for _, user := range users {
		result[user.Name] = []string{}

		sql := c.DB.Table("group_users").Select("groups.name").Joins("JOIN groups ON groups.id = group_users.group_id").Where("group_users.user_id = ?", user.ID)

		rows, _ := sql.Rows()
		defer rows.Close()

		for rows.Next() {
			var group string

			rows.Scan(&group)

			result[user.Name] = append(result[user.Name], group)
		}
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

func (c *Config) memberOfGroup(user string) bool {
	return false
}
