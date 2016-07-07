package cmd

import (
	"fmt"
	"log"

	"github.com/harbourmaster/hbm/pkg/db"
	"github.com/spf13/cobra"
)

var portCmd = &cobra.Command{
	Use:   "port",
	Short: "Manage whitelisted ports",
	Long:  "Manage whitelisted ports",
}

var portListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List whitelisted ports",
	Long:  "List whitelisted ports",
}

var portAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add port to the whitelist",
	Long:  "Add port to the whitelist",
}

var portRemoveCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove port from the whitelist",
	Long:  "Remove port from the whitelist",
}

func init() {
	RootCmd.AddCommand(portCmd)
	portCmd.AddCommand(portListCmd)
	portCmd.AddCommand(portAddCmd)
	portCmd.AddCommand(portRemoveCmd)

	portCmd.Run = portList
	portListCmd.Run = portList
	portAddCmd.Run = portAdd
	portRemoveCmd.Run = portRemove
}

func portList(cmd *cobra.Command, args []string) {
	defer db.RecoverFunc()

	d, err := db.NewDB(appPath)
	if err != nil {
		log.Fatal(err)
	}

	ports := d.List("port")

	if len(ports) > 0 {
		fmt.Println("Port")
		fmt.Println("----")
	}

	for _, v := range ports {
		fmt.Printf("%s\n", v.Key)
	}
}

func portAdd(cmd *cobra.Command, args []string) {
	defer db.RecoverFunc()

	d, err := db.NewDB(appPath)
	if err != nil {
		log.Fatal(err)
	}

	d.Add("port", args[0], []byte(""))
}

func portRemove(cmd *cobra.Command, args []string) {
	defer db.RecoverFunc()

	d, err := db.NewDB(appPath)
	if err != nil {
		log.Fatal(err)
	}

	d.Remove("port", args[0])
}
