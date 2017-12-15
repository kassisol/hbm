package resource

import (
	"fmt"

	"github.com/juliengk/go-utils"
	"github.com/juliengk/go-utils/json"
	"github.com/juliengk/go-utils/validation"
	"github.com/kassisol/hbm/cli/command"
	resourcepkg "github.com/kassisol/hbm/docker/resource"
	"github.com/kassisol/hbm/storage"
	"github.com/kassisol/hbm/storage/driver"
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

	s, err := storage.NewDriver("sqlite", command.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer s.End()

	if err = validation.IsValidName(args[0]); err != nil {
		log.Fatal(err)
	}

	if s.FindResource(args[0]) {
		log.Fatalf("%s already exists", args[0])
	}

	res, err := resourcepkg.NewDriver(resourceAddType)
	if err != nil {
		log.Fatal(err)
	}

	if err = res.Valid(resourceAddValue); err != nil {
		log.Fatal(err)
	}

	options := utils.ConvertSliceToMap("=", resourceAddOption)
	if err = res.ValidOptions(options); err != nil {
		log.Fatal(err)
	}

	opts := ""
	if resourceAddType == "volume" {
		vo := driver.VolumeOptions{}
		if _, ok := options["recursive"]; ok {
			vo.Recursive = true
		}
		if _, ok := options["nosuid"]; ok {
			vo.NoSuid = true
		}
		jsonR := json.Encode(vo)
		opts = jsonR.String()
	}

	// Add to DB
	s.AddResource(args[0], resourceAddType, resourceAddValue, opts)
}

var addDescription = `
Add resource to the whitelist

`
