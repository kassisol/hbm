package collection

import (
	"fmt"

	"github.com/juliengk/go-utils/validation"
	"github.com/kassisol/hbm/storage"
	"github.com/kassisol/hbm/storage/driver"
)

type Collection interface {
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

func New(driver, options string) (Collection, error) {
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
	if err := validation.IsValidUsername(name); err != nil {
		return err
	}

	if c.Storage.FindCollection(name) {
		return fmt.Errorf("collection '%s' already exists", name)
	}

	c.Storage.AddCollection(name)

	return nil
}

func (c *Config) Remove(name string) error {
	if !c.Storage.FindCollection(name) {
		return fmt.Errorf("collection '%s' does not exist", name)
	}

	if len(c.Storage.ListPolicies(map[string]string{"collection": name})) > 0 {
		return fmt.Errorf("collection '%s' cannot be removed. It is used by a policy", name)
	}

	if len(c.Storage.ListResources(map[string]string{"elem": name})) > 0 {
		return fmt.Errorf("collection '%s' cannot be removed. It is used by at least one resource", name)
	}

	if err := c.Storage.RemoveCollection(name); err != nil {
		return err
	}

	return nil
}

func (c *Config) List(filters map[string]string) (map[string][]string, error) {
	return c.Storage.ListCollections(filters), nil
}

func (c *Config) Count() int {
	return c.Storage.CountCollection()
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
