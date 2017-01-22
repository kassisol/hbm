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

var userMemberAdd bool
var userMemberRemove bool

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage whitelisted users",
	Long:  "Manage whitelisted users",
}

var userListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List whitelisted users",
	Long:  "List whitelisted users",
}

var userAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add user to the whitelist",
	Long:  "Add user to the whitelist",
}

var userRemoveCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove user from the whitelist",
	Long:  "Remove user from the whitelist",
}

var userExistsCmd = &cobra.Command{
	Use:   "find",
	Short: "Verify if user exists in the whitelist",
	Long:  "Verify if user exists in the whitelist",
}

var userMemberCmd = &cobra.Command{
	Use:   "member",
	Short: "Verify if user exists in the whitelist",
	Long:  "Verify if user exists in the whitelist",
}

func init() {
	RootCmd.AddCommand(userCmd)
	userCmd.AddCommand(userListCmd)
	userCmd.AddCommand(userAddCmd)
	userCmd.AddCommand(userRemoveCmd)
	userCmd.AddCommand(userExistsCmd)
	userCmd.AddCommand(userMemberCmd)

	userMemberCmd.Flags().BoolVarP(&userMemberAdd, "add", "a", false, "Add user to group")
	userMemberCmd.Flags().BoolVarP(&userMemberRemove, "remove", "r", false, "Remove user to group")

	userCmd.Run = userUsage
	userListCmd.Run = userList
	userAddCmd.Run = userAdd
	userRemoveCmd.Run = userRemove
	userExistsCmd.Run = userExists
	userMemberCmd.Run = userMember
}

func userUsage(cmd *cobra.Command, args []string) {
	userCmd.Usage()
	os.Exit(-1)
}

func userList(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", appPath)
	if err != nil {
		utils.Exit(err)
	}
	defer s.End()

	users := s.ListUsers()

	if len(users) > 0 {
		w := tabwriter.NewWriter(os.Stdout, 20, 1, 2, ' ', 0)
		fmt.Fprintln(w, "NAME\tGROUPS")

		for user, groups := range users {
			if len(groups) > 0 {
				fmt.Fprintf(w, "%s\t%s\n", user, strings.Join(groups, ", "))
			} else {
				fmt.Fprintf(w, "%s\n", user)
			}
		}

		w.Flush()
	}
}

func userAdd(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", appPath)
	if err != nil {
		utils.Exit(err)
	}
	defer s.End()

	if err = validation.IsValidUsername(args[0]); err != nil {
		utils.Exit(err)
	}

	if s.FindUser(args[0]) {
		utils.Exit(fmt.Errorf("%s already exists", args[0]))
	}

	s.AddUser(args[0])
}

func userRemove(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", appPath)
	if err != nil {
		utils.Exit(err)
	}
	defer s.End()

	if !s.FindUser(args[0]) {
		utils.Exit(fmt.Errorf("%s does not exist", args[0]))
	}

	s.RemoveUser(args[0])
}

func userExists(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", appPath)
	if err != nil {
		utils.Exit(err)
	}
	defer s.End()

	result := s.FindUser(args[0])

	fmt.Println(result)
}

func userMember(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", appPath)
	if err != nil {
		utils.Exit(err)
	}
	defer s.End()

	if !s.FindGroup(args[0]) {
		utils.Exit(fmt.Errorf("%s does not exist", args[0]))
	}

	if !s.FindUser(args[1]) {
		utils.Exit(fmt.Errorf("%s does not exist", args[1]))
	}

	if userMemberAdd {
		s.AddUserToGroup(args[0], args[1])
	}
	if userMemberRemove {
		s.RemoveUserFromGroup(args[0], args[1])
	}
}
