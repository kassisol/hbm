package policy

import (
	"fmt"
	"os"
	"text/tabwriter"

	log "github.com/sirupsen/logrus"
	"github.com/juliengk/go-utils"
	"github.com/kassisol/hbm/cli/command"
	"github.com/kassisol/hbm/cli/validation"
	"github.com/kassisol/hbm/storage"
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

	s, err := storage.NewDriver("sqlite", command.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer s.End()

	filters := utils.ConvertSliceToMap("=", policyListFilter)

	if err = validation.IsValidPolicyFilterKeys(filters); err != nil {
		log.Fatal(err)
	}

	policies := s.ListPolicies(filters)

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
