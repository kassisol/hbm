package system

import (
	"fmt"

	"github.com/kassisol/hbm/version"
	"github.com/spf13/cobra"
)

func NewVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Show the HBM version information",
		Long:  versionDescription,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("HBM", version.VERSION, "-- HEAD")
		},
	}

	return cmd
}

var versionDescription = `
All software has versions. This is HBM's

`
