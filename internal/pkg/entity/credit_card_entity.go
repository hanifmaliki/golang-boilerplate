package entity

import "github.com/hanifmaliki/golang-boilerplate/pkg/model"

type CreditCard struct {
	model.Base

	UserID uint // Foreign key for User struct
	Number string
}
