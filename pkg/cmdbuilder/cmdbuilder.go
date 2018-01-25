package cmdbuilder

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type Config struct {
	Params url.Values
	Cmd    []string
}

func New(cmd string) *Config {
	return &Config{Cmd: []string{"docker", cmd}}
}

func (c *Config) GetParams(r string) {
	u, err := url.ParseRequestURI(r)
	if err == nil {
		c.Params = u.Query()
	}
}

func (c *Config) Add(t string) {
	c.Cmd = append(c.Cmd, t)
}

func (c *Config) AddFilters() {
	if len(c.Params) > 0 {
		if _, ok := c.Params["filters"]; ok {
			compose := false
			if strings.Contains(c.Params["filters"][0], "[") {
				compose = true
			}

			if compose {
				var v map[string][]string

				err := json.Unmarshal([]byte(c.Params["filters"][0]), &v)
				if err != nil {
					panic(err)
				}

				for k, val := range v {
					for _, ka := range val {
						c.Add(fmt.Sprintf("--filter \"%s=%s\"", k, ka))
					}
				}
			} else {
				var v map[string]map[string]bool

				err := json.Unmarshal([]byte(c.Params["filters"][0]), &v)
				if err != nil {
					panic(err)
				}

				for k, val := range v {
					for ka := range val {
						c.Add(fmt.Sprintf("--filter \"%s=%s\"", k, ka))
					}
				}
			}
		}
	}
}

func (c *Config) GetParamAndAdd(k, p string, b bool) {
	if len(c.Params) > 0 {
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
}

func (c *Config) String() string {
	return strings.Join(c.Cmd, " ")
}
