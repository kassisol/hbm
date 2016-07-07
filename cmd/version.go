package cmd

import (
	"fmt"

	"github.com/harbourmaster/hbm/version"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the HBM version information",
	Long:  "All software has versions. This is HBM's",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("HBM", version.VERSION, "-- HEAD")
	},
}
