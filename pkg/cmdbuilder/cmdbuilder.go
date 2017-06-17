package cmdbuilder

import (
	"fmt"
	"net/url"
	"strings"
)

// Config structure
type Config struct {
	Params url.Values
	Cmd    []string
}

// New function
func New(cmd string) *Config {
	return &Config{Cmd: []string{"docker", cmd}}
}

// GetParams function
func (c *Config) GetParams(r string) {
	u, err := url.ParseRequestURI(r)
	if err == nil {
		c.Params = u.Query()
	}
}

// Add function
func (c *Config) Add(t string) {
	c.Cmd = append(c.Cmd, t)
}

// GetParamAndAdd function
func (c *Config) GetParamAndAdd(k, p string, b bool) {
	if b {
		if v, ok := c.Params[k]; ok {
			if v[0] == "1" || v[0] == "True" || v[0] == "true" {
				c.Add(p)
			}
		}
	} else {
		if v, ok := c.Params[k]; ok {
			c.Add(fmt.Sprintf("%s=%s", p, v[0]))
		}
	}
}

func (c *Config) String() string {
	return strings.Join(c.Cmd, " ")
}
