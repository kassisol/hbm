package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/boltdb/bolt"
	"github.com/kassisol/hbm/pkg/db"
	"github.com/kassisol/hbm/pkg/utils"
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
	var dockerPluginPath = "/etc/docker/plugins"
	var dockerPluginFile = path.Join(dockerPluginPath, "hbm.spec")
	var pluginSpecContent = []byte("unix://run/docker/plugins/hbm.sock")

	_, err := exec.LookPath("docker")
	if err != nil {
		fmt.Println("Docker does not seem to be installed. Please check your installation.")

		os.Exit(-1)
	}

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
