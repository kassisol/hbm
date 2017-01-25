package cmd

import (
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
	"github.com/kassisol/hbm/plugin"
	"github.com/kassisol/hbm/version"
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

	l, err := log.NewDriver("standard", nil)
	if err != nil {
		l.Fatal(err)
	}

	if _, err = exec.LookPath("docker"); err != nil {
		l.Fatal("Docker does not seem to be installed. Please check your installation.")
	}

	if _, err := os.Stat(dockerPluginPath); os.IsNotExist(err) {
		if err := os.Mkdir(dockerPluginPath, 0755); err != nil {
			l.Fatal(err)
		}
	}

	if !filedir.FileExists(dockerPluginFile) {
		if err := ioutil.WriteFile(dockerPluginFile, pluginSpecContent, 0644); err != nil {
			l.Fatal(err)
		}
	}

	l.Info("Server has completed initialization")
}

func server(cmd *cobra.Command, args []string) {
	l, err := log.NewDriver("standard", nil)
	if err != nil {
		l.Fatal(err)
	}

	serverInitConfig()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, syscall.SIGTERM)

	go func() {
		p, err := plugin.NewPlugin(appPath)
		if err != nil {
			l.Fatal(err)
		}

		h := authorization.NewHandler(p)

		l.WithFields(driver.Fields{
			"storagedriver": "sqlite",
			"logdriver":     "standard",
			"version":       version.VERSION,
		}).Info("HBM server")

		l.Info("Listening on socket file")
		l.Fatal(h.ServeUnix("root", "hbm"))
	}()

	s := <-ch
	l.Info("Processing signal '%s'", s)
}
