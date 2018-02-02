package system

import (
	"github.com/juliengk/go-utils/filedir"
	"github.com/kassisol/hbm/cli/command"
	"github.com/kassisol/hbm/docker/endpoint"
	configobj "github.com/kassisol/hbm/object/config"
	groupobj "github.com/kassisol/hbm/object/group"
	resourceobj "github.com/kassisol/hbm/object/resource"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var initAction bool

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

	return cmd
}

func runInit(cmd *cobra.Command, args []string) {
	if err := filedir.CreateDirIfNotExist(command.AppPath, false, 0700); err != nil {
		log.Fatal(err)
	}

	s, err := configobj.New("sqlite", command.AppPath)
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

	g, err := groupobj.New("sqlite", command.AppPath)
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

	if initAction {
		r, err := resourceobj.New("sqlite", command.AppPath)
		if err != nil {
			log.Fatal(err)
		}
		defer r.End()

		if r.Count("action") == 0 {
			for _, u := range *endpoint.GetUris() {
				if err := r.Add(u.Action, "action", u.Action, []string{}); err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}

var initDescription = `
Initialize config

`
