package sqlite

import (
	"time"
)

type Model struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"created_at"`
}

type AppConfig struct {
	Model
	Key string
}

type User struct {
	Model
	Name string `gorm:"unique;"`
}

type Group struct {
	Model
	Name  string `gorm:"unique;"`
	Users []User `gorm:"many2many:group_users;"`
}

type Resource struct {
	Model
	Name   string `gorm:"unique;"`
	Type   string
	Value  string
	Option string
}

type Collection struct {
	Model
	Name      string     `gorm:"unique;"`
	Resources []Resource `gorm:"many2many:collection_resources;"`
}

type Policy struct {
	Model
	Name         string `gorm:"unique;"`
	Group        Group
	GroupID      uint
	Collection   Collection
	CollectionID uint
}
