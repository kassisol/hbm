package config

type Config struct {
	Action		string
	Description	string
}

func New() []Config {
	config := []Config{}

	config = append(config, Config{Action: "container_create_privileged", Description: "--privileged param is not allowed"})
	config = append(config, Config{Action: "container_create_ipc_host", Description: "--ipc=\"host\" param is not allowed"})
	config = append(config, Config{Action: "container_create_net_host", Description: "--net=\"host\" param is not allowed"})
	config = append(config, Config{Action: "container_create_pid_host", Description: "--pid=\"host\" param is not allowed"})
	config = append(config, Config{Action: "container_create_userns_host", Description: "--userns=\"host\" param is not allowed"})
	config = append(config, Config{Action: "container_create_uts_host", Description: "--uts=\"host\" param is not allowed"})

	config = append(config, Config{Action: "image_create_official", Description: "pull of Official image is allowed"})

	return config
}
