package repository

import (
	"github.com/hanifmaliki/go-boilerplate/internal/pkg/entity"
	"gorm.io/gorm"
)

type AddressRepository interface {
	BaseRepository[entity.Address]
	Test() (*entity.Address, error)
}

type addressRepository struct {
	baseRepository[entity.Address]
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &addressRepository{
		baseRepository: baseRepository[entity.Address]{db: db},
	}
}

func (r *addressRepository) Test() (*entity.Address, error) {
	var address entity.Address
	return &address, nil
}
