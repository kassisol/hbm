package cmd

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/plugin"
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

func server(cmd *cobra.Command, args []string) {
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

	<-ch
}
