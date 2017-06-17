package config

// Config structure
type Config struct {
	Action      string
	Description string
}

// Configs array
type Configs []Config

// New configs array
func New() Configs {
	config := []Config{}

	config = append(config, Config{Action: "container_create_privileged", Description: "--privileged param"})
	config = append(config, Config{Action: "container_create_ipc_host", Description: "--ipc=\"host\" param"})
	config = append(config, Config{Action: "container_create_net_host", Description: "--net=\"host\" param"})
	config = append(config, Config{Action: "container_create_pid_host", Description: "--pid=\"host\" param"})
	config = append(config, Config{Action: "container_create_userns_host", Description: "--userns=\"host\" param"})
	config = append(config, Config{Action: "container_create_uts_host", Description: "--uts=\"host\" param"})
	config = append(config, Config{Action: "container_create_user_root", Description: "--user=\"root\" param"})
	config = append(config, Config{Action: "container_publish_all", Description: "--publish-all param"})

	config = append(config, Config{Action: "image_create_official", Description: "Pull of Official image"})

	return config
}

// ActionExists does action exist?
func (c Configs) ActionExists(name string) bool {
	for _, config := range c {
		if config.Action == name {
			return true
		}
	}

	return false
}
