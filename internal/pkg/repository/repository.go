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
	r := &repository{}
	r.db = db
	return r
}

func (r *repository) Transaction(fn func(r Repository) error) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		newRepository := NewRepository(tx)
		err := fn(newRepository)
		return err
	})

	return err
}

func (r *repository) User() UserRepository {
	return NewUserRepository(r.db)
}

type BaseRepository[T any] struct {
	db *gorm.DB
}

func (r *BaseRepository[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

func (r *BaseRepository[T]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

func (r *BaseRepository[T]) Delete(entity *T) error {
	return r.db.Delete(entity).Error
}

func (r *BaseRepository[T]) CountById(id any) (int64, error) {
	var total int64
	err := r.db.Model(new(T)).Where("id = ?", id).Count(&total).Error
	return total, err
}

func (r *BaseRepository[T]) FindById(entity *T, id any) error {
	return r.db.Where("id = ?", id).Take(entity).Error
}
