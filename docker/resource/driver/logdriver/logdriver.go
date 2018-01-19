package logdriver

import (
	"fmt"

	"github.com/kassisol/hbm/docker/resource"
	"github.com/kassisol/hbm/docker/resource/driver"
)

type Config struct {
	Drivers []string
}

func init() {
	resource.RegisterDriver("logdriver", New)
}

func New() (driver.Resourcer, error) {
	drivers := []string{
		"none",
		"json-file",
		"syslog",
		"journald",
		"gelf",
		"fluentd",
		"awslogs",
		"splunk",
		"etwlogs",
		"gcplogs",
	}

	return &Config{Drivers: drivers}, nil
}

func (c *Config) List() interface{} {
	return c.Drivers
}

func (c *Config) Valid(value string) error {
	if value == "*" {
		return nil
	}

	for _, logdriver := range c.Drivers {
		if logdriver == value {
			return nil
		}
	}

	return fmt.Errorf("Log driver '%s' is not valid", value)
}

func (c *Config) ValidOptions(options map[string]string) error {
	return nil
}
