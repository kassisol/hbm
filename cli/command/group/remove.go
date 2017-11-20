package group

import (
	"github.com/juliengk/go-utils"
	"github.com/kassisol/hbm/cli/command"
	"github.com/kassisol/hbm/storage"
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

	s, err := storage.NewDriver("sqlite", command.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer s.End()

	if !s.FindGroup(args[0]) {
		log.Fatalf("%s does not exist", args[0])
	}

	if err = s.RemoveGroup(args[0]); err != nil {
		log.Fatal(err)
	}
}

var removeDescription = `
Remove a group. You cannot remove a group that is in use by a policy.

`
