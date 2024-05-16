package entity

import "github.com/hanifmaliki/golang-boilerplate/pkg/model"

type User struct {
	model.Base

	Name      string
	Email     string
	CompanyID uint

	Company     *Company      // Belongs To relationship (one user belongs to one company)
	Address     *Address      // Has One relationship (one user has one address)
	CreditCards []*CreditCard // Has Many relationship (one user has many credit cards)
	UserRoles   []*UserRole   // Many to Many relationship (one user can have many roles, and one role can have many users)
}
