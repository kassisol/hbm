package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:	"hbm",
	Short:	"HBM is a command line to restrict docker use",
	Long:	"HBM is a command line to restrict docker use",
}

var appPath = "/var/lib/hbm"
