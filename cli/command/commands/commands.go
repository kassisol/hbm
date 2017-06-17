package commands

import (
	"github.com/kassisol/hbm/cli/command/collection"
	"github.com/kassisol/hbm/cli/command/config"
	"github.com/kassisol/hbm/cli/command/group"
	"github.com/kassisol/hbm/cli/command/policy"
	"github.com/kassisol/hbm/cli/command/resource"
	"github.com/kassisol/hbm/cli/command/server"
	"github.com/kassisol/hbm/cli/command/system"
	"github.com/kassisol/hbm/cli/command/user"
	"github.com/spf13/cobra"
)

// AddCommands adds the commands
func AddCommands(cmd *cobra.Command) {
	cmd.AddCommand(
		collection.NewCommand(),
		config.NewCommand(),
		group.NewCommand(),
		policy.NewCommand(),
		resource.NewCommand(),
		user.NewCommand(),
		server.NewServerCommand(),
		system.NewInfoCommand(),
		system.NewInitCommand(),
		system.NewVersionCommand(),
	)
}
