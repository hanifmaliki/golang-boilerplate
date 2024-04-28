package entity

import "github.com/hanifmaliki/go-boilerplate/pkg/model"

type Address struct {
	model.BaseModel

	UserID  uint // Foreign key for User struct
	Street  string
	City    string
	Country string
}
