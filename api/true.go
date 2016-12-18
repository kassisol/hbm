package api

import (
	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/api/types"
)

func AllowTrue(req authorization.Request, config *types.Config) *types.AllowResult {
	return &types.AllowResult{Allow: true}
}
