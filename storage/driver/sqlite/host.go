package sqlite

import (
	"fmt"
)

func (c *Config) AddHost(name string) {
	c.DB.Create(&Host{Name: name})
}

func (c *Config) RemoveHost(name string) error {
	if c.memberOfCluster(name) {
		return fmt.Errorf("host \"%s\" cannot be removed. It is being used by a cluster", name)
	}

	c.DB.Where("name = ?", name).Delete(Host{})

	return nil
}

func (c *Config) ListHosts(filter map[string]string) map[string][]string {
	result := make(map[string][]string)

	sql := c.DB.Table("hosts").Select("hosts.name, clusters.name").Joins("LEFT JOIN cluster_hosts ON cluster_hosts.host_id = hosts.id").Joins("LEFT JOIN clusters ON clusters.id = cluster_hosts.cluster_id")

	if v, ok := filter["name"]; ok {
		sql = sql.Where("hosts.name = ?", v)
	}

	if v, ok := filter["elem"]; ok {
		sql = sql.Where("clusters.name = ?", v)
	}

	rows, _ := sql.Rows()
	defer rows.Close()

	for rows.Next() {
		var host string
		var cluster string

		rows.Scan(&host, &cluster)

		result[host] = append(result[host], cluster)
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

func (c *Config) memberOfCluster(name string) bool {
	var count int64

	c.DB.Table("clusters").Joins("JOIN cluster_hosts ON cluster_hosts.cluster_id = clusters.id").Joins("JOIN hosts ON hosts.id = cluster_hosts.host_id").Where("hosts.name = ?", name).Count(&count)

	if count > 0 {
		return true
	}

	return false
}
