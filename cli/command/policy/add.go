package policy

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/juliengk/go-utils"
	"github.com/juliengk/go-utils/validation"
	"github.com/kassisol/hbm/cli/command"
	"github.com/kassisol/hbm/storage"
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
		Run:   runAdd,
	}

	flags := cmd.Flags()
	flags.StringVarP(&policyAddGroup, "group", "g", "", "Set group")
	flags.StringVarP(&policyAddCollection, "collection", "c", "", "Set collection")

	return cmd
}

func runAdd(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", command.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer s.End()

	if len(args) < 1 || len(args) > 1 {
		cmd.Usage()
		os.Exit(-1)
	}

	if err = validation.IsValidName(args[0]); err != nil {
		log.Fatal(err)
	}

	if s.FindPolicy(args[0]) {
		log.Fatalf("%s already exists", args[0])
	}

	if policyAddGroup != "all" && !s.FindGroup(policyAddGroup) {
		log.Fatalf("%s does not exist", policyAddGroup)
	}

	if !s.FindCollection(policyAddCollection) {
		log.Fatalf("%s does not exist", policyAddCollection)
	}

	s.AddPolicy(args[0], policyAddGroup, policyAddCollection)
}

var addDescription = `
Add policy

`
