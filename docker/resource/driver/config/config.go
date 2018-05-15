package config

import (
	"fmt"

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
	resource.RegisterDriver("config", New)
}

func New() (driver.Resourcer, error) {
	actions := []Action{}

	actions = append(actions, Action{Key: "container_create_param_ipc_host", Description: "--ipc=\"host\" param"})
	actions = append(actions, Action{Key: "container_create_param_net_host", Description: "--net=\"host\" param"})
	actions = append(actions, Action{Key: "container_create_param_oom_kill_disable", Description: "--oom-kill-disable param"})
	actions = append(actions, Action{Key: "container_create_param_oom_score_adj", Description: "--oom-score-adj param"})
	actions = append(actions, Action{Key: "container_create_param_pid_host", Description: "--pid=\"host\" param"})
	actions = append(actions, Action{Key: "container_create_param_privileged", Description: "--privileged param"})
	actions = append(actions, Action{Key: "container_create_param_publish_all", Description: "--publish-all param"})
	actions = append(actions, Action{Key: "container_create_param_security_opt", Description: "--security-opt param"})
	actions = append(actions, Action{Key: "container_create_param_sysctl", Description: "--sysctl param"})
	actions = append(actions, Action{Key: "container_create_param_tmpfs", Description: "--tmpfs param"})
	actions = append(actions, Action{Key: "container_create_param_user_root", Description: "--user=\"root\" param"})
	actions = append(actions, Action{Key: "container_create_param_userns_host", Description: "--userns=\"host\" param"})
	actions = append(actions, Action{Key: "container_create_param_uts_host", Description: "--uts=\"host\" param"})

	actions = append(actions, Action{Key: "image_create_official", Description: "Pull of Official image"})

	return &Config{Actions: actions}, nil
}

func (c *Config) List() interface{} {
	return c.Actions
}

func (c *Config) Valid(value string) error {
	for _, cfg := range c.Actions {
		if cfg.Key == value {
			return nil
		}
	}

	return fmt.Errorf("Config '%s' is not valid", value)
}

func (c *Config) ValidOptions(options map[string]string) error {
	return nil
}
