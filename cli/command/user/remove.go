package user

import (
	"github.com/juliengk/go-utils"
	"github.com/kassisol/hbm/cli/command"
	userobj "github.com/kassisol/hbm/object/user"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func newRemoveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "rm [name]",
		Aliases: []string{"remove"},
		Short:   "Remove user from the whitelist",
		Long:    removeDescription,
		Args:    cobra.ExactArgs(1),
		Run:     runRemove,
	}

	return cmd
}

func runRemove(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	u, err := userobj.New("sqlite", command.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer u.End()

	if err := u.Remove(args[0]); err !=nil {
		log.Fatal(err)
	}
}

var removeDescription = `
Remove user from the whitelist

`
