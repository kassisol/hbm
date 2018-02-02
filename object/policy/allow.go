package policy

import (
	"strconv"

	"github.com/kassisol/hbm/pkg/utils"
)

func (c *Config) allowPort(username, port string) bool {
	resources := c.Storage.GetResourceValues(username, "port")

	p, err := strconv.Atoi(port)
	if err != nil {
		return false
	}

	for _, r := range resources {
		startPort, endPort, err := utils.GetPortRangeFromString(r.Value)
		if err != nil {
			return false
		}

		if p >= startPort && p <= endPort {
			return true
		}
	}

	return false
}
