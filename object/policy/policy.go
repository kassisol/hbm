package policy

import (
	"fmt"

	"github.com/juliengk/go-utils/validation"
	"github.com/kassisol/hbm/object/types"
	"github.com/kassisol/hbm/storage"
	"github.com/kassisol/hbm/storage/driver"
)

type Policy interface {
	Add(name, group, collection string) error
	Remove(name string) error
	List(filters map[string]string) ([]types.Policy, error)
	Count() int
	Find(name string) bool
	Validate(user, rType, rValue, rOptions string) bool
	End()
}

type Config struct {
	Storage driver.Storager
}

func New(driver, options string) (Policy, error) {
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

func (c *Config) Add(name, group, collection string) error {
	defer c.Storage.End()

	if err := validation.IsValidName(name); err != nil {
		return err
	}

	if len(group) == 0 {
		return fmt.Errorf("group is MANDATORY")
	}

	if len(collection) == 0 {
		return fmt.Errorf("collection is MANDATORY")
	}

	if c.Storage.FindPolicy(name) {
		return fmt.Errorf("policy '%s' already exists", name)
	}

	if !c.Storage.FindGroup(group) {
		return fmt.Errorf("group '%s' does not exist", group)
	}

	if !c.Storage.FindCollection(collection) {
		return fmt.Errorf("collection '%s' does not exist", collection)
	}

	c.Storage.AddPolicy(name, group, collection)

	return nil
}

func (c *Config) Remove(name string) error {
	if !c.Storage.FindPolicy(name) {
		return fmt.Errorf("policy '%s' does not exist", name)
	}

	c.Storage.RemovePolicy(name)

	return nil
}

func (c *Config) List(filters map[string]string) ([]types.Policy, error) {
	if err := isValidFilterKeys(filters); err != nil {
		return []types.Policy{}, err
	}

	return c.Storage.ListPolicies(filters), nil
}

func (c *Config) Count() int {
	return c.Storage.CountPolicy()
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

func (c *Config) Validate(user, rType, rValue, rOptions string) bool {
	var filters []map[string]string

	filters = append(filters, map[string]string{
		"user":             "all",
		"resource-type":    "all",
		"resource-value":   "all",
		"resource-options": "all",
	})

	filters = append(filters, map[string]string{
		"user":             "all",
		"resource-type":    "all",
		"resource-value":   "all",
		"resource-options": "",
	})

	filters = append(filters, map[string]string{
		"user":             "all",
		"resource-type":    rType,
		"resource-value":   "all",
		"resource-options": "all",
	})

	filters = append(filters, map[string]string{
		"user":             "all",
		"resource-type":    rType,
		"resource-value":   "all",
		"resource-options": "",
	})

	filters = append(filters, map[string]string{
		"user":             "all",
		"resource-type":    rType,
		"resource-value":   rValue,
		"resource-options": rOptions,
	})

	filters = append(filters, map[string]string{
		"user":             user,
		"resource-type":    rType,
		"resource-value":   "all",
		"resource-options": "all",
	})

	filters = append(filters, map[string]string{
		"user":             user,
		"resource-type":    rType,
		"resource-value":   "all",
		"resource-options": "",
	})

	filters = append(filters, map[string]string{
		"user":             user,
		"resource-type":    rType,
		"resource-value":   rValue,
		"resource-options": rOptions,
	})

	for _, filter := range filters {
		result, err := c.List(filter)
		if err != nil {
			return false
		}

		if len(result) > 0 {
			return true
		}
	}

	if rType == "port" {
		return c.allowPort(user, rValue) || c.allowPort("all", rValue)
	}

	return false
}
