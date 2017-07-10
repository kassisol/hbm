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

func GetURIInfo(defaultVersion string, req authorization.Request) (URIInfo, error) {
	reURIWithVersion := regexp.MustCompile(`^/(v[0-9]+\.[0-9]+)(/.*)`)
	reURIWithoutVersion := regexp.MustCompile(`^(/.*)`)

	u, err := url.ParseRequestURI(req.RequestURI)
	if err != nil {
		return URIInfo{}, err
	}

	result := reURIWithVersion.FindStringSubmatch(u.Path)

	if len(result) == 0 {
		r := reURIWithoutVersion.FindStringSubmatch(u.Path)
		if len(r) > 0 {
			return URIInfo{Version: defaultVersion, Path: r[1]}, nil
		}
	} else {
		return URIInfo{Version: result[1], Path: result[2]}, nil
	}

	return URIInfo{}, fmt.Errorf("%s is not compatible", u.Path)
}
