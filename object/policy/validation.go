package policy

import (
	"fmt"

	"github.com/juliengk/go-utils"
)

func isValidFilterKeys(filters map[string]string) error {
	validKeys := []string{
		"user",
		"group",
		"resource-type",
		"resource-value",
		"resource-options",
		"collection",
	}

	for k := range filters {
		if !utils.StringInSlice(k, validKeys, false) {
			return fmt.Errorf("%s is not a valid filter key", k)
		}
	}

	return nil
}
