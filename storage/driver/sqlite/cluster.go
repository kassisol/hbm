package sqlite

import (
	"fmt"
)

func (c *Config) AddCluster(name string) {
	c.DB.Create(&Cluster{Name: name})
}

func (c *Config) RemoveCluster(name string) error {
	if c.clusterUsedInPolicy(name) {
		return fmt.Errorf("cluster \"%s\" cannot be removed. It is being used by a policy", name)
	}

	c.DB.Where("name = ?", name).Delete(Cluster{})

	return nil
}

func (c *Config) ListClusters(filter map[string]string) map[string][]string {
	result := make(map[string][]string)

	sql := c.DB.Table("clusters").Select("clusters.name, hosts.name").Joins("LEFT JOIN cluster_hosts ON cluster_hosts.cluster_id = clusters.id").Joins("LEFT JOIN hosts ON hosts.id = cluster_hosts.host_id")

	if v, ok := filter["name"]; ok {
		sql = sql.Where("clusters.name = ?", v)
	}

	if v, ok := filter["elem"]; ok {
		sql = sql.Where("hosts.name = ?", v)
	}

	rows, _ := sql.Rows()
	defer rows.Close()

	for rows.Next() {
		var cluster string
		var host string

		rows.Scan(&cluster, &host)

		result[cluster] = append(result[cluster], host)
	}

	return result
}

func (c *Config) FindCluster(name string) bool {
	var count int64

	c.DB.Model(&Cluster{}).Where("name = ?", name).Count(&count)

	if count == 1 {
		return true
	}

	return false
}

func (c *Config) CountCluster() int {
	var count int64

	c.DB.Model(&Cluster{}).Count(&count)

	return int(count)
}

func (c *Config) clusterUsedInPolicy(name string) bool {
	var count int64

	c.DB.Table("policies").Joins("JOIN clusters ON clusters.id = policies.cluster_id").Where("clusters.name = ?", name).Count(&count)

	if count > 0 {
		return true
	}

	return false
}
