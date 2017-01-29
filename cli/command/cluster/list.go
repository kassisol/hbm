package cluster

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/juliengk/go-utils"
	"github.com/kassisol/hbm/cli/command"
	"github.com/kassisol/hbm/storage"
	"github.com/spf13/cobra"
)

var clusterListFilter []string

func newListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ls",
		Aliases: []string{"list"},
		Short:   "List whitelisted clusters",
		Long:    listDescription,
		Run:     runList,
	}

	flags := cmd.Flags()
	flags.StringSliceVarP(&clusterListFilter, "filter", "f", []string{}, "Filter output based on conditions provided")

	return cmd
}

func runList(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", command.AppPath)
	if err != nil {
		utils.Exit(err)
	}
	defer s.End()

	filters := utils.ConvertSliceToMap("=", clusterListFilter)

	clusters := s.ListClusters(filters)

	if len(clusters) > 0 {
		w := tabwriter.NewWriter(os.Stdout, 20, 1, 2, ' ', 0)
		fmt.Fprintln(w, "NAME\tHOSTS")

		for cluster, hosts := range clusters {
			if len(hosts) > 0 {
				fmt.Fprintf(w, "%s\t%s\n", cluster, strings.Join(hosts, ", "))
			} else {
				fmt.Fprintf(w, "%s\n", cluster)
			}
		}

		w.Flush()
	}
}

var listDescription = `
Lists all the clusters Harbormaster knows about. You can filter using the **-f** or
**--filter** flag. The filtering format is a **key=value** pair. To specify more
than one filter, pass multiple flags (for example, **--filter "foo=bar" --filter "bif=baz"**).

The currently supported filters are:

* **name** (a cluster's name)
* **elem** (a host's name)

`
