package sqlite

import (
	"fmt"
)

func (c *Config) AddCollection(name string) {
	c.DB.Create(&Collection{Name: name})
}

func (c *Config) RemoveCollection(name string) error {
	if c.collectionUsedInPolicy(name) {
		return fmt.Errorf("collection \"%s\" cannot be removed. It is being used by a policy", name)
	}

	c.DB.Where("name = ?", name).Delete(Collection{})

	return nil
}

func (c *Config) ListCollections() map[string][]string {
	var collections []Collection
	var resources []Resource

	result := make(map[string][]string)

	c.DB.Find(&collections)

	for _, collection := range collections {
		c.DB.Model(collection).Related(&resources, "Resources")

		result[collection.Name] = []string{}

		if len(resources) > 0 {
			for _, resource := range resources {
				result[collection.Name] = append(result[collection.Name], resource.Name)
			}
		}
	}

	return result
}

func (c *Config) FindCollection(name string) bool {
	var count int64

	c.DB.Model(&Collection{}).Where("name = ?", name).Count(&count)

	if count == 1 {
		return true
	}

	return false
}

func (c *Config) CountCollection() int {
	var count int64

	c.DB.Model(&Collection{}).Count(&count)

	return int(count)
}

func (c *Config) collectionUsedInPolicy(name string) bool {
	var count int64

	c.DB.Table("policies").Joins("JOIN collections ON collections.id = policies.collection_id").Where("collections.name = ?", name).Count(&count)

	if count > 0 {
		return true
	}

	return false
}
