package migrations

import (
	"github.com/hanifmaliki/golang-boilerplate/pkg/model"

	gormigrate "github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var M20240428123801_companies = gormigrate.Migration{
	ID: "M20240428123801_companies",
	Migrate: func(tx *gorm.DB) error {
		type User struct {
			model.Base

			UserID  uint // Foreign key for User struct
			Street  string
			City    string
			Country string
		}
		return tx.Migrator().CreateTable(&User{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("users")
	},
}
