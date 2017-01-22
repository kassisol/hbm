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

var collectionCmd = &cobra.Command{
	Use:   "collection",
	Short: "Manage whitelisted collections",
	Long:  "Manage whitelisted collections",
}

var collectionListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List whitelisted collections",
	Long:  "List whitelisted collections",
}

var collectionAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add collection to the whitelist",
	Long:  "Add collection to the whitelist",
}

var collectionRemoveCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove collection from the whitelist",
	Long:  "Remove collection from the whitelist",
}

var collectionExistsCmd = &cobra.Command{
	Use:   "find",
	Short: "Verify if collection exists in the whitelist",
	Long:  "Verify if collection exists in the whitelist",
}

func init() {
	RootCmd.AddCommand(collectionCmd)
	collectionCmd.AddCommand(collectionListCmd)
	collectionCmd.AddCommand(collectionAddCmd)
	collectionCmd.AddCommand(collectionRemoveCmd)
	collectionCmd.AddCommand(collectionExistsCmd)

	collectionCmd.Run = collectionUsage
	collectionListCmd.Run = collectionList
	collectionAddCmd.Run = collectionAdd
	collectionRemoveCmd.Run = collectionRemove
	collectionExistsCmd.Run = collectionExists
}

func collectionUsage(cmd *cobra.Command, args []string) {
	collectionCmd.Usage()
	os.Exit(-1)
}
func collectionList(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", appPath)
	if err != nil {
		utils.Exit(err)
	}
	defer s.End()

	collections := s.ListCollections()

	if len(collections) > 0 {
		w := tabwriter.NewWriter(os.Stdout, 20, 1, 2, ' ', 0)
		fmt.Fprintln(w, "NAME\tRESOURCES")

		for collection, resources := range collections {
			if len(resources) > 0 {
				fmt.Fprintf(w, "%s\t%s\n", collection, strings.Join(resources, ", "))
			} else {
				fmt.Fprintf(w, "%s\n", collection)
			}
		}

		w.Flush()
	}
}

func collectionAdd(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", appPath)
	if err != nil {
		utils.Exit(err)
	}
	defer s.End()

	if err = validation.IsValidName(args[0]); err != nil {
		utils.Exit(err)
	}

	s.AddCollection(args[0])
}

func collectionRemove(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", appPath)
	if err != nil {
		utils.Exit(err)
	}
	defer s.End()

	if !s.FindCollection(args[0]) {
		utils.Exit(fmt.Errorf("%s does not exist", args[0]))
	}

	s.RemoveCollection(args[0])
}

func collectionExists(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", appPath)
	if err != nil {
		utils.Exit(err)
	}
	defer s.End()

	result := s.FindCollection(args[0])

	fmt.Println(result)
}
