package cluster

import (
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cluster",
		Short: "Manage whitelisted clusters",
		Long:  clusterDescription,
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

var clusterDescription = `
The **hbm cluster** command has subcommands for managing whitelisted clusters.

To see help for a subcommand, use:

    hbm cluster [command] --help

For full details on using hbm cluster visit Harbormaster's online documentation.

`
