package entity

import "github.com/hanifmaliki/go-boilerplate/pkg/model"

type User struct {
	model.Base

	Name      string
	Email     string
	CompanyID uint

	Company     *Company      // Belongs To relationship (one user belongs to one company)
	Address     *Address      // Has One relationship (one user has one address)
	CreditCards []*CreditCard // Has Many relationship (one user has many credit cards)
	Roles       []*Role       `gorm:"many2many:user_roles;"` // Many to Many relationship (one user can have many roles, and one role can have many users)
}
