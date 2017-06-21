package validation

import (
	"fmt"

	"github.com/juliengk/go-utils"
)

// IsValidPolicyFilterKeys Is this a valid policy filter key map
func IsValidPolicyFilterKeys(filters map[string]string) error {
	validKeys := []string{
		"user",
		"group",
		"resource-type",
		"resource-value",
		"collection",
	}

	for k := range filters {
		if !utils.StringInSlice(k, validKeys, false) {
			return fmt.Errorf("%s is not a valid filter key", k)
		}
	}

	return nil
}
