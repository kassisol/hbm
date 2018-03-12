package group

import (
	"fmt"

	"github.com/juliengk/go-utils"
	groupobj "github.com/kassisol/hbm/object/group"
	"github.com/kassisol/hbm/pkg/adf"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func newFindCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "find [name]",
		Short: "Verify if group exists in the whitelist",
		Long:  findDescription,
		Args:  cobra.ExactArgs(1),
		Run:   runFind,
	}

	return cmd
}

func runFind(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	g, err := groupobj.New("sqlite", adf.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer g.End()

	fmt.Println(g.Find(args[0]))
}

var findDescription = `
Verify if group exists in the whitelist

`
