package collection

import (
	"github.com/juliengk/go-utils"
	"github.com/kassisol/hbm/cli/command"
	collectionobj "github.com/kassisol/hbm/object/collection"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func newAddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add [name]",
		Short: "Add collection to the whitelist",
		Long:  addDescription,
		Args:  cobra.ExactArgs(1),
		Run:   runAdd,
	}

	return cmd
}

func runAdd(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	c, err := collectionobj.New("sqlite", command.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer c.End()

	if err := c.Add(args[0]); err != nil {
		log.Fatal(err)
	}
}

var addDescription = `
Add collection to the whitelist

`
