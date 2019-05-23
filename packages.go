package main

import (
	_ "github.com/juliengk/go-log/driver/standard"
	_ "github.com/kassisol/hbm/storage/driver/sqlite"

	_ "github.com/kassisol/hbm/docker/resource/driver/action"
	_ "github.com/kassisol/hbm/docker/resource/driver/capability"
	_ "github.com/kassisol/hbm/docker/resource/driver/config"
	_ "github.com/kassisol/hbm/docker/resource/driver/device"
	_ "github.com/kassisol/hbm/docker/resource/driver/dns"
	_ "github.com/kassisol/hbm/docker/resource/driver/image"
	_ "github.com/kassisol/hbm/docker/resource/driver/logdriver"
	_ "github.com/kassisol/hbm/docker/resource/driver/logopt"
	_ "github.com/kassisol/hbm/docker/resource/driver/plugin"
	_ "github.com/kassisol/hbm/docker/resource/driver/port"
	_ "github.com/kassisol/hbm/docker/resource/driver/registry"
	_ "github.com/kassisol/hbm/docker/resource/driver/volume"
	_ "github.com/kassisol/hbm/docker/resource/driver/volumedriver"
)
