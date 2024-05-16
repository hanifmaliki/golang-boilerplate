package seeds

import (
	"github.com/hanifmaliki/golang-boilerplate/pkg/model"

	gormigrate "github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var S20240428123803_credit_cards = gormigrate.Migration{
	ID: "M20240428123803_credit_cards",
	Migrate: func(tx *gorm.DB) error {
		type User struct {
			model.Base

			UserID uint // Foreign key for User struct
			Number string
		}
		return tx.Migrator().CreateTable(&User{})
	},
	Rollback: func(tx *gorm.DB) error {
		return nil
	},
}
