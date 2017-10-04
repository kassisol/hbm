package log

import (
	"fmt"
	"sort"
	"strings"

	"github.com/juliengk/go-log/driver"
)

type Initialize func(options interface{}) (driver.Logger, error)

var initializers = make(map[string]Initialize)

func supportedDriver() string {
	drivers := make([]string, 0, len(initializers))

	for d := range initializers {
		drivers = append(drivers, string(d))
	}

	sort.Strings(drivers)

	return strings.Join(drivers, ",")
}

func NewDriver(driver string, options interface{}) (driver.Logger, error) {
	if init, exists := initializers[driver]; exists {
		return init(options)
	}

	return nil, fmt.Errorf("The Logger Driver: %s is not supported. Supported drivers are %s", driver, supportedDriver())
}

func RegisterDriver(driver string, init Initialize) {
	initializers[driver] = init
}
