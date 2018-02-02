package group

import (
	"fmt"

	"github.com/juliengk/go-utils/validation"
	"github.com/kassisol/hbm/storage"
	"github.com/kassisol/hbm/storage/driver"
)

type Group interface {
	Add(name string) error
	Remove(name string) error
	List(filters map[string]string) (map[string][]string, error)
	Count() int
	Find(name string) bool
	End()
}

type Config struct {
	Storage driver.Storager
}

func New(driver, options string) (Group, error) {
	s, err := storage.NewDriver(driver, options)
	if err != nil {
		return new(Config), err
	}

	return &Config{
		Storage: s,
	}, nil
}

func (c *Config) End() {
	c.Storage.End()
}

func (c *Config) Add(name string) error {
	if err := validation.IsValidGroupname(name); err != nil {
		return err
	}

	if c.Storage.FindGroup(name) {
		return fmt.Errorf("%s already exists", name)
	}

	c.Storage.AddGroup(name)

	return nil
}

func (c *Config) Remove(name string) error {
	if name == "administrators" {
		return fmt.Errorf("group \"administrators\" cannot be removed")
	}

	if !c.Storage.FindGroup(name) {
		return fmt.Errorf("%s does not exist", name)
	}

	if err := c.Storage.RemoveGroup(name); err != nil {
		return err
	}

	return nil
}

func (c *Config) List(filters map[string]string) (map[string][]string, error) {
	return c.Storage.ListGroups(filters), nil

}

func (c *Config) Count() int {
	return c.Storage.CountGroup()
}

func (c *Config) Find(name string) bool {
	filters := map[string]string{"name": name}

	result, err := c.List(filters)
	if err != nil {
		return false
	}

	if len(result) > 0 {
		return true
	}

	return false
}
