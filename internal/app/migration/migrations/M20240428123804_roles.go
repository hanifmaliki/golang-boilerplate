package migrations

import (
	"github.com/hanifmaliki/golang-boilerplate/pkg/model"

	gormigrate "github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var M20240428123804_roles = gormigrate.Migration{
	ID: "M20240428123804_roles",
	Migrate: func(tx *gorm.DB) error {
		type User struct {
			model.Base

			Name string
		}
		return tx.Migrator().CreateTable(&User{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("users")
	},
}
