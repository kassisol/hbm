package sqlite

import (
	"time"
)

// Model structure
type Model struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"created_at"`
}

// AppConfig structure
type AppConfig struct {
	Model
	Key string
}

// User structure
type User struct {
	Model
	Name string `gorm:"unique;"`
}

// Group structure
type Group struct {
	Model
	Name  string `gorm:"unique;"`
	Users []User `gorm:"many2many:group_users;"`
}

// Resource structure
type Resource struct {
	Model
	Name   string `gorm:"unique;"`
	Type   string
	Value  string
	Option string
}

// Collection structure
type Collection struct {
	Model
	Name      string     `gorm:"unique;"`
	Resources []Resource `gorm:"many2many:collection_resources;"`
}

// Policy structure
type Policy struct {
	Model
	Name         string `gorm:"unique;"`
	Group        Group
	GroupID      uint
	Collection   Collection
	CollectionID uint
}
