package allow

import (
	"fmt"

	"github.com/juliengk/go-log"
	"github.com/juliengk/go-log/driver"
	"github.com/juliengk/go-utils"
	"github.com/kassisol/hbm/allow/types"
	"github.com/kassisol/hbm/storage"
	"github.com/kassisol/hbm/version"
)

// Action called from plugin
func Action(config *types.Config, action, cmd string) *types.AllowResult {
	defer utils.RecoverFunc()

	l, _ := log.NewDriver("standard", nil)

	s, err := storage.NewDriver("sqlite", config.AppPath)
	if err != nil {
		l.WithFields(driver.Fields{
			"storagedriver": "sqlite",
			"logdriver":     "standard",
			"version":       version.Version,
		}).Fatal(err)
	}
	defer s.End()

	if !s.ValidatePolicy(config.Username, "action", action, "") {
		return &types.AllowResult{Allow: false, Error: fmt.Sprintf("%s is not allowed", cmd)}
	}

	return &types.AllowResult{Allow: true}
}
