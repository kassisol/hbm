package user

import (
	"fmt"

	"github.com/juliengk/go-utils/validation"
	"github.com/kassisol/hbm/storage"
	"github.com/kassisol/hbm/storage/driver"
)

type User interface {
	Add(name string) error
	Remove(name string) error
	List(filters map[string]string) (map[string][]string, error)
	Count() int
	Find(name string) bool
	AddToGroup(user, group string) error
	RemoveFromGroup(user, group string) error
	End()
}

type Config struct {
	Storage driver.Storager
}

func New(driver, options string) (User, error) {
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

	if c.Storage.FindUser(name) {
		return fmt.Errorf("user '%s' already exists", name)
	}

	c.Storage.AddUser(name)

	return nil
}

func (c *Config) Remove(name string) error {
	if !c.Storage.FindUser(name) {
		return fmt.Errorf("user '%s' does not exist", name)
	}

	if len(c.Storage.ListGroups(map[string]string{"elem": name})) > 0 {
		return fmt.Errorf("user '%s' cannot be removed. It is used by at least one group", name)
	}

	if err := c.Storage.RemoveUser(name); err != nil {
		return err
	}

	return nil
}

func (c *Config) List(filters map[string]string) (map[string][]string, error) {
	return c.Storage.ListUsers(filters), nil
}

func (c *Config) Count() int {
	return c.Storage.CountUser()
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

func (c *Config) AddToGroup(user, group string) error {
	if !c.Storage.FindUser(user) {
		return fmt.Errorf("user '%s' does not exist", user)
	}

	if !c.Storage.FindGroup(group) {
		return fmt.Errorf("group '%s' does not exist", group)
	}

	c.Storage.AddUserToGroup(group, user)

	return nil
}

func (c *Config) RemoveFromGroup(user, group string) error {
	if !c.Storage.FindUser(user) {
		return fmt.Errorf("user '%s' does not exist", user)
	}

	if !c.Storage.FindGroup(group) {
		return fmt.Errorf("group '%s' does not exist", group)
	}

	c.Storage.RemoveUserFromGroup(group, user)

	return nil
}
