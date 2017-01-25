package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/juliengk/go-utils"
	"github.com/kassisol/hbm/config"
	"github.com/kassisol/hbm/storage"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage HBM features",
	Long:  "Manage HBM features",
}

var configListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List hbm enabled features",
	Long:  "List hbm enabled features",
}

var configFindCmd = &cobra.Command{
	Use:   "find",
	Short: "Verify if feature is enabled",
	Long:  "Verify if feature is enabled",
}

var configAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Enable hbm feature",
	Long:  "Enable hbm feature",
}

var configRemoveCmd = &cobra.Command{
	Use:   "rm",
	Short: "Disable hbm feature",
	Long:  "Disable hbm feature",
}

func init() {
	RootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configListCmd)
	configCmd.AddCommand(configFindCmd)
	configCmd.AddCommand(configAddCmd)
	configCmd.AddCommand(configRemoveCmd)

	configCmd.Run = configList
	configListCmd.Run = configList
	configFindCmd.Run = configFind
	configAddCmd.Run = configAdd
	configRemoveCmd.Run = configRemove
}

func configList(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", appPath)
	if err != nil {
		utils.Exit(err)
	}
	defer s.End()

	configs := s.ListConfigs()

	if len(configs) > 0 {
		w := tabwriter.NewWriter(os.Stdout, 20, 1, 2, ' ', 0)
		fmt.Fprintln(w, "NAME")

		for _, config := range configs {
			fmt.Fprintf(w, "%s\n", config)
		}

		w.Flush()
	}
}

func configFind(cmd *cobra.Command, args []string) {
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

	result := s.FindConfig(args[0])

	fmt.Println(result)
}

func configAdd(cmd *cobra.Command, args []string) {
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

	configs := config.New()
	if configs.ConfigExists(args[0]) {
		utils.Exit(fmt.Errorf("This feature does not exist"))
	}

	if s.FindConfig(args[0]) {
		utils.Exit(fmt.Errorf("%s is already enabled", args[0]))
	}

	s.AddConfig(args[0])
}

func configRemove(cmd *cobra.Command, args []string) {
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

	if !s.FindConfig(args[0]) {
		utils.Exit(fmt.Errorf("%s is not enabled", args[0]))
	}

	if err = s.RemoveConfig(args[0]); err != nil {
		utils.Exit(err)
	}
}
