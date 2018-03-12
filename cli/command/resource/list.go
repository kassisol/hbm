package resource

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/juliengk/go-utils"
	resourceobj "github.com/kassisol/hbm/object/resource"
	"github.com/kassisol/hbm/pkg/adf"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var resourceListFilter []string

func newListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ls",
		Aliases: []string{"list"},
		Short:   "List whitelisted resources",
		Long:    listDescription,
		Args:    cobra.NoArgs,
		Run:     runList,
	}

	flags := cmd.Flags()
	flags.StringSliceVarP(&resourceListFilter, "filter", "f", []string{}, "Filter output based on conditions provided")

	return cmd
}

func runList(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	r, err := resourceobj.New("sqlite", adf.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer r.End()

	filters := utils.ConvertSliceToMap("=", resourceListFilter)

	resources, err := r.List(filters)
	if err != nil {
		log.Fatal(err)
	}

	if len(resources) > 0 {
		tw := tabwriter.NewWriter(os.Stdout, 20, 1, 2, ' ', 0)
		fmt.Fprintln(tw, "NAME\tTYPE\tVALUE\tOPTIONS\tCOLLECTIONS")

		for resource, collections := range resources {
			if len(collections) > 0 {
				fmt.Fprintf(tw, "%s\t%s\t%s\t%s\t%s\n", resource.Name, resource.Type, resource.Value, utils.RemoveLastChar(resource.Option), strings.Join(collections, ", "))
			} else {
				fmt.Fprintf(tw, "%s\t%s\t%s\t%s\n", resource.Name, resource.Type, resource.Value, utils.RemoveLastChar(resource.Option))
			}
		}

		tw.Flush()
	}
}

var listDescription = `
List whitelisted resources

`
