package host

import (
	"fmt"
	"os"

	"github.com/juliengk/go-utils"
	"github.com/kassisol/hbm/cli/command"
	"github.com/kassisol/hbm/storage"
	//	"github.com/juliengk/go-utils/validation"
	"github.com/spf13/cobra"
)

func newAddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add [name]",
		Short: "Add host to the whitelist",
		Long:  addDescription,
		Run:   runAdd,
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

	//if err = utils.IsValidHostname(args[0]); err != nil {
	//	utils.Exit(err)
	//}

	if s.FindHost(args[0]) {
		utils.Exit(fmt.Errorf("%s already exists", args[0]))
	}

	s.AddHost(args[0])
}

var addDescription = `
Add host to the whitelist

`
