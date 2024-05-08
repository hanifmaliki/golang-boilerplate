package seeds

import (
	"github.com/hanifmaliki/go-boilerplate/pkg/model"

	gormigrate "github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var S20240428123801_companies = gormigrate.Migration{
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
		return nil
	},
}
