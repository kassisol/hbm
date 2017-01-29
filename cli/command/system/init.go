package system

import (
	"os"

	"github.com/juliengk/go-utils"
	"github.com/kassisol/hbm/cli/command"
	"github.com/kassisol/hbm/storage"
	"github.com/spf13/cobra"
)

func NewInitCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize config",
		Long:  initDescription,
		Run:   runInit,
	}

	return cmd
}

func runInit(cmd *cobra.Command, args []string) {
	if _, err := os.Stat(command.AppPath); os.IsNotExist(err) {
		err := os.Mkdir(command.AppPath, 0700)
		if err != nil {
			utils.Exit(err)
		}
	}

	s, err := storage.NewDriver("sqlite", command.AppPath)
	if err != nil {
		utils.Exit(err)
	}
	s.End()
}

var initDescription = `
Initialize config

`
