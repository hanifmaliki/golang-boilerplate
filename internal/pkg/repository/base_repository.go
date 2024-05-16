package repository

import (
	"context"

	"github.com/hanifmaliki/golang-boilerplate/pkg/model"

	"gorm.io/gorm"
)

type BaseRepository[T any] interface {
	Find(ctx context.Context, conds *T, query model.Query) ([]*T, error)
	FindOne(ctx context.Context, conds *T, query model.Query) (*T, error)
	Create(ctx context.Context, data *T) (*T, error)
	Update(ctx context.Context, data *T) (*T, error)
	Delete(ctx context.Context, conds *T) error
	Count(ctx context.Context, conds *T) (int64, error)
}

type baseRepository[T any] struct {
	db *gorm.DB
}

func (r *baseRepository[T]) Find(ctx context.Context, conds *T, query model.Query) ([]*T, error) {
	var datas []*T
	db := r.db.WithContext(ctx)
	for _, expand := range query.Expand {
		db = db.Preload(expand)
	}
	if query.SortBy != "" {
		db = db.Order(query.SortBy)
	}
	err := db.Where(conds).Find(datas).Error
	return datas, err
}

func (r *baseRepository[T]) FindOne(ctx context.Context, conds *T, query model.Query) (*T, error) {
	var data *T
	db := r.db.WithContext(ctx)
	for _, expand := range query.Expand {
		db = db.Preload(expand)
	}
	err := db.Where(conds).First(data).Error
	return data, err
}

func (r *baseRepository[T]) Create(ctx context.Context, data *T) (*T, error) {
	err := r.db.WithContext(ctx).Create(data).Error
	return data, err
}

func (r *baseRepository[T]) Update(ctx context.Context, data *T) (*T, error) {
	err := r.db.WithContext(ctx).Save(data).Error
	return data, err
}

func (r *baseRepository[T]) Delete(ctx context.Context, conds *T) error {
	return r.db.WithContext(ctx).Where(conds).Delete(new(T)).Error
}

func (r *baseRepository[T]) Count(ctx context.Context, conds *T) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(new(T)).Where(conds).Count(&count).Error
	return count, err
}
