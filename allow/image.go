package allow

import (
	"fmt"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/juliengk/go-docker/image"
	"github.com/juliengk/go-log"
	"github.com/juliengk/go-log/driver"
	"github.com/juliengk/go-utils"
	"github.com/kassisol/hbm/allow/types"
	u "github.com/kassisol/hbm/pkg/utils"
	"github.com/kassisol/hbm/storage"
	"github.com/kassisol/hbm/version"
)

func AllowImageCreate(req authorization.Request, config *types.Config) *types.AllowResult {
	params := u.GetURLParams(req.RequestURI)

	if !AllowImage(params["fromImage"][0], config) {
		return &types.AllowResult{Allow: false, Msg: fmt.Sprintf("Image %s is not allowed to be pulled", params["fromImage"][0])}
	}

	return &types.AllowResult{Allow: true}
}

func AllowImage(img string, config *types.Config) bool {
	defer utils.RecoverFunc()

	l, _ := log.NewDriver("standard", nil)

	s, err := storage.NewDriver("sqlite", config.AppPath)
	if err != nil {
		l.WithFields(driver.Fields{
			"storagedriver": "sqlite",
			"logdriver":     "standard",
			"version":       version.Version,
		}).Fatal(err)
	}
	defer s.End()

	i := image.NewImage(img)

	if i.Official {
		if s.ValidatePolicy(config.Username, config.Hostname, "config", "image_create_official", "") {
			return true
		}
	}

	if s.ValidatePolicy(config.Username, config.Hostname, "registry", i.Registry, "") {
		return true
	}

	if s.ValidatePolicy(config.Username, config.Hostname, "image", i.String(), "") {
		return true
	}

	return false
}
