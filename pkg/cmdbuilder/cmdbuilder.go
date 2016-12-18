package cmdbuilder

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/kassisol/hbm/pkg/utils"
)

type Config struct {
	Params url.Values
	Cmd    []string
}

func New(cmd string) *Config {
	return &Config{Cmd: []string{"docker", cmd}}
}

func (c *Config) GetParams(r string) {
	c.Params = utils.GetURLParams(r)
}

func (c *Config) Add(t string) {
	c.Cmd = append(c.Cmd, t)
}

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
