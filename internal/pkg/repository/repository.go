package repository

import "gorm.io/gorm"

type Repository interface {
	Transaction(fn func(r Repository) error) error
	User() UserRepository
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
