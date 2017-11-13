package config

import (
	"strconv"

	"github.com/juliengk/go-utils"
	"github.com/kassisol/hbm/cli/command"
	"github.com/kassisol/hbm/config"
	"github.com/kassisol/hbm/storage"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func newSetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set [key] [value]",
		Short: "Set HBM config option",
		Long:  setDescription,
		Args:  cobra.ExactArgs(2),
		Run:   runSet,
	}

	return cmd
}

func runSet(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", command.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer s.End()

	configs := config.New()
	if err := configs.IsValid(args[0]); err != nil {
		log.Fatal(err)
	}

	value, err := strconv.ParseBool(args[1])
	if err != nil {
		log.Fatal(err)
	}

	s.SetConfig(args[0], value)
}

var setDescription = `
Set HBM config option

`
