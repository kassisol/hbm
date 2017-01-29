package host

import (
	"fmt"
	"os"

	"github.com/juliengk/go-utils"
	"github.com/kassisol/hbm/cli/command"
	"github.com/kassisol/hbm/storage"
	"github.com/spf13/cobra"
)

var (
	hostMemberAdd    bool
	hostMemberRemove bool
)

func newMemberCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "member [cluster] [ host]",
		Short: "Manage host membership to cluster",
		Long:  memberDescription,
		Run:   runMember,
	}

	flags := cmd.Flags()
	flags.BoolVarP(&hostMemberAdd, "add", "a", false, "Add host to cluster")
	flags.BoolVarP(&hostMemberRemove, "remove", "r", false, "Remove host from cluster")

	return cmd
}

func runMember(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", command.AppPath)
	if err != nil {
		utils.Exit(err)
	}
	defer s.End()

	if len(args) < 2 || len(args) > 2 {
		cmd.Usage()
		os.Exit(-1)
	}

	if !s.FindCluster(args[0]) {
		utils.Exit(fmt.Errorf("%s does not exist", args[0]))
	}

	if !s.FindHost(args[1]) {
		utils.Exit(fmt.Errorf("%s does not exist", args[1]))
	}

	if hostMemberAdd {
		s.AddHostToCluster(args[0], args[1])
	}
	if hostMemberRemove {
		s.RemoveHostFromCluster(args[0], args[1])
	}
}

var memberDescription = `
Manage host membership to cluster

`
