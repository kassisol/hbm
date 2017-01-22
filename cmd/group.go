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

var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "Manage whitelisted groups",
	Long:  "Manage whitelisted groups",
}

var groupListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List whitelisted groups",
	Long:  "List whitelisted groups",
}

var groupAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add group to the whitelist",
	Long:  "Add group to the whitelist",
}

var groupRemoveCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove group from the whitelist",
	Long:  "Remove group from the whitelist",
}

var groupExistsCmd = &cobra.Command{
	Use:   "find",
	Short: "Verify if group exists in the whitelist",
	Long:  "Verify if group exists in the whitelist",
}

func init() {
	RootCmd.AddCommand(groupCmd)
	groupCmd.AddCommand(groupListCmd)
	groupCmd.AddCommand(groupAddCmd)
	groupCmd.AddCommand(groupRemoveCmd)
	groupCmd.AddCommand(groupExistsCmd)

	groupCmd.Run = groupUsage
	groupListCmd.Run = groupList
	groupAddCmd.Run = groupAdd
	groupRemoveCmd.Run = groupRemove
	groupExistsCmd.Run = groupExists
}

func groupUsage(cmd *cobra.Command, args []string) {
	groupCmd.Usage()
	os.Exit(-1)
}

func groupList(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", appPath)
	if err != nil {
		utils.Exit(err)
	}
	defer s.End()

	groups := s.ListGroups()

	if len(groups) > 0 {
		w := tabwriter.NewWriter(os.Stdout, 20, 1, 2, ' ', 0)
		fmt.Fprintln(w, "NAME\tUSERS")

		for group, users := range groups {
			if len(users) > 0 {
				fmt.Fprintf(w, "%s\t%s\n", group, strings.Join(users, ", "))
			} else {
				fmt.Fprintf(w, "%s\n", group)
			}
		}

		w.Flush()
	}
}

func groupAdd(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", appPath)
	if err != nil {
		utils.Exit(err)
	}
	defer s.End()

	if err = validation.IsValidGroupname(args[0]); err != nil {
		utils.Exit(err)
	}

	if s.FindGroup(args[0]) {
		utils.Exit(fmt.Errorf("%s already exists", args[0]))
	}

	s.AddGroup(args[0])
}

func groupRemove(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", appPath)
	if err != nil {
		utils.Exit(err)
	}
	defer s.End()

	if !s.FindGroup(args[0]) {
		utils.Exit(fmt.Errorf("%s does not exist", args[0]))
	}

	s.RemoveGroup(args[0])
}

func groupExists(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", appPath)
	if err != nil {
		utils.Exit(err)
	}
	defer s.End()

	result := s.FindGroup(args[0])

	fmt.Println(result)
}
