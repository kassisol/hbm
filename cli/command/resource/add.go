package resource

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/juliengk/go-utils"
	"github.com/juliengk/go-utils/json"
	"github.com/juliengk/go-utils/validation"
	"github.com/kassisol/hbm/cli/command"
	clivalidation "github.com/kassisol/hbm/cli/validation"
	"github.com/kassisol/hbm/docker/config"
	"github.com/kassisol/hbm/docker/endpoint"
	"github.com/kassisol/hbm/storage"
	"github.com/kassisol/hbm/storage/driver"
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
		Run:   runAdd,
	}

	flags := cmd.Flags()
	flags.StringVarP(&resourceAddType, "type", "t", "action", "Set resource type (action|cap|config|device|dns|image|logdriver|logopt|port|registry|volume)")
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

	// Inputs validation
	if len(args) < 1 || len(args) > 1 {
		cmd.Usage()
		os.Exit(-1)
	}

	options := utils.ConvertSliceToMap("=", resourceAddOption)
	if len(options) > 0 {
		if err := clivalidation.IsValidResourceOptionKeys(options); err != nil {
			log.Fatal(err)
		}
	}

	rt := clivalidation.NewResourceTypes()
	if err = rt.IsValidResourceType(resourceAddType); err != nil {
		log.Fatal(err)
	}

	if resourceAddType == "action" {
		uris := endpoint.GetUris()

		if !uris.ActionExists(resourceAddValue) {
			log.Fatalf("%s is not a valid action", resourceAddValue)
		}
	}

	if resourceAddType == "cap" {
		if !clivalidation.IsValidCapability(resourceAddValue) {
			log.Fatalf("%s is not a valid cap", resourceAddValue)
		}
	}

	if resourceAddType == "config" {
		configs := config.New()

		if !configs.ActionExists(resourceAddValue) {
			log.Fatalf("%s is not a valid config", resourceAddValue)
		}
	}

	if resourceAddType == "logdriver" {
		if !clivalidation.IsValidLogDriver(resourceAddValue) {
			log.Fatalf("%s is not a valid logdriver", resourceAddValue)
		}
	}

	if err = validation.IsValidName(args[0]); err != nil {
		log.Fatal(err)
	}

	if s.FindResource(args[0]) {
		log.Fatalf("%s already exists", args[0])
	}

	// Add to DB
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

	s.AddResource(args[0], resourceAddType, resourceAddValue, opts)
}

var addDescription = `
Add resource to the whitelist

`
