package cmd

import (
	"fmt"
	log "github.com/Sirupsen/logrus"

	"github.com/harbourmaster/hbm/api"
	"github.com/harbourmaster/hbm/pkg/db"
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
	data, err := api.NewApi(dockerApiVersion, "")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%-25s | %-20s | %s\n", "Action", "Command Name", "Description")
		fmt.Println("-----------------------------------------------------------------------------------------------------------------------")

		for _, a := range *data.Uris {
			fmt.Printf("%-25s | %-20s | %s\n", a.Action, a.CmdName, a.Description)
		}
	}
}

func actionList(cmd *cobra.Command, args []string) {
	defer db.RecoverFunc()
	bucket := getBucket("action")

	d, err := db.NewDB(appPath)
	if err != nil {
		log.Fatal(err)
	}

	actions := d.List(bucket)

	if len(actions) > 0 {
		fmt.Println("Action")
		fmt.Println("------")
	}

	for _, v := range actions {
		fmt.Printf("%s\n", v.Key)
	}
}

func validAction(action string) bool {
	data, err := api.NewApi(dockerApiVersion, "")
	if err != nil {
		fmt.Println(err)
		return false
	}
	for _, a := range *data.Uris {
		if action == a.Action {
			return true
		}
	}
	return false
}

func actionAdd(cmd *cobra.Command, args []string) {
	action := args[0]
	if validAction(action) {

		defer db.RecoverFunc()
		bucket := getBucket("action")

		d, err := db.NewDB(appPath)
		if err != nil {
			log.Fatal(err)
		}

		d.Add(bucket, action, []byte(""))
	} else {
		fmt.Printf("Error: action %s not valid for docker API %s\n", action, dockerApiVersion)
	}
}

func actionRemove(cmd *cobra.Command, args []string) {
	defer db.RecoverFunc()
	bucket := getBucket("action")

	d, err := db.NewDB(appPath)
	if err != nil {
		log.Fatal(err)
	}

	d.Remove(bucket, args[0])
}
