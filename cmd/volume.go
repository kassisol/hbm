package cmd

import (
	"fmt"
	"log"

	"github.com/kassisol/hbm/pkg/db"
	"github.com/spf13/cobra"
)

var volumeRecursive bool

var volumeCmd = &cobra.Command{
	Use:   "volume",
	Short: "Manage whitelisted volumes",
	Long:  "Manage whitelisted volumes",
}

var volumeListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List whitelisted volumes",
	Long:  "List whitelisted volumes",
}

var volumeAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add volume to the whitelist",
	Long:  "Add volume to the whitelist",
}

var volumeRemoveCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove volume from the whitelist",
	Long:  "Remove volume from the whitelist",
}

func init() {
	RootCmd.AddCommand(volumeCmd)
	volumeCmd.AddCommand(volumeListCmd)
	volumeCmd.AddCommand(volumeAddCmd)
	volumeCmd.AddCommand(volumeRemoveCmd)

	volumeAddCmd.Flags().BoolVarP(&volumeRecursive, "recursive", "r", false, "Allow recursive volume")

	volumeCmd.Run = volumeList
	volumeListCmd.Run = volumeList
	volumeAddCmd.Run = volumeAdd
	volumeRemoveCmd.Run = volumeRemove
}

func volumeList(cmd *cobra.Command, args []string) {
	defer db.RecoverFunc()

	d, err := db.NewDB(appPath)
	if err != nil {
		log.Fatal(err)
	}

	volumes := d.List("volume")

	if len(volumes) > 0 {
		fmt.Println("Recursive Path")
		fmt.Println("--------- ----")
	}

	for _, v := range volumes {
		value := "false"
		if v.Value == "r" {
			value = "true"
		}

		fmt.Printf("%-9s %s\n", value, v.Key)
	}
}

func volumeAdd(cmd *cobra.Command, args []string) {
	defer db.RecoverFunc()

	d, err := db.NewDB(appPath)
	if err != nil {
		log.Fatal(err)
	}

	v := ""
	if volumeRecursive {
		v = "r"
	}

	d.Add("volume", args[0], []byte(v))
}

func volumeRemove(cmd *cobra.Command, args []string) {
	defer db.RecoverFunc()

	d, err := db.NewDB(appPath)
	if err != nil {
		log.Fatal(err)
	}

	d.Remove("volume", args[0])
}
