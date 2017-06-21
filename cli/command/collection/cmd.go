package collection

import (
	"github.com/spf13/cobra"
)

// NewCommand new collection command
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "collection",
		Short: "Manage whitelisted collections",
		Long:  collectionDescription,
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

var collectionDescription = `
The **hbm collection** command has subcommands for managing whitelisted collections.

To see help for a subcommand, use:

    hbm collection [command] --help

For full details on using hbm cluster visit Harbormaster's online documentation.

`
