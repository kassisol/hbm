package api

import (
	"fmt"
	log "github.com/Sirupsen/logrus"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/harbourmaster/hbm/api/types"
	"github.com/harbourmaster/hbm/pkg/db"
	"github.com/harbourmaster/hbm/pkg/image"
	"github.com/harbourmaster/hbm/pkg/utils"
)

func AllowImageCreate(req authorization.Request, config *types.Config) *types.AllowResult {
	params := utils.GetURLParams(req.RequestURI)

	if req.User != "" {
		for _, val := range []string{"config", "image", "registry"} {
			initUserBucket(val, req.User, config.AppPath)
		}
	}

	if !AllowImage(req.User, params["fromImage"][0], config) {
		userstr := " for anonymous user"
		if req.User != "" {
			userstr = " for user " + req.User
		}
		return &types.AllowResult{Allow: false, Msg: fmt.Sprintf("Image %s is not allowed to be pulled"+userstr, params["fromImage"][0])}
	}

	return &types.AllowResult{Allow: true}
}

func AllowImage(user string, img string, config *types.Config) bool {
	configBuckets := []string{"config"}
	imageBuckets := []string{"image"}
	registryBuckets := []string{"registry"}

	if user != "" {
		configBuckets = append(configBuckets, configBuckets[0]+"_"+user)
		imageBuckets = append(imageBuckets, imageBuckets[0]+"_"+user)
		registryBuckets = append(registryBuckets, registryBuckets[0]+"_"+user)
	}
	defer db.RecoverFunc()

	d, err := db.NewDB(config.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer d.Conn.Close()

	i := image.NewImage(img)

	if i.Official {
		if d.KeyExistsInBuckets(configBuckets, "image_create_official") {
			log.Debug("Passing on image_create_official")
			return true
		}
	}

	if d.KeyExistsInBuckets(registryBuckets, i.Registry) {
		log.Debug("Passing on allowed registry")
		return true
	}

	if d.KeyExistsInBuckets(imageBuckets, i.String()) {
		log.Debug("Passing on allowed image")
		return true
	}

	return false
}
