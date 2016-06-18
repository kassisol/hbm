package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:	"hbm",
	Short:	"HBM is an application to authorize and manage authorized docker command",
	Long:	"HBM is an application to authorize and manage authorized docker command",
}

var appPath = "/var/lib/hbm"
