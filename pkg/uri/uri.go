package uri

import (
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/harbourmaster/hbm/api/types"
)

type URI struct {
	Method		string
	Re		*regexp.Regexp
	Func		func(authorization.Request, *types.Config) (string, string)
	Action		string
	CmdName		string
	Description	string
}

type URIs []URI

func New() *URIs {
	return &URIs{}
}

func (uris *URIs) Register(method, uri string, f func(authorization.Request, *types.Config) (string, string), action, cmdName, desc string) {
	*uris = append(*uris, URI{Method: method, Re: regexp.MustCompile(uri), Func: f, Action: action, CmdName: cmdName, Description: desc})
}
