package cmd

import (
	"fmt"
	"log"

	"github.com/kassisol/hbm/pkg/db"
	"github.com/spf13/cobra"
)

var registryCmd = &cobra.Command{
	Use:   "registry",
	Short: "Manage whitelisted registries",
	Long:  "Manage whitelisted registries",
}

var registryListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List whitelisted registries",
	Long:  "List whitelisted registries",
}

var registryAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add registry to the whitelist",
	Long:  "Add registry to the whitelist",
}

var registryRemoveCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove registry from the whitelist",
	Long:  "Remove registry from the whitelist",
}

func init() {
	RootCmd.AddCommand(registryCmd)
	registryCmd.AddCommand(registryListCmd)
	registryCmd.AddCommand(registryAddCmd)
	registryCmd.AddCommand(registryRemoveCmd)

	registryCmd.Run = registryList
	registryListCmd.Run = registryList
	registryAddCmd.Run = registryAdd
	registryRemoveCmd.Run = registryRemove
}

func registryList(cmd *cobra.Command, args []string) {
	defer db.RecoverFunc()

	d, err := db.NewDB(appPath)
	if err != nil {
		log.Fatal(err)
	}

	registries := d.List("registry")

	if len(registries) > 0 {
		fmt.Println("Registry")
		fmt.Println("--------")
	}

	for _, v := range registries {
		fmt.Println(v.Key)
	}
}

func registryAdd(cmd *cobra.Command, args []string) {
	defer db.RecoverFunc()

	d, err := db.NewDB(appPath)
	if err != nil {
		log.Fatal(err)
	}

	d.Add("registry", args[0], []byte(""))
}

func registryRemove(cmd *cobra.Command, args []string) {
	defer db.RecoverFunc()

	d, err := db.NewDB(appPath)
	if err != nil {
		log.Fatal(err)
	}

	d.Remove("registry", args[0])
}
