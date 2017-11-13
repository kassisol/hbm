package system

import (
	"github.com/juliengk/go-utils/filedir"
	"github.com/kassisol/hbm/cli/command"
	"github.com/kassisol/hbm/storage"
	log "github.com/sirupsen/logrus"
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
	if err := filedir.CreateDirIfNotExist(command.AppPath, false, 0700); err != nil {
		log.Fatal(err)
	}

	s, err := storage.NewDriver("sqlite", command.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer s.End()

	s.SetConfig("authorization", false)
	s.SetConfig("default-allow-action-error", false)
}

var initDescription = `
Initialize config

`
