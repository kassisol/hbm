package sqlite

import (
	"path"

	"github.com/kassisol/hbm/storage"
	"github.com/kassisol/hbm/storage/driver"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func init() {
	storage.RegisterDriver("sqlite", New)
}

type Config struct {
	DB *gorm.DB
}

func New(config string) (driver.Storager, error) {
	debug := false

	file := path.Join(config, "data.db")

	db, err := gorm.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}

	db.LogMode(debug)

	db.AutoMigrate(&AppConfig{}, &User{}, &Group{}, &Host{}, &Cluster{}, &Resource{}, &Collection{}, &Policy{})

	return &Config{DB: db}, nil
}

func (c *Config) End() {
	c.DB.Close()
}
