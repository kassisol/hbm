package sqlite

import (
	"github.com/kassisol/hbm/storage/driver"
)

func (c *Config) SetConfig(name string, value bool) {
	r := c.ListConfigs(map[string]string{"name": name})

	if len(r) == 0 {
		conf := AppConfig{
			Key:   name,
			Value: value,
		}

		c.DB.Create(&conf)
	} else {
		conf := AppConfig{}
		c.DB.Where("key = ?", name).First(&conf)

		conf.Value = value

		c.DB.Save(&conf)
	}
}

func (c *Config) ListConfigs(filter map[string]string) []driver.ConfigResult {
	var configs []AppConfig

	result := []driver.ConfigResult{}

	sql := c.DB

	if v, ok := filter["name"]; ok {
		sql = sql.Where("key = ?", v)
	}

	if v, ok := filter["value"]; ok {
		sql = sql.Where("value = ?", v)
	}

	sql.Find(&configs)

	for _, config := range configs {
		result = append(result, driver.ConfigResult{Key: config.Key, Value: config.Value})
	}

	return result
}

func (c *Config) GetConfig(name string) bool {
	r := c.ListConfigs(map[string]string{"name": name})
	if len(r) == 1 {
		return r[0].Value
	}

	return false
}
