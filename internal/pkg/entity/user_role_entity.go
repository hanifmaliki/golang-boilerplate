package entity

import "github.com/hanifmaliki/golang-boilerplate/pkg/model"

type UserRole struct {
	model.Base

	UserID uint
	RoleID uint

	Role *Role
}
