package repository

import "gorm.io/gorm"

type Repository interface {
	Transaction(fn func(r Repository) error) error
	User() UserRepository
	Company() CompanyRepository
	Address() AddressRepository
	CreditCard() CreditCardRepository
	Role() RoleRepository
	UserRole() UserRoleRepository
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Transaction(fn func(r Repository) error) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		return fn(NewRepository(tx))
	})

	return err
}

func (r *repository) User() UserRepository {
	return NewUserRepository(r.db)
}

func (r *repository) Company() CompanyRepository {
	return NewCompanyRepository(r.db)
}

func (r *repository) Address() AddressRepository {
	return NewAddressRepository(r.db)
}

func (r *repository) CreditCard() CreditCardRepository {
	return NewCreditCardRepository(r.db)
}

func (r *repository) Role() RoleRepository {
	return NewRoleRepository(r.db)
}

func (r *repository) UserRole() UserRoleRepository {
	return NewUserRoleRepository(r.db)
}
