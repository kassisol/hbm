package cmd

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/juliengk/go-utils"
	"github.com/juliengk/go-utils/validation"
	"github.com/kassisol/hbm/storage"
	"github.com/spf13/cobra"
)

var clusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Manage whitelisted clusters",
	Long:  "Manage whitelisted clusters",
}

var clusterListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List whitelisted clusters",
	Long:  "List whitelisted clusters",
}

var clusterAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add cluster to the whitelist",
	Long:  "Add cluster to the whitelist",
}

var clusterRemoveCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove cluster from the whitelist",
	Long:  "Remove cluster from the whitelist",
}

var clusterExistsCmd = &cobra.Command{
	Use:   "find",
	Short: "Verify if cluster exists in the whitelist",
	Long:  "Verify if cluster exists in the whitelist",
}

func init() {
	RootCmd.AddCommand(clusterCmd)
	clusterCmd.AddCommand(clusterListCmd)
	clusterCmd.AddCommand(clusterAddCmd)
	clusterCmd.AddCommand(clusterRemoveCmd)
	clusterCmd.AddCommand(clusterExistsCmd)

	clusterCmd.Run = clusterUsage
	clusterListCmd.Run = clusterList
	clusterAddCmd.Run = clusterAdd
	clusterRemoveCmd.Run = clusterRemove
	clusterExistsCmd.Run = clusterExists
}

func clusterUsage(cmd *cobra.Command, args []string) {
	clusterCmd.Usage()
	os.Exit(-1)
}

func clusterList(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", appPath)
	if err != nil {
		utils.Exit(err)
	}
	defer s.End()

	clusters := s.ListClusters()

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

func clusterAdd(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", appPath)
	if err != nil {
		utils.Exit(err)
	}
	defer s.End()

	if err = validation.IsValidName(args[0]); err != nil {
		utils.Exit(err)
	}

	s.AddCluster(args[0])
}

func clusterRemove(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", appPath)
	if err != nil {
		utils.Exit(err)
	}
	defer s.End()

	if !s.FindCluster(args[0]) {
		utils.Exit(fmt.Errorf("%s does not exist", args[0]))
	}

	s.RemoveCluster(args[0])
}

func clusterExists(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", appPath)
	if err != nil {
		utils.Exit(err)
	}
	defer s.End()

	result := s.FindCluster(args[0])

	fmt.Println(result)
}
