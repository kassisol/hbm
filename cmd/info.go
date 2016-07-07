package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/docker/engine-api/client"
	"github.com/harbourmaster/hbm/pkg/db"
	"github.com/harbourmaster/hbm/version"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Display information about HBM",
	Long:  "Display information about HBM",
}

func init() {
	RootCmd.AddCommand(infoCmd)

	infoCmd.Run = infoView
}

func infoView(cmd *cobra.Command, args []string) {
	defer db.RecoverFunc()

	d, err := db.NewDB(appPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Whitelists:")
	fmt.Println(" Actions:", strconv.Itoa(d.Count("action")))
	fmt.Println(" Capabilities:", strconv.Itoa(d.Count("cap")))
	fmt.Println(" Configs:", strconv.Itoa(d.Count("config")))
	fmt.Println(" Devices:", strconv.Itoa(d.Count("device")))
	fmt.Println(" DNS servers:", strconv.Itoa(d.Count("dns")))
	fmt.Println(" Images:", strconv.Itoa(d.Count("image")))
	fmt.Println(" Ports:", strconv.Itoa(d.Count("port")))
	fmt.Println(" Registries:", strconv.Itoa(d.Count("registry")))
	fmt.Println(" Volumes:", strconv.Itoa(d.Count("volume")))

	d.Conn.Close()

	fmt.Println("Server Version:", version.VERSION)
	fmt.Println("Harbourmaster Root Dir:", appPath)
	fmt.Println("Server Status: running")
	fmt.Println("Docker AuthZ Plugin Enabled:", PluginEnabled())
}

func PluginEnabled() bool {
	defaultHeaders := map[string]string{"User-Agent": "engine-api-cli-1.0"}
	cli, err := client.NewClient("unix:///var/run/docker.sock", "v1.23", nil, defaultHeaders)
	if err != nil {
		return false
	}

	info, err := cli.Info(context.Background())
	if err != nil {
		return false
	}

	for _, p := range info.Plugins.Authorization {
		if p == "hbm" {
			return true
		}
	}

	return false
}
