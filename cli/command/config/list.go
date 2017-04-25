package config

import (
	"fmt"
	"os"
	"text/tabwriter"

	log "github.com/Sirupsen/logrus"
	"github.com/juliengk/go-utils"
	"github.com/kassisol/hbm/cli/command"
	"github.com/kassisol/hbm/storage"
	"github.com/spf13/cobra"
)

func newListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ls",
		Aliases: []string{"list"},
		Short:   "List hbm enabled features",
		Long:    listDescription,
		Run:     runList,
	}

	return cmd
}

func runList(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", command.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer s.End()

	configs := s.ListConfigs()

	if len(configs) > 0 {
		w := tabwriter.NewWriter(os.Stdout, 20, 1, 2, ' ', 0)
		fmt.Fprintln(w, "NAME")

		for _, config := range configs {
			fmt.Fprintf(w, "%s\n", config)
		}

		w.Flush()
	}
}

var listDescription = `
List hbm enabled features

`
