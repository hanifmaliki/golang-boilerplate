package migrations

import (
	"github.com/hanifmaliki/go-boilerplate/pkg/model"

	gormigrate "github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var M20240428123802_addresses = gormigrate.Migration{
	ID: "M20240428123802_addresses",
	Migrate: func(tx *gorm.DB) error {
		type User struct {
			model.BaseModel

			Name string
		}
		return tx.Migrator().CreateTable(&User{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("users")
	},
}
