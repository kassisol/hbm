package resource

import (
	"github.com/juliengk/go-utils"
	"github.com/kassisol/hbm/cli/command"
	"github.com/kassisol/hbm/storage"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	resourceMemberAdd    bool
	resourceMemberRemove bool
)

func newMemberCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "member [collection] [resource]",
		Short: "Manage resource membership to collection",
		Long:  memberDescription,
		Args:  cobra.ExactArgs(2),
		Run:   runMember,
	}

	flags := cmd.Flags()
	flags.BoolVarP(&resourceMemberAdd, "add", "a", false, "Add resource to collection")
	flags.BoolVarP(&resourceMemberRemove, "remove", "r", false, "Remove resource from collection")

	return cmd
}

func runMember(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", command.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer s.End()

	if !s.FindCollection(args[0]) {
		log.Fatalf("%s does not exist", args[0])
	}

	if !s.FindResource(args[1]) {
		log.Fatalf("%s does not exist", args[1])
	}

	if resourceMemberAdd {
		s.AddResourceToCollection(args[0], args[1])
	}
	if resourceMemberRemove {
		s.RemoveResourceFromCollection(args[0], args[1])
	}
}

var memberDescription = `
Manage resource membership to collection

`
