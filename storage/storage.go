package storage

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kassisol/hbm/storage/driver"
)

// Initialize structure
type Initialize func(string) (driver.Storager, error)

var initializers = make(map[string]Initialize)

func supportedDriver() string {
	drivers := make([]string, 0, len(initializers))

	for d := range initializers {
		drivers = append(drivers, string(d))
	}

	sort.Strings(drivers)

	return strings.Join(drivers, ",")
}

// NewDriver function
func NewDriver(driver, config string) (driver.Storager, error) {
	if init, exists := initializers[driver]; exists {
		return init(config)
	}

	return nil, fmt.Errorf("The Storage Driver: %s is not supported. Supported drivers are %s", driver, supportedDriver())
}

// RegisterDriver function
func RegisterDriver(driver string, init Initialize) {
	initializers[driver] = init
}
