package cmd

import (
	"os"

	"github.com/juliengk/go-utils"
	"github.com/kassisol/hbm/storage"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize config",
	Long:  "Initialize config",
}

func init() {
	RootCmd.AddCommand(initCmd)

	initCmd.Run = initConfig
}

func initConfig(cmd *cobra.Command, args []string) {
	if _, err := os.Stat(appPath); os.IsNotExist(err) {
		err := os.Mkdir(appPath, 0700)
		if err != nil {
			utils.Exit(err)
		}
	}

	s, err := storage.NewDriver("sqlite", appPath)
	if err != nil {
		utils.Exit(err)
	}
	s.End()
}
