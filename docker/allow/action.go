package allow

import (
	"fmt"

	"github.com/juliengk/go-log"
	"github.com/juliengk/go-log/driver"
	"github.com/juliengk/go-utils"
	"github.com/kassisol/hbm/docker/allow/types"
	policyobj "github.com/kassisol/hbm/object/policy"
	"github.com/kassisol/hbm/version"
)

func Action(config *types.Config, action, cmd string) *types.AllowResult {
	defer utils.RecoverFunc()

	l, _ := log.NewDriver("standard", nil)

	p, err := policyobj.New("sqlite", config.AppPath)
	if err != nil {
		l.WithFields(driver.Fields{
			"storagedriver": "sqlite",
			"logdriver":     "standard",
			"version":       version.Version,
		}).Fatal(err)
	}
	defer p.End()

	if !p.Validate(config.Username, "action", action, "") {
		return &types.AllowResult{Allow: false, Msg: fmt.Sprintf("%s is not allowed", cmd)}
	}

	return &types.AllowResult{Allow: true}
}
