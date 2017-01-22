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

func (c *Config) ListCollections(filter map[string]string) map[string][]string {
	result := make(map[string][]string)

	sql := c.DB.Table("collections").Select("collections.name, resources.name").Joins("LEFT JOIN collection_resources ON collection_resources.collection_id = collections.id").Joins("LEFT JOIN resources ON resources.id = collection_resources.resource_id")

	if v, ok := filter["name"]; ok {
		sql = sql.Where("collections.name = ?", v)
	}

	if v, ok := filter["elem"]; ok {
		sql = sql.Where("resources.name = ?", v)
	}

	rows, _ := sql.Rows()
	defer rows.Close()

	for rows.Next() {
		var collection string
		var resource string

		rows.Scan(&collection, &resource)

		result[collection] = append(result[collection], resource)
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
