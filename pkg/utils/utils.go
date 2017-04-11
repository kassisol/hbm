package utils

import (
	"log"
	"net/url"
)

func GetURLParams(r string) url.Values {
	u, err := url.ParseRequestURI(r)
	if err != nil {
		log.Fatal(err)
	}

	return u.Query()
}

func RemoveLastChar(s string) string {
        strLen := len(s) - 1
        newStr := s
        if strLen > 0 {
                newStr = s[0:strLen]
        }

        return newStr
}
