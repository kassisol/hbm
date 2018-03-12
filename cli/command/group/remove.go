package group

import (
	"github.com/juliengk/go-utils"
	groupobj "github.com/kassisol/hbm/object/group"
	"github.com/kassisol/hbm/pkg/adf"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func newRemoveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "rm [name]",
		Aliases: []string{"remove"},
		Short:   "Remove group from the whitelist",
		Long:    removeDescription,
		Args:    cobra.ExactArgs(1),
		Run:     runRemove,
	}

	return cmd
}

func runRemove(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	g, err := groupobj.New("sqlite", adf.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer g.End()

	if err := g.Remove(args[0]); err != nil {
		log.Fatal(err)
	}
}

var removeDescription = `
Remove a group. You cannot remove a group that is in use by a policy.

`
