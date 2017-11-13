package config

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/juliengk/go-utils"
	"github.com/kassisol/hbm/cli/command"
	"github.com/kassisol/hbm/storage"
	"github.com/spf13/cobra"
)

func newGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "get [key]",
		Aliases: []string{"find"},
		Short:   "Get config option value",
		Long:    getDescription,
		Args:    cobra.ExactArgs(1),
		Run:     runGet,
	}

	return cmd
}

func runGet(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", command.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer s.End()

	result := s.GetConfig(args[0])

	fmt.Println(result)
}

var getDescription = `
Get config option value

`
