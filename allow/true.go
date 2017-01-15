package allow

import (
	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/allow/types"
)

func AllowTrue(req authorization.Request, config *types.Config) *types.AllowResult {
	return &types.AllowResult{Allow: true}
}
