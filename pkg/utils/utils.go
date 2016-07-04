package utils

import (
	"os"
	"regexp"
	"sort"

	"github.com/docker/go-plugins-helpers/authorization"
)

func StringInSlice(haystack []string, needle string) bool {
	s := sort.StringSlice(haystack)
	s.Sort()
	return s.Search(needle) < len(haystack)
}

func GetURIInfo(req authorization.Request) (string, string) {
	reURI := regexp.MustCompile(`^/(v[0-9]+\.[0-9]+)(/.*)`)

	result := reURI.FindStringSubmatch(req.RequestURI)

	return result[1], result[2]
}

func FileExists(f string) bool {
	_, err := os.Lstat(f)
	if err != nil {
		return false
	}

	return true
}
