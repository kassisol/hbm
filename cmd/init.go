package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/boltdb/bolt"
	"github.com/harbourmaster/hbm/pkg/db"
	"github.com/harbourmaster/hbm/pkg/utils"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize config",
	Long:  "Initialize config",
}

func init() {
	RootCmd.AddCommand(initCmd)

	initCmd.Run = initialconfig
}

func initialconfig(cmd *cobra.Command, args []string) {
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

	err = d.Conn.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("action"))
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists([]byte("cap"))
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists([]byte("config"))
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists([]byte("device"))
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists([]byte("dns"))
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists([]byte("image"))
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists([]byte("port"))
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists([]byte("registry"))
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists([]byte("volume"))
		if err != nil {
			return err
		}

		return nil
	})

	var dockerPluginPath = "/etc/docker/plugins"
	var dockerPluginFile = path.Join(dockerPluginPath, "hbm.spec")
	var pluginSpecContent = []byte("unix://run/docker/plugins/hbm.sock")

	if _, err := os.Stat(dockerPluginPath); os.IsNotExist(err) {
		err := os.Mkdir(dockerPluginPath, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	if !utils.FileExists(dockerPluginFile) {
		err := ioutil.WriteFile(dockerPluginFile, pluginSpecContent, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}
