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

	// Initialize the main query
	mainQuery := db

	// Handle preload/expansion of related entities
	for _, expand := range query.Expand {
		mainQuery = mainQuery.Preload(expand)
	}

	// Handle sorting
	if query.SortBy != "" {
		mainQuery = mainQuery.Order(query.SortBy)
	}

	// Execute the query
	err := mainQuery.Find(&datas).Error
	if err != nil {
		return nil, err
	}

	return datas, nil
}
