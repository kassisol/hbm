package sqlite

func (c *Config) AddHost(name string) {
	c.DB.Create(&Host{Name: name})
}

func (c *Config) RemoveHost(name string) {
	c.DB.Where("name = ?", name).Delete(Host{})
}

func (c *Config) ListHosts() []string {
	var hosts []Host
	var result []string

	c.DB.Find(&hosts)

	for _, host := range hosts {
		result = append(result, host.Name)
	}

	return result
}

func (c *Config) FindHost(name string) bool {
	var count int64

	c.DB.Model(&Host{}).Where("name = ?", name).Count(&count)

	if count == 1 {
		return true
	}

	return false
}

func (c *Config) CountHost() int {
	var count int64

	c.DB.Model(&Host{}).Count(&count)

	return int(count)
}

func (c *Config) AddHostToCluster(cluster, host string) {
	cl := Cluster{}
	h := Host{}

	c.DB.Where("name = ?", host).Find(&h)
	c.DB.Where("name = ?", cluster).Find(&cl)

	c.DB.Model(&cl).Association("Hosts").Append(&h)
}

func (c *Config) RemoveHostFromCluster(cluster, host string) {
	cl := Cluster{}
	h := Host{}

	c.DB.Where("name = ?", host).Find(&h)
	c.DB.Where("name = ?", cluster).Find(&cl)

	c.DB.Model(&cl).Association("Hosts").Delete(&h)
}
