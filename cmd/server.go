package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path"
	"syscall"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/plugin"
	"github.com/kassisol/hbm/pkg/utils"
	"github.com/spf13/cobra"
)

var serverConfig string

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts a Docker AuthZ server",
	Long:  "Starts a Docker AuthZ server",
}

func init() {
	RootCmd.AddCommand(serverCmd)

	serverCmd.Run = server
}

func serverInitConfig() {
	var dockerPluginPath = "/etc/docker/plugins"
	var dockerPluginFile = path.Join(dockerPluginPath, "hbm.spec")
	var pluginSpecContent = []byte("unix://run/docker/plugins/hbm.sock")

	_, err := exec.LookPath("docker")
	if err != nil {
		fmt.Println("Docker does not seem to be installed. Please check your installation.")

		os.Exit(-1)
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

	log.Print("Server has completed initialization")
}

func server(cmd *cobra.Command, args []string) {
	serverInitConfig()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, syscall.SIGTERM)

	go func() {
		p, err := plugin.NewPlugin(appPath)
		if err != nil {
			log.Fatal(err)
		}

		h := authorization.NewHandler(p)

		log.Print("Listening on socket file")
		log.Fatal(h.ServeUnix("root", "hbm"))
	}()

	s := <-ch
	log.Printf("Processing signal '%s'", s)
}
