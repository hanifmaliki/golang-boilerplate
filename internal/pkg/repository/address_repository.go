package repository

import (
	"context"

	"github.com/hanifmaliki/golang-boilerplate/internal/pkg/entity"

	"gorm.io/gorm"
)

type AddressRepository interface {
	BaseRepository[entity.Address]
	Test(ctx context.Context) (*entity.Address, error)
}

type addressRepository struct {
	baseRepository[entity.Address]
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &addressRepository{
		baseRepository: baseRepository[entity.Address]{db: db},
	}
}

func (r *addressRepository) Test(ctx context.Context) (*entity.Address, error) {
	var address entity.Address
	return &address, nil
}
