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

func (c *Config) ListClusters() map[string][]string {
	var clusters []Cluster
	var hosts []Host

	result := make(map[string][]string)

	c.DB.Find(&clusters)

	for _, cluster := range clusters {
		c.DB.Model(cluster).Related(&hosts, "Hosts")

		result[cluster.Name] = []string{}

		if len(hosts) > 0 {
			for _, host := range hosts {
				result[cluster.Name] = append(result[cluster.Name], host.Name)
			}
		}
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
