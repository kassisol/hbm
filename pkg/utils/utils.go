package utils

import (
	"log"
	"net/url"
	"os"
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
)

func GetURIInfo(req authorization.Request) (string, string) {
	reURI := regexp.MustCompile(`^/(v[0-9]+\.[0-9]+)(/.*)`)

	u, err := url.ParseRequestURI(req.RequestURI)
	if err != nil {
		log.Fatal(err)
	}

	result := reURI.FindStringSubmatch(u.Path)

	return result[1], result[2]
}

func GetURLParams(r string) url.Values {
	u, err := url.ParseRequestURI(r)
	if err != nil {
		log.Fatal(err)
	}

	return u.Query()
}

func FileExists(f string) bool {
	_, err := os.Lstat(f)
	if err != nil {
		return false
	}

	return true
}
