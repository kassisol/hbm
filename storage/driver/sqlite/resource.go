package sqlite

import (
	"fmt"

	"github.com/kassisol/hbm/storage/driver"
)

func (c *Config) AddResource(name, rtype, value, options string) {
	c.DB.Create(&Resource{
		Name:   name,
		Type:   rtype,
		Value:  value,
		Option: options,
	})
}

func (c *Config) RemoveResource(name string) error {
	if c.memberOfCollection(name) {
		return fmt.Errorf("resource \"%s\" cannot be removed. It is being used by a collection", name)
	}

	c.DB.Where("name = ?", name).Delete(Resource{})

	return nil
}

func (c *Config) ListResources() map[driver.ResourceResult][]string {
	var resources []Resource

	result := make(map[driver.ResourceResult][]string)

	c.DB.Find(&resources)

	for _, res := range resources {
		rr := driver.ResourceResult{Name: res.Name, Type: res.Type, Value: res.Value, Option: res.Option}
		result[rr] = []string{}

		sql := c.DB.Table("collection_resources").Select("collections.name").Joins("JOIN collections ON collections.id = collection_resources.collection_id").Where("collection_resources.resource_id = ?", res.ID)

		rows, _ := sql.Rows()
		defer rows.Close()

		for rows.Next() {
			var collection string

			rows.Scan(&collection)

			result[rr] = append(result[rr], collection)
		}
	}

	return result
}

func (c *Config) FindResource(name string) bool {
	var count int64

	c.DB.Model(&Resource{}).Where("name = ?", name).Count(&count)

	if count == 1 {
		return true
	}

	return false
}

func (c *Config) CountResource(rtype string) int {
	var count int64

	if rtype == "all" {
		c.DB.Model(&Resource{}).Count(&count)
	} else {
		c.DB.Model(&Resource{}).Where("type = ?", rtype).Count(&count)
	}

	return int(count)
}

func (c *Config) AddResourceToCollection(collection, resource string) {
	col := Collection{}
	res := Resource{}

	c.DB.Where("name = ?", resource).Find(&res)
	c.DB.Where("name = ?", collection).Find(&col)

	c.DB.Model(&col).Association("Resources").Append(&res)
}

func (c *Config) RemoveResourceFromCollection(collection, resource string) {
	col := Collection{}
	res := Resource{}

	c.DB.Where("name = ?", resource).Find(&res)
	c.DB.Where("name = ?", collection).Find(&col)

	c.DB.Model(&col).Association("Resources").Delete(&res)
}

func (c *Config) memberOfCollection(name string) bool {
	var count int64

	c.DB.Table("collections").Joins("JOIN collection_resources ON collection_resources.collection_id = collections.id").Joins("JOIN resources ON resources.id = collection_resources.resource_id").Where("resources.name = ?", name).Count(&count)

	if count > 0 {
		return true
	}

	return false
}
