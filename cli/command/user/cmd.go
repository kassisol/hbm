package user

import (
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "Manage whitelisted users",
		Long:  userDescription,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}

	cmd.AddCommand(
		newAddCommand(),
		newFindCommand(),
		newListCommand(),
		newMemberCommand(),
		newRemoveCommand(),
	)

	return cmd
}

var userDescription = `
The **hbm user** command has subcommands for managing whitelisted users.

To see help for a subcommand, use:

    hbm user [command] --help

For full details on using hbm user visit Harbormaster's online documentation.

`
