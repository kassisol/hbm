package cmd

import (
	"fmt"
	"log"

	"github.com/kassisol/hbm/plugin"
	"github.com/kassisol/hbm/pkg/db"
	"github.com/spf13/cobra"
)

var actionCmd = &cobra.Command{
	Use:   "action",
	Short: "Manage whitelisted actions",
	Long:  "Manage whitelisted actions",
}

var actionListOptionsCmd = &cobra.Command{
	Use:   "options",
	Short: "List options actions",
	Long:  "List options actions",
}

var actionListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List whitelisted actions",
	Long:  "List whitelisted actions",
}

var actionAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add action to the whitelist",
	Long:  "Add action to the whitelist",
}

var actionRemoveCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove action from the whitelist",
	Long:  "Remove action from the whitelist",
}

func init() {
	RootCmd.AddCommand(actionCmd)
	actionCmd.AddCommand(actionListOptionsCmd)
	actionCmd.AddCommand(actionListCmd)
	actionCmd.AddCommand(actionAddCmd)
	actionCmd.AddCommand(actionRemoveCmd)

	actionCmd.Run = actionList
	actionListOptionsCmd.Run = actionListOptions
	actionListCmd.Run = actionList
	actionAddCmd.Run = actionAdd
	actionRemoveCmd.Run = actionRemove
}

func actionListOptions(cmd *cobra.Command, args []string) {
	data, _ := plugin.NewApi(plugin.SupportedVersion, "")

	fmt.Printf("%-25s | %-20s | %s\n", "Action", "Command Name", "Description")
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------")

	for _, a := range *data.Uris {
		fmt.Printf("%-25s | %-20s | %s\n", a.Action, a.CmdName, a.Description)
	}
}

func actionList(cmd *cobra.Command, args []string) {
	defer db.RecoverFunc()

	d, err := db.NewDB(appPath)
	if err != nil {
		log.Fatal(err)
	}

	actions := d.List("action")

	if len(actions) > 0 {
		fmt.Println("Action")
		fmt.Println("------")
	}

	for _, v := range actions {
		fmt.Printf("%s\n", v.Key)
	}
}

func actionAdd(cmd *cobra.Command, args []string) {
	defer db.RecoverFunc()

	d, err := db.NewDB(appPath)
	if err != nil {
		log.Fatal(err)
	}

	d.Add("action", args[0], []byte(""))
}

func actionRemove(cmd *cobra.Command, args []string) {
	defer db.RecoverFunc()

	d, err := db.NewDB(appPath)
	if err != nil {
		log.Fatal(err)
	}

	d.Remove("action", args[0])
}
