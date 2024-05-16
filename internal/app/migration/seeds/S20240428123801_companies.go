package seeds

import (
	"github.com/hanifmaliki/golang-boilerplate/pkg/model"

	gormigrate "github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var S20240428123802_addresses = gormigrate.Migration{
	ID: "M20240428123802_addresses",
	Migrate: func(tx *gorm.DB) error {
		type User struct {
			model.Base

			Name string
		}
		return tx.Migrator().CreateTable(&User{})
	},
	Rollback: func(tx *gorm.DB) error {
		return nil
	},
}
