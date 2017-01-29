package config

import (
	"fmt"
	"os"

	"github.com/juliengk/go-utils"
	"github.com/kassisol/hbm/cli/command"
	"github.com/kassisol/hbm/storage"
	"github.com/spf13/cobra"
)

func newFindCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "find [key]",
		Short: "Verify if feature is enabled",
		Long:  findDescription,
		Run:   runFind,
	}

	return cmd
}

func runFind(cmd *cobra.Command, args []string) {
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

	result := s.FindConfig(args[0])

	fmt.Println(result)
}

var findDescription = `
Verify if feature is enabled

`
