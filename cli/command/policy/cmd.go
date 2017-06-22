package policy

import (
	"github.com/spf13/cobra"
)

// NewCommand new policy command
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "policy",
		Short: "Manage policies",
		Long:  policyDescription,
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

var policyDescription = `
The **hbm policy** command has subcommands for managing whitelisted policies.

To see help for a subcommand, use:

    hbm policy [command] --help

For full details on using hbm cluster visit Harbormaster's online documentation.

`
