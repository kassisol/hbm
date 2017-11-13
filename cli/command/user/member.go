package user

import (
	log "github.com/sirupsen/logrus"
	"github.com/juliengk/go-utils"
	"github.com/kassisol/hbm/cli/command"
	"github.com/kassisol/hbm/storage"
	"github.com/spf13/cobra"
)

var (
	userMemberAdd    bool
	userMemberRemove bool
)

func newMemberCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "member [group] [user]",
		Short: "Manage user membership to group",
		Long:  memberDescription,
		Args:  cobra.ExactArgs(2),
		Run:   runMember,
	}

	flags := cmd.Flags()
	flags.BoolVarP(&userMemberAdd, "add", "a", false, "Add user to group")
	flags.BoolVarP(&userMemberRemove, "remove", "r", false, "Remove user from group")

	return cmd
}

func runMember(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", command.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer s.End()

	if !s.FindGroup(args[0]) {
		log.Fatalf("%s does not exist", args[0])
	}

	if !s.FindUser(args[1]) {
		log.Fatalf("%s does not exist", args[1])
	}

	if userMemberAdd {
		s.AddUserToGroup(args[0], args[1])
	}
	if userMemberRemove {
		s.RemoveUserFromGroup(args[0], args[1])
	}
}

var memberDescription = `
Manage user membership to group

`
