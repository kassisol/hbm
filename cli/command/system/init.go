package system

import (
	"reflect"

	"github.com/juliengk/go-utils"
	"github.com/juliengk/go-utils/filedir"
	"github.com/kassisol/hbm/docker/endpoint"
	resourcepkg "github.com/kassisol/hbm/docker/resource"
	rconfigdrv "github.com/kassisol/hbm/docker/resource/driver/config"
	configobj "github.com/kassisol/hbm/object/config"
	groupobj "github.com/kassisol/hbm/object/group"
	resourceobj "github.com/kassisol/hbm/object/resource"
	"github.com/kassisol/hbm/pkg/adf"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	initAction bool
	initConfig bool
)

func NewInitCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize config",
		Long:  initDescription,
		Args:  cobra.NoArgs,
		Run:   runInit,
	}

	flags := cmd.Flags()
	flags.BoolVarP(&initAction, "action", "", false, "Initialize action resources")
	flags.BoolVarP(&initConfig, "config", "", false, "Initialize config resources")

	return cmd
}

func runInit(cmd *cobra.Command, args []string) {
	if err := filedir.CreateDirIfNotExist(adf.AppPath, false, 0700); err != nil {
		log.Fatal(err)
	}

	s, err := configobj.New("sqlite", adf.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer s.End()

	config, err := s.List(map[string]string{})
	if err != nil {
		log.Fatal(err)
	}

	if len(config) == 0 {
		s.Set("authorization", "false")
		s.Set("default-allow-action-error", "false")
	}

	g, err := groupobj.New("sqlite", adf.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer g.End()

	filters := map[string]string{
		"name": "administrators",
	}
	groups, _ := g.List(filters)
	if len(groups) == 0 {
		g.Add("administrators")
	}

	r, err := resourceobj.New("sqlite", adf.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer r.End()

	if initAction {
		if r.Count("action") == 0 {
			for _, u := range *endpoint.GetUris() {
				if err := r.Add(u.Action, "action", u.Action, []string{}); err != nil {
					log.Fatal(err)
				}
			}
		}
	}

	if initConfig {
		if r.Count("config") == 0 {
			res, err := resourcepkg.NewDriver("config")
			if err != nil {
				log.Fatal(err)
			}

			val := utils.GetReflectValue(reflect.Slice, res.List())
			v := val.Interface().([]rconfigdrv.Action)

			for _, c := range v {
				if err := r.Add(c.Key, "config", c.Key, []string{}); err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}

var initDescription = `
Initialize config

`
