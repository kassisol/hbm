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

	config = append(config, Config{Action: "authorization", Description: "Authorization"})

	return config
}

// ConfigExists does config exist?
func (c Configs) ConfigExists(name string) bool {
	for _, config := range c {
		if config.Action == name {
			return true
		}
	}

	return false
}
