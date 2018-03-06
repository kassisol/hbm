package config

import (
	"fmt"

	"github.com/juliengk/go-utils"
	"github.com/kassisol/hbm/cli/command"
	configobj "github.com/kassisol/hbm/object/config"
	log "github.com/sirupsen/logrus"
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

	c, err := configobj.New("sqlite", command.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer c.End()

	result, err := c.Get(args[0])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}

var getDescription = `
Get config option value

`
