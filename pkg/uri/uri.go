package uri

import (
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/allow/types"
)

// URI structure
type URI struct {
	Method      string
	Re          *regexp.Regexp
	AllowFunc   func(authorization.Request, *types.Config) *types.AllowResult
	DCBFunc     func(authorization.Request, string, *regexp.Regexp) string
	Action      string
	CmdName     string
	Description string
}

// URIs structure
type URIs []URI

// New function
func New() *URIs {
	return &URIs{}
}

// Register function
func (uris *URIs) Register(method, uri string, af func(authorization.Request, *types.Config) *types.AllowResult, dcbf func(authorization.Request, string, *regexp.Regexp) string, action, cmdName, desc string) {
	*uris = append(*uris, URI{Method: method, Re: regexp.MustCompile(uri), AllowFunc: af, DCBFunc: dcbf, Action: action, CmdName: cmdName, Description: desc})
}

// GetURI function
func (uris *URIs) GetURI(method, url string) (URI, error) {
	for _, u := range *uris {
		if u.Method == method {
			if u.Re.MatchString(url) {
				return u, nil
			}
		}
	}

	return URI{}, nil
}

// ActionExists function
func (uris *URIs) ActionExists(action string) bool {
	for _, u := range *uris {
		if u.Action == action {
			return true
		}
	}

	return false
}
