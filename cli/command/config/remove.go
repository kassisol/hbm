package config

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/juliengk/go-utils"
	"github.com/kassisol/hbm/cli/command"
	"github.com/kassisol/hbm/storage"
	"github.com/spf13/cobra"
)

func newRemoveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "rm [name]",
		Aliases: []string{"remove", "disable", "dis", "off"},
		Short:   "Disable hbm feature",
		Long:    removeDescription,
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

	if len(args) < 1 || len(args) > 1 {
		cmd.Usage()
		os.Exit(-1)
	}

	if !s.FindConfig(args[0]) {
		log.Fatalf("%s is not enabled", args[0])
	}

	if err = s.RemoveConfig(args[0]); err != nil {
		log.Fatal(err)
	}
}

var removeDescription = `
Disable hbm feature

`
