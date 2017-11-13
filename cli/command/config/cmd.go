package config

import (
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage HBM configs",
		Long:  configDescription,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}

	cmd.AddCommand(
		newGetCommand(),
		newSetCommand(),
		newListCommand(),
	)

	return cmd
}

var configDescription = `
The **hbm config** command has subcommands for managing hbm configs.

To see help for a subcommand, use:

    hbm config [command] --help

For full details on using hbm config visit Harbormaster's online documentation.

`
