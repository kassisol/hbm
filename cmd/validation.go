package cmd

import (
	"fmt"

	"github.com/juliengk/go-utils"
)

type ResourceTypes map[string]string

func NewResourceTypes() ResourceTypes {
	return map[string]string{
		"action":   "Action",
		"cap":      "Capability",
		"config":   "Config",
		"device":   "Device",
		"dns":      "DNS server",
		"image":    "Image",
		"port":     "Port",
		"registry": "Registry",
		"volume":   "Volume",
	}
}

func (rt ResourceTypes) IsValidResourceType(name string) error {
	for t, _ := range rt {
		if t == name {
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid resource type", name)
}

func IsValidPolicyFilterKeys(filters map[string]string) error {
	validKeys := []string{
		"user",
		"group",
		"host",
		"cluster",
		"resource-type",
		"resource-value",
		"collection",
	}

	for k, _ := range filters {
		if !utils.StringInSlice(k, validKeys) {
			return fmt.Errorf("%s is not a valid filter key", k)
		}
	}

	return nil
}
