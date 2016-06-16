package cmd

import (
	"fmt"
	"log"

	"github.com/harbourmaster/hbm/pkg/db"
        "github.com/spf13/cobra"
)

var capCmd = &cobra.Command{
        Use:    "cap",
        Short:  "Manage whitelisted caps",
        Long:	"Manage whitelisted caps",
}

var capListCmd = &cobra.Command{
        Use:    "ls",
        Short:  "List whitelisted caps",
        Long:	"List whitelisted caps",
}

var capAddCmd = &cobra.Command{
        Use:    "add",
        Short:  "Add cap to the whitelist",
        Long:	"Add cap to the whitelist",
}

var capRemoveCmd = &cobra.Command{
        Use:    "rm",
        Short:  "Remove cap from the whitelist",
        Long:	"Remove cap from the whitelist",
}

func init() {
        RootCmd.AddCommand(capCmd)
	capCmd.AddCommand(capListCmd)
	capCmd.AddCommand(capAddCmd)
	capCmd.AddCommand(capRemoveCmd)

        capCmd.Run = capList
        capListCmd.Run = capList
        capAddCmd.Run = capAdd
        capRemoveCmd.Run = capRemove
}

func capList(cmd *cobra.Command, args []string) {
	defer db.RecoverFunc()

	d, err := db.NewDB(appPath)
	if err != nil {
		log.Fatal(err)
	}

	caps := d.List("cap")

	if len(caps) > 0 {
		fmt.Println("Capability")
		fmt.Println("----------")
	}

	for _, v := range caps {
		fmt.Printf("%s\n", v.Key)
	}
}

func capAdd(cmd *cobra.Command, args []string) {
	defer db.RecoverFunc()

	d, err := db.NewDB(appPath)
	if err != nil {
		log.Fatal(err)
	}

	d.Add("cap", args[0], []byte(""))
}

func capRemove(cmd *cobra.Command, args []string) {
	defer db.RecoverFunc()

	d, err := db.NewDB(appPath)
	if err != nil {
		log.Fatal(err)
	}

	d.Remove("cap", args[0])
}
