package policy

import (
	"fmt"
	"os"

	"github.com/juliengk/go-utils"
	"github.com/juliengk/go-utils/validation"
	"github.com/kassisol/hbm/cli/command"
	"github.com/kassisol/hbm/storage"
	"github.com/spf13/cobra"
)

var (
	policyAddGroup      string
	policyAddCluster    string
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
	flags.StringVarP(&policyAddCluster, "cluster", "c", "", "Set cluster")
	flags.StringVarP(&policyAddCollection, "collection", "", "", "Set collection")

	return cmd
}

func runAdd(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", command.AppPath)
	if err != nil {
		utils.Exit(err)
	}
	defer s.End()

	if len(args) < 1 || len(args) > 1 {
		cmd.Usage()
		os.Exit(-1)
	}

	if err = validation.IsValidName(args[0]); err != nil {
		utils.Exit(err)
	}

	if s.FindPolicy(args[0]) {
		utils.Exit(fmt.Errorf("%s already exists", args[0]))
	}

	if policyAddCluster != "all" && !s.FindCluster(policyAddCluster) {
		utils.Exit(fmt.Errorf("%s does not exist", policyAddCluster))
	}

	if policyAddGroup != "all" && !s.FindGroup(policyAddGroup) {
		utils.Exit(fmt.Errorf("%s does not exist", policyAddGroup))
	}

	if !s.FindCollection(policyAddCollection) {
		utils.Exit(fmt.Errorf("%s does not exist", policyAddCollection))
	}

	s.AddPolicy(args[0], policyAddGroup, policyAddCluster, policyAddCollection)
}

var addDescription = `
Add policy

`
