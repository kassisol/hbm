package dns

import (
	"github.com/kassisol/hbm/docker/resource"
	"github.com/kassisol/hbm/docker/resource/driver"
)

type Config struct {}

func init() {
	resource.RegisterDriver("dns", New)
}

func New() (driver.Resourcer, error) {
	return &Config{}, nil
}

func (c *Config) List() interface{} {
	return []string{}
}

func (c *Config) Valid(value string) error {
	return nil
}

func (c *Config) ValidOptions(options map[string]string) error {
	return nil
}
