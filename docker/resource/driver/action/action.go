package action

import (
	"fmt"

	"github.com/kassisol/hbm/docker/endpoint"
	"github.com/kassisol/hbm/docker/resource"
	"github.com/kassisol/hbm/docker/resource/driver"
)

type Action struct {
	Key         string
	Description string
}

type Config struct {
	Actions []Action
}

func init() {
	resource.RegisterDriver("action", New)
}

func New() (driver.Resourcer, error) {
	actions := []Action{}
	uris := endpoint.GetUris()

	for _, uri := range *uris {
		actions = append(actions, Action{Key: uri.Action, Description: uri.Description})
	}

	return &Config{Actions: actions}, nil
}

func (c *Config) List() interface{} {
	return c.Actions
}

func (c *Config) Valid(value string) error {
	for _, a := range c.Actions {
		if a.Key == value {
			return nil
		}
	}

	return fmt.Errorf("Action '%s' is not valid", value)
}

func (c *Config) ValidOptions(options map[string]string) error {
	return nil
}
