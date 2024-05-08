package entity

import "github.com/hanifmaliki/go-boilerplate/pkg/model"

type Address struct {
	model.Base

	UserID  uint // Foreign key for User struct
	Street  string
	City    string
	Country string
}
