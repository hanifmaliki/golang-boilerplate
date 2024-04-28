package entity

import "github.com/hanifmaliki/go-boilerplate/pkg/model"

type Role struct {
	model.BaseModel

	Name  string
	Users []User `gorm:"many2many:user_roles;"` // Many to Many relationship (one role can have many users, and one user can have many roles)
}
