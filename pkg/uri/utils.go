package uri

import (
	"fmt"
	"net/url"
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
)

type URIInfo struct {
	Version string
	Path    string
}

func GetURIInfo(req authorization.Request) (URIInfo, error) {
	reURI := regexp.MustCompile(`^/(v[0-9]+\.[0-9]+)(/.*)`)

	u, err := url.ParseRequestURI(req.RequestURI)
	if err != nil {
		return URIInfo{}, err
	}

	result := reURI.FindStringSubmatch(u.Path)

	if len(result) == 0 {
		return URIInfo{}, fmt.Errorf("%s is not compatible", u.Path)
	}

	return URIInfo{Version: result[1], Path: result[2]}, nil
}
