package cmd

import (
	"log"
	"os"

	"github.com/boltdb/bolt"
	"github.com/kassisol/hbm/pkg/db"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize config",
	Long:  "Initialize config",
}

func init() {
	RootCmd.AddCommand(initCmd)

	initCmd.Run = initConfig
}

func initConfig(cmd *cobra.Command, args []string) {
	if _, err := os.Stat(appPath); os.IsNotExist(err) {
		err := os.Mkdir(appPath, 0700)
		if err != nil {
			log.Fatal(err)
		}
	}

	d, err := db.NewDB(appPath)
	if err != nil {
		log.Fatal(err)
	}
	defer d.Conn.Close()

	buckets := []string{
                "action",
                "cap",
                "config",
                "device",
                "dns",
                "image",
                "port",
                "registry",
                "volume",
        }

	err = d.Conn.Update(func(tx *bolt.Tx) error {
		for _, bucket := range buckets {
			_, err := tx.CreateBucketIfNotExists([]byte(bucket))
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
