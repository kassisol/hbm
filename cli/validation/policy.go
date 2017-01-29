package validation

import (
	"fmt"

	"github.com/juliengk/go-utils"
)

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
