package group

import (
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "group",
		Short: "Manage whitelisted groups",
		Long:  groupDescription,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}

	cmd.AddCommand(
		newAddCommand(),
		newFindCommand(),
		newListCommand(),
		newRemoveCommand(),
	)

	return cmd
}

var groupDescription = `
The **hbm group** command has subcommands for managing whitelisted groups.

To see help for a subcommand, use:

    hbm group [command] --help

For full details on using hbm cluster visit Harbormaster's online documentation.

`
