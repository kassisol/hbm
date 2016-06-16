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
        Use:    "version",
        Short:  "Print the version number of Harbourmaster",
        Long:   "All software has versions. This is Harbourmaster's",
        Run:    func(cmd *cobra.Command, args []string) {
                fmt.Println("Harbourmaster", version.VERSION, "-- HEAD")
        },
}
