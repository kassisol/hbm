package cmd

import (
	"fmt"
	"log"

	"github.com/harbourmaster/hbm/pkg/db"
        "github.com/spf13/cobra"
)

var dnsCmd = &cobra.Command{
        Use:    "dns",
        Short:  "Manage whitelisted DNS server",
        Long:	"Manage whitelisted DNS server",
}

var dnsListCmd = &cobra.Command{
        Use:    "ls",
        Short:  "List whitelisted DNS server",
        Long:	"List whitelisted DNS server",
}

var dnsAddCmd = &cobra.Command{
        Use:    "add",
        Short:  "Add DNS server to the whitelist",
        Long:	"Add DNS server to the whitelist",
}

var dnsRemoveCmd = &cobra.Command{
        Use:    "rm",
        Short:  "Remove DNS server from the whitelist",
        Long:	"Remove DNS server from the whitelist",
}

func init() {
        RootCmd.AddCommand(dnsCmd)
	dnsCmd.AddCommand(dnsListCmd)
	dnsCmd.AddCommand(dnsAddCmd)
	dnsCmd.AddCommand(dnsRemoveCmd)

        dnsCmd.Run = dnsList
        dnsListCmd.Run = dnsList
        dnsAddCmd.Run = dnsAdd
        dnsRemoveCmd.Run = dnsRemove
}

func dnsList(cmd *cobra.Command, args []string) {
	defer db.RecoverFunc()

	d, err := db.NewDB(appPath)
	if err != nil {
		log.Fatal(err)
	}

	dnsservers := d.List("dns")

	if len(dnsservers) > 0 {
		fmt.Println("DNS Server")
		fmt.Println("----------")
	}

	for _, v := range dnsservers {
		fmt.Printf("%s\n", v.Key)
	}
}

func dnsAdd(cmd *cobra.Command, args []string) {
	defer db.RecoverFunc()

	d, err := db.NewDB(appPath)
	if err != nil {
		log.Fatal(err)
	}

	d.Add("dns", args[0], []byte(""))
}

func dnsRemove(cmd *cobra.Command, args []string) {
	defer db.RecoverFunc()

	d, err := db.NewDB(appPath)
	if err != nil {
		log.Fatal(err)
	}

	d.Remove("dns", args[0])
}
