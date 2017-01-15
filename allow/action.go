package allow

import (
	"fmt"
	"log"

	"github.com/kassisol/hbm/allow/types"
	"github.com/kassisol/hbm/pkg/db"
)

func AllowAction(config *types.Config, action, cmd string) *types.AllowResult {
	defer db.RecoverFunc()

	d, err := db.NewDB(config.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer d.Conn.Close()

	// Validate Docker command is allowed
	if !d.KeyExists("action", action) {
		return &types.AllowResult{Allow: false, Error: fmt.Sprintf("%s is not allowed", cmd)}
	}

	return &types.AllowResult{Allow: true}
}
