package policy

import (
	"github.com/juliengk/go-utils"
	policyobj "github.com/kassisol/hbm/object/policy"
	"github.com/kassisol/hbm/pkg/adf"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	policyAddGroup      string
	policyAddCollection string
)

func newAddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add [name]",
		Short: "Add policy",
		Long:  addDescription,
		Args:  cobra.ExactArgs(1),
		Run:   runAdd,
	}

	flags := cmd.Flags()
	flags.StringVarP(&policyAddGroup, "group", "g", "", "Set group")
	flags.StringVarP(&policyAddCollection, "collection", "c", "", "Set collection")

	return cmd
}

func runAdd(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	p, err := policyobj.New("sqlite", adf.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer p.End()

	if err := p.Add(args[0], policyAddGroup, policyAddCollection); err != nil {
		log.Fatal(err)
	}
}

var addDescription = `
Add policy

`
