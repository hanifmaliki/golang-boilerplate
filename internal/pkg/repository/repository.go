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
	return nil
}

func (r *repository) User() UserRepository {
	return NewUserRepository(r.db)
}

type BaseRepository[T any] struct {
	db *gorm.DB
}

func (r *BaseRepository[T]) Create(db *gorm.DB, entity *T) error {
	return db.Create(entity).Error
}

func (r *BaseRepository[T]) Update(db *gorm.DB, entity *T) error {
	return db.Save(entity).Error
}

func (r *BaseRepository[T]) Delete(db *gorm.DB, entity *T) error {
	return db.Delete(entity).Error
}

func (r *BaseRepository[T]) CountById(db *gorm.DB, id any) (int64, error) {
	var total int64
	err := db.Model(new(T)).Where("id = ?", id).Count(&total).Error
	return total, err
}

func (r *BaseRepository[T]) FindById(db *gorm.DB, entity *T, id any) error {
	return db.Where("id = ?", id).Take(entity).Error
}
