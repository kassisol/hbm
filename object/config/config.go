package config

import (
	"strconv"

	configpkg "github.com/kassisol/hbm/config"
	"github.com/kassisol/hbm/object/types"
	"github.com/kassisol/hbm/storage"
	"github.com/kassisol/hbm/storage/driver"
)

type Conf interface {
	Get(name string) (bool, error)
	Set(name, value string) error
	List(filters map[string]string) ([]types.Config, error)
	End()
}

type Config struct {
	Storage driver.Storager
}

func New(driver, options string) (Conf, error) {
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

func (c *Config) Get(name string) (bool, error) {
	conf := configpkg.New()
	if err := conf.IsValid(name); err != nil {
		return false, err
	}

	return c.Storage.GetConfig(name), nil
}

func (c *Config) Set(name, value string) error {
	conf := configpkg.New()
	if err := conf.IsValid(name); err != nil {
		return err
	}

	val, err := strconv.ParseBool(value)
	if err != nil {
		return err
	}

	c.Storage.SetConfig(name, val)

	return nil
}

func (c *Config) List(filters map[string]string) ([]types.Config, error) {
	return c.Storage.ListConfigs(filters), nil
}
