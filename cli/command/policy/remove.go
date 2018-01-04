package policy

import (
	"github.com/juliengk/go-utils"
	"github.com/kassisol/hbm/cli/command"
	policyobj "github.com/kassisol/hbm/object/policy"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func newRemoveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "rm [name]",
		Aliases: []string{"remove"},
		Short:   "Remove policy",
		Long:    removeDescription,
		Args:    cobra.ExactArgs(1),
		Run:     runRemove,
	}

	return cmd
}

func runRemove(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	p, err := policyobj.New("sqlite", command.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer p.End()

	if err := p.Remove(args[0]); err != nil {
		log.Fatal(err)
	}
}

var removeDescription = `
Remove policy

`
