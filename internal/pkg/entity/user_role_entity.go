package entity

import "github.com/hanifmaliki/go-boilerplate/pkg/model"

type UserRole struct {
	model.BaseModel

	UserID uint
	RoleID uint
}
