package validation

import (
	"fmt"

	"github.com/juliengk/go-utils"
)

// ResourceTypes map
type ResourceTypes map[string]string

// NewResourceTypes new resource type map
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

// IsValidResourceType is this a valid resource type?
func (rt ResourceTypes) IsValidResourceType(name string) error {
	for t := range rt {
		if t == name {
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid resource type", name)
}

// IsValidResourceOptionKeys is this valid resource option keys?
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
