package api

import (
	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/harbourmaster/hbm/api/types"
)

func AllowTrue(req authorization.Request, config *types.Config) (string, string) {
	return "", ""
}
