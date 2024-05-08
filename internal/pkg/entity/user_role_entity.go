package entity

import "github.com/hanifmaliki/go-boilerplate/pkg/model"

type UserRole struct {
	model.Base

	UserID uint
	RoleID uint
}
