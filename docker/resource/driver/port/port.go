package port

import (
	"strconv"

	"github.com/juliengk/go-utils/validation"
	"github.com/kassisol/hbm/docker/resource"
	"github.com/kassisol/hbm/docker/resource/driver"
)

type Config struct {
}

func init() {
	resource.RegisterDriver("port", New)
}

func New() (driver.Resourcer, error) {
	return &Config{}, nil
}

func (c *Config) List() interface{} {
	return []string{}
}

func (c *Config) Valid(value string) error {
	port, err := strconv.Atoi(value)
	if err != nil {
		return err
	}

	return validation.IsValidPort(port)
}

func (c *Config) ValidOptions(options map[string]string) error {
	return nil
}
