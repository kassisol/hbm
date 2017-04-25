package server

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/signal"
	"path"
	"syscall"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/juliengk/go-log"
	"github.com/juliengk/go-log/driver"
	"github.com/juliengk/go-utils/filedir"
	"github.com/kassisol/hbm/cli/command"
	"github.com/kassisol/hbm/plugin"
	"github.com/kassisol/hbm/version"
	"github.com/spf13/cobra"
)

var serverConfig string

func NewServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Launch the HBM server",
		Long:  serverDescription,
		Run:   runStart,
	}

	return cmd
}

func serverInitConfig() {
	dockerPluginPath := "/etc/docker/plugins"
	dockerPluginFile := path.Join(dockerPluginPath, "hbm.spec")
	pluginSpecContent := []byte("unix://run/docker/plugins/hbm.sock")

	l, err := log.NewDriver("standard", nil)
	if err != nil {
		fmt.Println(err)

		os.Exit(-1)
	}

	_, err = exec.LookPath("docker")
	if err != nil {
		fmt.Println("Docker does not seem to be installed. Please check your installation.")

		os.Exit(-1)
	}

	if err := filedir.CreateDirIfNotExist(dockerPluginPath, false, 0755); err != nil {
		l.Fatal(err)
	}

	if !filedir.FileExists(dockerPluginFile) {
		err := ioutil.WriteFile(dockerPluginFile, pluginSpecContent, 0644)
		if err != nil {
			l.Fatal(err)
		}
	}

	l.Info("Server has completed initialization")
}

func runStart(cmd *cobra.Command, args []string) {
	l, err := log.NewDriver("standard", nil)
	if err != nil {
		fmt.Println(err)

		os.Exit(-1)
	}

	serverInitConfig()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, syscall.SIGTERM)

	go func() {
		p, err := plugin.NewPlugin(command.AppPath)
		if err != nil {
			l.Fatal(err)
		}

		h := authorization.NewHandler(p)

		l.WithFields(driver.Fields{
			"storagedriver": "sqlite",
			"logdriver":     "standard",
			"version":       version.Version,
		}).Info("HBM server")

		l.Info("Listening on socket file")
		l.Fatal(h.ServeUnix("root", "hbm"))
	}()

	s := <-ch
	l.Infof("Processing signal '%s'", s)
}

var serverDescription = `
Launch the HBM server

`
