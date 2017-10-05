package system

import (
	"context"
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/docker/docker/client"
	"github.com/juliengk/go-utils"
	"github.com/kassisol/hbm/cli/command"
	"github.com/kassisol/hbm/plugin"
	"github.com/kassisol/hbm/storage"
	"github.com/kassisol/hbm/version"
	"github.com/spf13/cobra"
)

func NewInfoCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info",
		Short: "Display information about HBM",
		Long:  infoDescription,
		Run:   runInfo,
	}

	return cmd
}

func runInfo(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", command.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer s.End()

	fmt.Println("Features Enabled:")
	fmt.Println(" Authorization:", s.FindConfig("authorization"))
	fmt.Println("Policies:", s.CountPolicy())
	fmt.Println("Groups:", s.CountGroup())
	fmt.Println(" Users:", s.CountUser())
	fmt.Println("Collections:", s.CountCollection())
	fmt.Println(" Resources:", s.CountResource("all"))
	fmt.Println("  Actions:", s.CountResource("action"))
	fmt.Println("  Configs:", s.CountResource("config"))
	fmt.Println("  Capabilities:", s.CountResource("cap"))
	fmt.Println("  Devices:", s.CountResource("device"))
	fmt.Println("  DNS Servers:", s.CountResource("dns"))
	fmt.Println("  Images:", s.CountResource("image"))
	fmt.Println("  Ports:", s.CountResource("port"))
	fmt.Println("  Registries:", s.CountResource("registry"))
	fmt.Println("  Volumes:", s.CountResource("volume"))

	fmt.Println("Server Version:", version.Version)
	fmt.Println("Storage Driver: sqlite")
	fmt.Println("Logging Driver: standard")
	fmt.Println("Harbormaster Root Dir:", command.AppPath)
	fmt.Println("Docker AuthZ Plugin Enabled:", pluginEnabled())
	fmt.Println("Docker API Version Supported:", strings.Join(plugin.SupportedDockerAPIVersions, ", "))
}

func pluginEnabled() bool {
	apiVersion := plugin.SupportedDockerAPIVersions[0]
	apiVersion = apiVersion[1:len(apiVersion)]
	os.Setenv("DOCKER_API_VERSION", apiVersion)

	cli, err := client.NewEnvClient()
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

var infoDescription = `
Display information about HBM

`
