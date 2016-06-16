package cmd

import (
	"fmt"
	"log"

	"github.com/harbourmaster/hbm/pkg/db"
        "github.com/spf13/cobra"
)

var deviceCmd = &cobra.Command{
        Use:    "device",
        Short:  "Manage whitelisted devices",
        Long:	"Manage whitelisted devices",
}

var deviceListCmd = &cobra.Command{
        Use:    "ls",
        Short:  "List whitelisted devices",
        Long:	"List whitelisted devices",
}

var deviceAddCmd = &cobra.Command{
        Use:    "add",
        Short:  "Add device to the whitelist",
        Long:	"Add device to the whitelist",
}

var deviceRemoveCmd = &cobra.Command{
        Use:    "rm",
        Short:  "Remove device from the whitelist",
        Long:	"Remove device from the whitelist",
}

func init() {
        RootCmd.AddCommand(deviceCmd)
	deviceCmd.AddCommand(deviceListCmd)
	deviceCmd.AddCommand(deviceAddCmd)
	deviceCmd.AddCommand(deviceRemoveCmd)

        deviceCmd.Run = deviceList
        deviceListCmd.Run = deviceList
        deviceAddCmd.Run = deviceAdd
        deviceRemoveCmd.Run = deviceRemove
}

func deviceList(cmd *cobra.Command, args []string) {
	defer db.RecoverFunc()

	d, err := db.NewDB(appPath)
	if err != nil {
		log.Fatal(err)
	}

	devices := d.List("device")

	if len(devices) > 0 {
		fmt.Println("Device")
		fmt.Println("------")
	}

	for _, v := range devices {
		fmt.Printf("%s\n", v.Key)
	}
}

func deviceAdd(cmd *cobra.Command, args []string) {
	defer db.RecoverFunc()

	d, err := db.NewDB(appPath)
	if err != nil {
		log.Fatal(err)
	}

	d.Add("device", args[0], []byte(""))
}

func deviceRemove(cmd *cobra.Command, args []string) {
	defer db.RecoverFunc()

	d, err := db.NewDB(appPath)
	if err != nil {
		log.Fatal(err)
	}

	d.Remove("device", args[0])
}
