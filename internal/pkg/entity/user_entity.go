package entity

import "github.com/hanifmaliki/golang-boilerplate/pkg/model"

type User struct {
	model.Base

	Name        string  `json:"name"`
	Email       string  `json:"email"`
	PhoneNumber *string `json:"phone_number"` // Optional
	CompanyID   uint    `json:"company_id"`

	Company     *Company      `json:"company" gorm:"->"`      // Belongs To relationship (one user belongs to one company)
	Address     *Address      `json:"address" gorm:"->"`      // Has One relationship (one user has one address)
	CreditCards []*CreditCard `json:"credit_cards" gorm:"->"` // Has Many relationship (one user has many credit cards)
	UserRoles   []*UserRole   `json:"user_roles" gorm:"->"`   // Many to Many relationship (one user can have many roles, and one role can have many users)
}
