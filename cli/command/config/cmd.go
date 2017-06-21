package config

import (
	"github.com/spf13/cobra"
)

// NewCommand new config command
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage HBM features",
		Long:  configDescription,
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

var configDescription = `
The **hbm config** command has subcommands for managing hbm features.

To see help for a subcommand, use:

    hbm config [command] --help

For full details on using hbm config visit Harbormaster's online documentation.

`
