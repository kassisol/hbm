package resource

import (
	"fmt"

	"github.com/juliengk/go-utils"
	"github.com/kassisol/hbm/cli/command"
	resourcepkg "github.com/kassisol/hbm/docker/resource"
	resourceobj "github.com/kassisol/hbm/object/resource"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	resourceAddType   string
	resourceAddValue  string
	resourceAddOption []string
)

func newAddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add [name]",
		Short: "Add resource to the whitelist",
		Long:  addDescription,
		Args:  cobra.ExactArgs(1),
		Run:   runAdd,
	}

	flags := cmd.Flags()
	flags.StringVarP(&resourceAddType, "type", "t", "action", fmt.Sprintf("Set resource type (%s)", resourcepkg.SupportedDrivers("|")))
	flags.StringVarP(&resourceAddValue, "value", "v", "", "Set resource value")
	flags.StringSliceVarP(&resourceAddOption, "option", "o", []string{}, "Specify options")

	return cmd
}

func runAdd(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	r, err := resourceobj.New("sqlite", command.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer r.End()

	if err := r.Add(args[0], resourceAddType, resourceAddValue, resourceAddOption); err != nil {
		log.Fatal(err)
	}
}

var addDescription = `
Add resource to the whitelist

`
