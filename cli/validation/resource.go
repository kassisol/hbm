package validation

import (
	"fmt"

	"github.com/juliengk/go-utils"
)

type ResourceTypes map[string]string

func NewResourceTypes() ResourceTypes {
	return map[string]string{
		"action":    "Action",
		"cap":       "Capability",
		"config":    "Config",
		"device":    "Device",
		"dns":       "DNS server",
		"image":     "Image",
		"logdriver": "Log driver",
		"logopt":    "Log option",
		"port":      "Port",
		"registry":  "Registry",
		"volume":    "Volume",
	}
}

func (rt ResourceTypes) IsValidResourceType(name string) error {
	for t := range rt {
		if t == name {
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid resource type", name)
}

func IsValidResourceOptionKeys(options map[string]string) error {
	validKeys := []string{
		"recursive",
		"nosuid",
	}

	if len(options) == 0 {
		return fmt.Errorf("Invalid option")
	}

	for k := range options {
		if !utils.StringInSlice(k, validKeys, false) {
			return fmt.Errorf("%s is not a valid option key", k)
			//fmt.Printf("Conflicting options --type %s and --recursive\n", resourceAddType)
		}
	}

	return nil
}
