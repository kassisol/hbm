package system

import (
	"context"
	"fmt"

	"github.com/docker/docker/client"
	"github.com/juliengk/go-utils"
	collectionobj "github.com/kassisol/hbm/object/collection"
	configobj "github.com/kassisol/hbm/object/config"
	groupobj "github.com/kassisol/hbm/object/group"
	policyobj "github.com/kassisol/hbm/object/policy"
	resourceobj "github.com/kassisol/hbm/object/resource"
	userobj "github.com/kassisol/hbm/object/user"
	"github.com/kassisol/hbm/pkg/adf"
	"github.com/kassisol/hbm/version"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewInfoCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info",
		Short: "Display information about HBM",
		Long:  infoDescription,
		Args:  cobra.NoArgs,
		Run:   runInfo,
	}

	return cmd
}

func runInfo(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	cfg, err := configobj.New("sqlite", adf.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer cfg.End()

	p, err := policyobj.New("sqlite", adf.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer p.End()

	g, err := groupobj.New("sqlite", adf.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer g.End()

	u, err := userobj.New("sqlite", adf.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer u.End()

	c, err := collectionobj.New("sqlite", adf.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer c.End()

	r, err := resourceobj.New("sqlite", adf.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer r.End()

	authz, _ := cfg.Get("authorization")
	daae, _ := cfg.Get("default-allow-action-error")

	fmt.Println("Features Enabled:")
	fmt.Println(" Authorization:", authz)
	fmt.Println(" Default Allow Action On Error:", daae)
	fmt.Println("Policies:", p.Count())
	fmt.Println("Groups:", g.Count())
	fmt.Println(" Users:", u.Count())
	fmt.Println("Collections:", c.Count())
	fmt.Println(" Resources:", r.Count("all"))
	fmt.Println("  Actions:", r.Count("action"))
	fmt.Println("  Configs:", r.Count("config"))
	fmt.Println("  Capabilities:", r.Count("cap"))
	fmt.Println("  Devices:", r.Count("device"))
	fmt.Println("  DNS Servers:", r.Count("dns"))
	fmt.Println("  Images:", r.Count("image"))
	fmt.Println("  Ports:", r.Count("port"))
	fmt.Println("  Registries:", r.Count("registry"))
	fmt.Println("  Volumes:", r.Count("volume"))

	fmt.Println("Server Version:", version.Version)
	fmt.Println("Storage Driver: sqlite")
	fmt.Println("Logging Driver: standard")
	fmt.Println("Harbormaster Root Dir:", adf.AppPath)
	fmt.Println("Docker AuthZ Plugin Enabled:", pluginEnabled())
}

func pluginEnabled() bool {
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
