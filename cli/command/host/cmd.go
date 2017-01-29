package host

import (
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "host",
		Short: "Manage whitelisted hosts",
		Long:  hostDescription,
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

var hostDescription = `
The **hbm host** command has subcommands for managing whitelisted hosts.

To see help for a subcommand, use:

    hbm host [command] --help

For full details on using hbm cluster visit Harbormaster's online documentation.

`
