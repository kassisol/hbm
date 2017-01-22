package cmd

import (
	"fmt"

	"github.com/docker/engine-api/client"
	"github.com/juliengk/go-utils"
	"github.com/kassisol/hbm/storage"
	"github.com/kassisol/hbm/version"
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
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", appPath)
	if err != nil {
		utils.Exit(err)
	}
	defer s.End()

	fmt.Println("Policies:", s.CountPolicy())
	fmt.Println("Groups:", s.CountGroup())
	fmt.Println(" Users:", s.CountUser())
	fmt.Println("Clusters:", s.CountCluster())
	fmt.Println(" Hosts:", s.CountHost())
	fmt.Println("Collections:", s.CountCollection())
	fmt.Println(" Resources:", s.CountResource("all"))
	fmt.Println("  Actions:", s.CountResource("action"))
	fmt.Println("  Config:", s.CountResource("config"))
	fmt.Println("  Capabilities:", s.CountResource("cap"))
	fmt.Println("  Devices:", s.CountResource("device"))
	fmt.Println("  DNS Servers:", s.CountResource("dns"))
	fmt.Println("  Images:", s.CountResource("image"))
	fmt.Println("  Ports:", s.CountResource("port"))
	fmt.Println("  Registries:", s.CountResource("registry"))
	fmt.Println("  Volumes:", s.CountResource("volume"))

	fmt.Println("Server Version:", version.VERSION)
	fmt.Println("Storage Driver: sqlite")
	fmt.Println("Logging Driver: standard")
	fmt.Println("Server Status: running")
	fmt.Println("Docker AuthZ Plugin Enabled:", PluginEnabled())
	fmt.Println("Harbormaster Root Dir:", appPath)
	fmt.Println("Debug mode (server): false")
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
