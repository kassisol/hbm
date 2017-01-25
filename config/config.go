package config

type Config struct {
	Action      string
	Description string
}

type Configs []Config

func New() Configs {
	config := []Config{}

	config = append(config, Config{Action: "authorization", Description: "Authorization"})

	return config
}

func (c Configs) ConfigExists(name string) bool {
	for _, config := range c {
		if config.Action == name {
			return true
		}
	}

	return false
}
