package user

import (
	"github.com/juliengk/go-utils"
	"github.com/kassisol/hbm/cli/command"
	userobj "github.com/kassisol/hbm/object/user"
	log "github.com/sirupsen/logrus"
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

	u, err := userobj.New("sqlite", command.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer u.End()

	if userMemberAdd {
		if err := u.AddToGroup(args[1], args[0]); err != nil {
			log.Fatal(err)
		}
	}
	if userMemberRemove {
		if err := u.RemoveFromGroup(args[1], args[0]); err != nil {
			log.Fatal(err)
		}
	}
}

var memberDescription = `
Manage user membership to group

`
