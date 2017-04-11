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
