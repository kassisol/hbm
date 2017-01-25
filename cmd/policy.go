package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/juliengk/go-utils"
	"github.com/juliengk/go-utils/validation"
	"github.com/kassisol/hbm/storage"
	"github.com/spf13/cobra"
)

var (
	policyListFilter []string

	policyAddGroup      string
	policyAddCluster    string
	policyAddCollection string
)

var policyCmd = &cobra.Command{
	Use:   "policy",
	Short: "Manage whitelisted policies",
	Long:  "Manage whitelisted policies",
}

var policyListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List whitelisted policies",
	Long:  "List whitelisted policies",
}

var policyAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add policy to the whitelist",
	Long:  "Add policy to the whitelist",
}

var policyRemoveCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove policy from the whitelist",
	Long:  "Remove policy from the whitelist",
}

var policyExistsCmd = &cobra.Command{
	Use:   "find",
	Short: "Verify if policy exists in the whitelist",
	Long:  "Verify if policy exists in the whitelist",
}

func init() {
	RootCmd.AddCommand(policyCmd)
	policyCmd.AddCommand(policyListCmd)
	policyCmd.AddCommand(policyAddCmd)
	policyCmd.AddCommand(policyRemoveCmd)
	policyCmd.AddCommand(policyExistsCmd)

	policyListCmd.Flags().StringSliceVarP(&policyListFilter, "filter", "f", []string{}, "Filter")

	policyAddCmd.Flags().StringVarP(&policyAddGroup, "group", "g", "", "Set group")
	policyAddCmd.Flags().StringVarP(&policyAddCluster, "cluster", "c", "", "Set cluster")
	policyAddCmd.Flags().StringVarP(&policyAddCollection, "collection", "", "", "Set collection")

	policyCmd.Run = policyUsage
	policyListCmd.Run = policyList
	policyAddCmd.Run = policyAdd
	policyRemoveCmd.Run = policyRemove
	policyExistsCmd.Run = policyExists
}

func policyUsage(cmd *cobra.Command, args []string) {
	policyCmd.Usage()
	os.Exit(-1)
}

func policyList(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", appPath)
	if err != nil {
		utils.Exit(err)
	}
	defer s.End()

	filters := utils.ConvertSliceToMap("=", policyListFilter)

	if err = IsValidPolicyFilterKeys(filters); err != nil {
		utils.Exit(err)
	}

	policies := s.ListPolicies(filters)

	if len(policies) > 0 {
		w := tabwriter.NewWriter(os.Stdout, 20, 1, 2, ' ', 0)
		fmt.Fprintln(w, "NAME\tGROUP\tCLUSTER\tCOLLECTION")

		for _, policy := range policies {
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", policy.Name, policy.Group, policy.Cluster, policy.Collection)
		}

		w.Flush()
	}
}

func policyAdd(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", appPath)
	if err != nil {
		utils.Exit(err)
	}
	defer s.End()

	if len(args) < 1 || len(args) > 1 {
		cmd.Usage()
		os.Exit(-1)
	}

	if err = validation.IsValidName(args[0]); err != nil {
		utils.Exit(err)
	}

	if s.FindPolicy(args[0]) {
		utils.Exit(fmt.Errorf("%s already exists", args[0]))
	}

	if !s.FindGroup(policyAddGroup) {
		utils.Exit(fmt.Errorf("%s does not exist", policyAddGroup))
	}

	if !s.FindCluster(policyAddCluster) {
		utils.Exit(fmt.Errorf("%s does not exist", policyAddCluster))
	}

	if !s.FindCollection(policyAddCollection) {
		utils.Exit(fmt.Errorf("%s does not exist", policyAddCollection))
	}

	s.AddPolicy(args[0], policyAddGroup, policyAddCluster, policyAddCollection)
}

func policyRemove(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", appPath)
	if err != nil {
		utils.Exit(err)
	}
	defer s.End()

	if len(args) < 1 || len(args) > 1 {
		cmd.Usage()
		os.Exit(-1)
	}

	if !s.FindPolicy(args[0]) {
		utils.Exit(fmt.Errorf("%s does not exist", args[0]))
	}

	s.RemovePolicy(args[0])
}

func policyExists(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", appPath)
	if err != nil {
		utils.Exit(err)
	}
	defer s.End()

	if len(args) < 1 || len(args) > 1 {
		cmd.Usage()
		os.Exit(-1)
	}

	result := s.FindPolicy(args[0])

	fmt.Println(result)
}
