package group

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/juliengk/go-utils"
	"github.com/juliengk/go-utils/validation"
	"github.com/kassisol/hbm/cli/command"
	"github.com/kassisol/hbm/storage"
	"github.com/spf13/cobra"
)

func newAddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add [name]",
		Short: "Add group to the whitelist",
		Long:  addDescription,
		Run:   runAdd,
	}

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

	if err = validation.IsValidGroupname(args[0]); err != nil {
		log.Fatal(err)
	}

	if s.FindGroup(args[0]) {
		log.Fatalf("%s already exists", args[0])
	}

	s.AddGroup(args[0])
}

var addDescription = `
Add group to the whitelist

`
