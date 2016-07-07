package cmd

import (
	log "github.com/Sirupsen/logrus"
	"github.com/harbourmaster/hbm/pkg/db"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "hbm",
	Short: "HBM is an application to authorize and manage authorized docker command",
	Long:  "HBM is an application to authorize and manage authorized docker command",
}

func init() {
	RootCmd.PersistentFlags().StringVar(&dockerApiVersion, "api", "v1.23", "Use Docker API version (does not affect server subcommand)")
	RootCmd.PersistentFlags().StringVar(&appPath, "apppath", "/var/lib/hbm", "Database path")
	RootCmd.PersistentFlags().BoolVarP(&debug, "debug", "D", false, "Enable debug messages")
	RootCmd.PersistentFlags().StringVarP(&user, "user", "u", "", "Query db for user. Default is default user (i.e. the empty user).")
	RootCmd.PersistentPreRun = setLogLevel
}

func getBucket(name string) string {
	if user != "" {
		defer db.RecoverFunc()

		d, err := db.NewDB(appPath)
		if err != nil {
			log.Fatal(err)
		}

		name += "_" + user
		d.InitBucket(name)
	}
	return name
}

func setLogLevel(cmd *cobra.Command, args []string) {
	if debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}

var user string
var debug bool
var dockerApiVersion string
var appPath string
