package repository

import (
	"context"

	"github.com/hanifmaliki/golang-boilerplate/pkg/model"

	"gorm.io/gorm"
)

type BaseRepository[T any] interface {
	Create(ctx context.Context, data *T) error
	Save(ctx context.Context, data *T) error
	Update(ctx context.Context, data *T, conds *T) error
	Delete(ctx context.Context, conds *T) error
	FindOne(ctx context.Context, conds *T, query *model.Query) (*T, error)
	Count(ctx context.Context, conds *T) (int64, error)
}

type baseRepository[T any] struct {
	db *gorm.DB
}

func (r *baseRepository[T]) Create(ctx context.Context, data *T) error {
	return r.db.WithContext(ctx).Create(data).Error
}

func (r *baseRepository[T]) Save(ctx context.Context, data *T) error {
	return r.db.WithContext(ctx).Save(data).Error
}

func (r *baseRepository[T]) Update(ctx context.Context, data *T, conds *T) error {
	return r.db.WithContext(ctx).Model(new(T)).Where(conds).Updates(data).Error
}

func (r *baseRepository[T]) Delete(ctx context.Context, conds *T) error {
	return r.db.WithContext(ctx).Where(conds).Delete(new(T)).Error
}

func (r *baseRepository[T]) FindOne(ctx context.Context, conds *T, query *model.Query) (*T, error) {
	var data *T
	db := r.db.WithContext(ctx)
	for _, expand := range query.Expand {
		db = db.Preload(expand)
	}
	if query.SortBy != "" {
		db = db.Order(query.SortBy)
	}
	err := db.Where(conds).First(data).Error
	return data, err
}

func (r *baseRepository[T]) Count(ctx context.Context, conds *T) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(new(T)).Where(conds).Count(&count).Error
	return count, err
}
