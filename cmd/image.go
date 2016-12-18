package cmd

import (
	"fmt"
	"log"

	"github.com/kassisol/hbm/pkg/db"
	"github.com/spf13/cobra"
)

var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "Manage whitelisted images",
	Long:  "Manage whitelisted images",
}

var imageListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List whitelisted images",
	Long:  "List whitelisted images",
}

var imageAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add image to the whitelist",
	Long:  "Add image to the whitelist",
}

var imageRemoveCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove image from the whitelist",
	Long:  "Remove image from the whitelist",
}

func init() {
	RootCmd.AddCommand(imageCmd)
	imageCmd.AddCommand(imageListCmd)
	imageCmd.AddCommand(imageAddCmd)
	imageCmd.AddCommand(imageRemoveCmd)

	imageCmd.Run = volumeList
	imageListCmd.Run = imageList
	imageAddCmd.Run = imageAdd
	imageRemoveCmd.Run = imageRemove
}

func imageList(cmd *cobra.Command, args []string) {
	defer db.RecoverFunc()

	d, err := db.NewDB(appPath)
	if err != nil {
		log.Fatal(err)
	}

	images := d.List("image")

	if len(images) > 0 {
		fmt.Println("Image")
		fmt.Println("-----")
	}

	for _, v := range images {
		fmt.Printf("%s\n", v.Key)
	}
}

func imageAdd(cmd *cobra.Command, args []string) {
	defer db.RecoverFunc()

	d, err := db.NewDB(appPath)
	if err != nil {
		log.Fatal(err)
	}

	d.Add("image", args[0], []byte(""))
}

func imageRemove(cmd *cobra.Command, args []string) {
	defer db.RecoverFunc()

	d, err := db.NewDB(appPath)
	if err != nil {
		log.Fatal(err)
	}

	d.Remove("image", args[0])
}
