package uri

import (
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/harbourmaster/hbm/api/types"
)

type URI struct {
	Method      string
	Re          *regexp.Regexp
	AllowFunc   func(authorization.Request, *types.Config) *types.AllowResult
	DCBFunc     func(authorization.Request, *regexp.Regexp) string
	Action      string
	CmdName     string
	Description string
}

type URIs []URI

func New() *URIs {
	return &URIs{}
}

func (uris *URIs) Register(method, uri string, af func(authorization.Request, *types.Config) *types.AllowResult, dcbf func(authorization.Request, *regexp.Regexp) string, action, cmdName, desc string) {
	*uris = append(*uris, URI{Method: method, Re: regexp.MustCompile(uri), AllowFunc: af, DCBFunc: dcbf, Action: action, CmdName: cmdName, Description: desc})
}

func (uris *URIs) UnRegister(method, uri string, f func(authorization.Request, *types.Config) *types.AllowResult, dcbf func(authorization.Request, *regexp.Regexp) string, action, cmdName, desc string) {
	for i, val := range *uris {
		if val.Action == action {
			*uris = append((*uris)[:i], (*uris)[i+1:]...)
		}
	}
}
