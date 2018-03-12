package policy

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/juliengk/go-utils"
	policyobj "github.com/kassisol/hbm/object/policy"
	"github.com/kassisol/hbm/pkg/adf"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var policyListFilter []string

func newListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ls",
		Aliases: []string{"list"},
		Short:   "List policies",
		Long:    listDescription,
		Args:    cobra.NoArgs,
		Run:     runList,
	}

	flags := cmd.Flags()
	flags.StringSliceVarP(&policyListFilter, "filter", "f", []string{}, "Filter output based on conditions provided")

	return cmd
}

func runList(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	p, err := policyobj.New("sqlite", adf.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer p.End()

	filters := utils.ConvertSliceToMap("=", policyListFilter)

	policies, err := p.List(filters)
	if err != nil {
		log.Fatal(err)
	}

	if len(policies) > 0 {
		w := tabwriter.NewWriter(os.Stdout, 20, 1, 2, ' ', 0)
		fmt.Fprintln(w, "NAME\tGROUP\tCOLLECTION")

		for _, policy := range policies {
			fmt.Fprintf(w, "%s\t%s\t%s\n", policy.Name, policy.Group, policy.Collection)
		}

		w.Flush()
	}
}

var listDescription = `
List policies

`
