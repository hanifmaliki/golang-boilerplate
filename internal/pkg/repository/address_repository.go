package repository

import (
	"context"

	"github.com/hanifmaliki/golang-boilerplate/internal/pkg/entity"
	pkg_model "github.com/hanifmaliki/golang-boilerplate/pkg/model"

	"gorm.io/gorm"
)

type AddressRepository interface {
	BaseRepository[entity.Address]
	Find(ctx context.Context, request *entity.Address, query *pkg_model.Query) ([]*entity.Address, error)
}

type addressRepository struct {
	baseRepository[entity.Address]
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &addressRepository{
		baseRepository: baseRepository[entity.Address]{db: db},
	}
}

func (r *addressRepository) Find(ctx context.Context, request *entity.Address, query *pkg_model.Query) ([]*entity.Address, error) {
	var datas []*entity.Address
	db := r.db.WithContext(ctx)

	// Handle preload/expansion of related entities
	for _, expand := range query.Expand {
		db = db.Preload(expand)
	}

	// Handle sorting
	if query.SortBy != "" {
		db = db.Order(query.SortBy)
	}

	// Execute the query
	if err := db.Find(&datas).Error; err != nil {
		return nil, err
	}

	return datas, nil
}
