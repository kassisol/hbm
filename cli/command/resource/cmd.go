package resource

import (
	"github.com/spf13/cobra"
)

// NewCommand new resource command
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resource",
		Short: "Manage whitelisted resources",
		Long:  resourceDescription,
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

var resourceDescription = `
The **hbm resource** command has subcommands for managing whitelisted resources.

To see help for a subcommand, use:

    hbm resource [command] --help

For full details on using hbm cluster visit Harbormaster's online documentation.

`
