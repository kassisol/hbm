package image

import (
        "strings"

        "github.com/juliengk/go-utils/validation"
)

func GetNameTag(name string) (string, string) {
        nt := strings.Split(name, ":")
        count := len(nt)

        if count > 2 {
                return strings.Join(nt[0:count-1], ":"), nt[count-1]
        } else if count == 2 {
                if strings.Contains(nt[1], "/") {
                        return name, "latest"
                }
                return nt[0], nt[1]
        } else if count == 1 {
                return nt[0], "latest"
        }

        return "", ""
}

func validateRegistry(value string) bool {
        if len(value) == 0 {
                return false
        }

        result := strings.Split(value, ":")

        if err := validation.IsValidFQDN(result[0]); err == nil {
                return true
        }

        return false
}
