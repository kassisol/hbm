package config

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/juliengk/go-utils"
	configobj "github.com/kassisol/hbm/object/config"
	"github.com/kassisol/hbm/pkg/adf"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var configListFilter []string

func newListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ls",
		Aliases: []string{"list"},
		Short:   "List HBM configs",
		Long:    listDescription,
		Args:    cobra.NoArgs,
		Run:     runList,
	}

	flags := cmd.Flags()
	flags.StringSliceVarP(&configListFilter, "filter", "f", []string{}, "Filter output based on conditions provided")

	return cmd
}

func runList(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	c, err := configobj.New("sqlite", adf.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer c.End()

	filters := utils.ConvertSliceToMap("=", configListFilter)

	configs, err := c.List(filters)
	if err != nil {
		log.Fatal(err)
	}

	if len(configs) > 0 {
		w := tabwriter.NewWriter(os.Stdout, 20, 1, 2, ' ', 0)
		fmt.Fprintln(w, "NAME\tVALUE")

		for _, config := range configs {
			fmt.Fprintf(w, "%s\t%t\n", config.Key, config.Value)
		}

		w.Flush()
	}
}

var listDescription = `
List HBM configs

`
