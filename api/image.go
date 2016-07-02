package api

import (
	"fmt"
	"log"
	"net/url"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/harbourmaster/hbm/api/types"
	"github.com/harbourmaster/hbm/pkg/db"
	"github.com/harbourmaster/hbm/pkg/image"
)

func AllowImageCreate(req authorization.Request, config *types.Config) *types.AllowResult {
	url, err := url.ParseRequestURI(req.RequestURI)
	if err != nil {
		log.Fatal(err)
	}

	params := url.Query()

	if ! AllowImage(params["fromImage"][0], config) {
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
