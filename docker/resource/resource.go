package resource

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kassisol/hbm/docker/resource/driver"
)

type Initialize func() (driver.Resourcer, error)

var initializers = make(map[string]Initialize)

func NewDriver(driver string) (driver.Resourcer, error) {
	if init, exists := initializers[driver]; exists {
		return init()
	}

	return nil, fmt.Errorf("The Resource Driver: %s is not supported. Supported drivers are %s", driver, SupportedDrivers(","))
}

func RegisterDriver(driver string, init Initialize) {
	initializers[driver] = init
}

func SupportedDrivers(sep string) string {
	drivers := make([]string, 0, len(initializers))

	for d := range initializers {
		drivers = append(drivers, string(d))
	}

	sort.Strings(drivers)

	return strings.Join(drivers, sep)
}
