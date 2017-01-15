package allow

import (
	"fmt"
	"log"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/allow/types"
	"github.com/kassisol/hbm/pkg/db"
	"github.com/kassisol/hbm/pkg/image"
	"github.com/kassisol/hbm/pkg/utils"
)

func AllowImageCreate(req authorization.Request, config *types.Config) *types.AllowResult {
	params := utils.GetURLParams(req.RequestURI)

	if !AllowImage(params["fromImage"][0], config) {
		return &types.AllowResult{Allow: false, Msg: fmt.Sprintf("Image %s is not allowed to be pulled", params["fromImage"][0])}
	}

	return &types.AllowResult{Allow: true}
}

func AllowImage(img string, config *types.Config) bool {
	defer db.RecoverFunc()

	d, err := db.NewDB(config.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer d.Conn.Close()

	i := image.NewImage(img)

	if i.Official {
		if d.KeyExists("config", "image_create_official") {
			return true
		}
	}

	if d.KeyExists("registry", i.Registry) {
		return true
	}

	if d.KeyExists("image", i.String()) {
		return true
	}

	return false
}
