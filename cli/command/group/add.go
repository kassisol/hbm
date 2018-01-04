package group

import (
	"github.com/juliengk/go-utils"
	"github.com/kassisol/hbm/cli/command"
	groupobj "github.com/kassisol/hbm/object/group"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func newAddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add [name]",
		Short: "Add group to the whitelist",
		Long:  addDescription,
		Args:  cobra.ExactArgs(1),
		Run:   runAdd,
	}

	return cmd
}

func runAdd(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	g, err := groupobj.New("sqlite", command.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer g.End()

	if err := g.Add(args[0]); err != nil {
		log.Fatal(err)
	}
}

var addDescription = `
Add group to the whitelist

`
