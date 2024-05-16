package repository

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	InTransaction(ctx context.Context, fn func(r Repository) error) error
	UserRepo() UserRepository
	CompanyRepo() CompanyRepository
	AddressRepo() AddressRepository
	CreditCardRepo() CreditCardRepository
	RoleRepo() RoleRepository
	UserRoleRepo() UserRoleRepository
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) InTransaction(ctx context.Context, fn func(r Repository) error) error {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(NewRepository(tx))
	})

	return err
}

func (r *repository) UserRepo() UserRepository {
	return NewUserRepository(r.db)
}

func (r *repository) CompanyRepo() CompanyRepository {
	return NewCompanyRepository(r.db)
}

func (r *repository) AddressRepo() AddressRepository {
	return NewAddressRepository(r.db)
}

func (r *repository) CreditCardRepo() CreditCardRepository {
	return NewCreditCardRepository(r.db)
}

func (r *repository) RoleRepo() RoleRepository {
	return NewRoleRepository(r.db)
}

func (r *repository) UserRoleRepo() UserRoleRepository {
	return NewUserRoleRepository(r.db)
}
