package cmd

import (
	"fmt"
	"log"

	"github.com/kassisol/hbm/docker/config"
	"github.com/kassisol/hbm/pkg/db"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage whitelisted configs",
	Long:  "Manage whitelisted configs",
}

var configListOptionsCmd = &cobra.Command{
	Use:   "options",
	Short: "List options configs",
	Long:  "List options configs",
}

var configListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List whitelisted configs",
	Long:  "List whitelisted configs",
}

var configAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add config to the whitelist",
	Long:  "Add config to the whitelist",
}

var configRemoveCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove config from the whitelist",
	Long:  "Remove config from the whitelist",
}

func init() {
	RootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configListOptionsCmd)
	configCmd.AddCommand(configListCmd)
	configCmd.AddCommand(configAddCmd)
	configCmd.AddCommand(configRemoveCmd)

	configCmd.Run = configList
	configListOptionsCmd.Run = configListOptions
	configListCmd.Run = configList
	configAddCmd.Run = configAdd
	configRemoveCmd.Run = configRemove
}

func configListOptions(cmd *cobra.Command, args []string) {
	data := config.New()

	fmt.Printf("%-30s | %s\n", "Action", "Description")
	fmt.Println("-------------------------------------------------------")

	for _, c := range data {
		fmt.Printf("%-30s | %s\n", c.Action, c.Description)
	}
}

func configList(cmd *cobra.Command, args []string) {
	defer db.RecoverFunc()

	d, err := db.NewDB(appPath)
	if err != nil {
		log.Fatal(err)
	}

	configs := d.List("config")

	if len(configs) > 0 {
		fmt.Println("Config")
		fmt.Println("------")
	}

	for _, v := range configs {
		fmt.Printf("%s\n", v.Key)
	}
}

func configAdd(cmd *cobra.Command, args []string) {
	defer db.RecoverFunc()

	d, err := db.NewDB(appPath)
	if err != nil {
		log.Fatal(err)
	}

	d.Add("config", args[0], []byte(""))
}

func configRemove(cmd *cobra.Command, args []string) {
	defer db.RecoverFunc()

	d, err := db.NewDB(appPath)
	if err != nil {
		log.Fatal(err)
	}

	d.Remove("config", args[0])
}
