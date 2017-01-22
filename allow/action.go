package allow

import (
	"fmt"
	"log"

	"github.com/juliengk/go-utils"
	"github.com/kassisol/hbm/allow/types"
	"github.com/kassisol/hbm/storage"
)

func AllowAction(config *types.Config, action, cmd string) *types.AllowResult {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", config.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer s.End()

	if !s.ValidatePolicy(config.Username, config.Hostname, "action", action, "") {
		return &types.AllowResult{Allow: false, Error: fmt.Sprintf("%s is not allowed", cmd)}
	}

	return &types.AllowResult{Allow: true}
}
