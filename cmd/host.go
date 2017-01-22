package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/juliengk/go-utils"
	"github.com/kassisol/hbm/storage"
	"github.com/spf13/cobra"
)

var hostMemberAdd bool
var hostMemberRemove bool

var hostCmd = &cobra.Command{
	Use:   "host",
	Short: "Manage whitelisted hosts",
	Long:  "Manage whitelisted hosts",
}

var hostListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List whitelisted hosts",
	Long:  "List whitelisted hosts",
}

var hostAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add host to the whitelist",
	Long:  "Add host to the whitelist",
}

var hostRemoveCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove host from the whitelist",
	Long:  "Remove host from the whitelist",
}

var hostExistsCmd = &cobra.Command{
	Use:   "find",
	Short: "Verify if host exists in the whitelist",
	Long:  "Verify if host exists in the whitelist",
}

var hostMemberCmd = &cobra.Command{
	Use:   "member",
	Short: "Verify if host exists in the whitelist",
	Long:  "Verify if host exists in the whitelist",
}

func init() {
	RootCmd.AddCommand(hostCmd)
	hostCmd.AddCommand(hostListCmd)
	hostCmd.AddCommand(hostAddCmd)
	hostCmd.AddCommand(hostRemoveCmd)
	hostCmd.AddCommand(hostExistsCmd)
	hostCmd.AddCommand(hostMemberCmd)

	hostMemberCmd.Flags().BoolVarP(&hostMemberAdd, "add", "a", false, "Add host to group")
	hostMemberCmd.Flags().BoolVarP(&hostMemberRemove, "remove", "r", false, "Remove host to group")

	hostCmd.Run = hostUsage
	hostListCmd.Run = hostList
	hostAddCmd.Run = hostAdd
	hostRemoveCmd.Run = hostRemove
	hostExistsCmd.Run = hostExists
	hostMemberCmd.Run = hostMember
}

func hostUsage(cmd *cobra.Command, args []string) {
	hostCmd.Usage()
	os.Exit(-1)
}

func hostList(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", appPath)
	if err != nil {
		utils.Exit(err)
	}
	defer s.End()

	hosts := s.ListHosts()

	if len(hosts) > 0 {
		w := tabwriter.NewWriter(os.Stdout, 20, 1, 2, ' ', 0)
		fmt.Fprintln(w, "NAME")

		for _, host := range hosts {
			fmt.Fprintf(w, "%s\n", host)
		}

		w.Flush()
	}
}

func hostAdd(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", appPath)
	if err != nil {
		utils.Exit(err)
	}
	defer s.End()

	//if err = utils.IsValidHostname(args[0]); err != nil {
	//      utils.Exit(err)
	//}

	s.AddHost(args[0])
}

func hostRemove(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", appPath)
	if err != nil {
		utils.Exit(err)
	}
	defer s.End()

	if !s.FindHost(args[0]) {
		utils.Exit(fmt.Errorf("%s does not exist", args[0]))
	}

	s.RemoveHost(args[0])
}

func hostExists(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", appPath)
	if err != nil {
		utils.Exit(err)
	}
	defer s.End()

	result := s.FindHost(args[0])

	fmt.Println(result)
}

func hostMember(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", appPath)
	if err != nil {
		utils.Exit(err)
	}
	defer s.End()

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
