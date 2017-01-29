package config

import (
	"fmt"
	"os"

	"github.com/juliengk/go-utils"
	"github.com/kassisol/hbm/cli/command"
	"github.com/kassisol/hbm/config"
	"github.com/kassisol/hbm/storage"
	"github.com/spf13/cobra"
)

func newAddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "add [feature]",
		Aliases: []string{"enable", "en", "on"},
		Short:   "Enable hbm feature",
		Long:    addDescription,
		Run:     runAdd,
	}

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

	configs := config.New()
	if configs.ConfigExists(args[0]) {
		utils.Exit(fmt.Errorf("This feature does not exist"))
	}

	if s.FindConfig(args[0]) {
		utils.Exit(fmt.Errorf("%s is already enabled", args[0]))
	}

	s.AddConfig(args[0])
}

var addDescription = `
Enable hbm feature

`
