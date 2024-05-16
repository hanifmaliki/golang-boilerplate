package seeds

import (
	"github.com/hanifmaliki/golang-boilerplate/pkg/model"

	gormigrate "github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var S20240428123805_user_roles = gormigrate.Migration{
	ID: "M20240428123805_user_roles",
	Migrate: func(tx *gorm.DB) error {
		type User struct {
			model.Base

			UserID uint
			RoleID uint
		}
		return tx.Migrator().CreateTable(&User{})
	},
	Rollback: func(tx *gorm.DB) error {
		return nil
	},
}
