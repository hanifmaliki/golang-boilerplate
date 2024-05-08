package repository

import "gorm.io/gorm"

type BaseRepository[T any] interface {
	Create(entity *T) error
	Update(entity *T) error
	Delete(entity *T) error
	CountById(id any) (int64, error)
	FindById(entity *T, id any) error
}

type baseRepository[T any] struct {
	db *gorm.DB
}

func (r *baseRepository[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

func (r *baseRepository[T]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

func (r *baseRepository[T]) Delete(entity *T) error {
	return r.db.Delete(entity).Error
}

func (r *baseRepository[T]) CountById(id any) (int64, error) {
	var total int64
	err := r.db.Model(new(T)).Where("id = ?", id).Count(&total).Error
	return total, err
}

func (r *baseRepository[T]) FindById(entity *T, id any) error {
	return r.db.Where("id = ?", id).Take(entity).Error
}
