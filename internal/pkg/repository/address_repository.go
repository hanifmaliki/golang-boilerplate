package repository

import (
	"github.com/hanifmaliki/go-boilerplate/internal/pkg/entity"

	"gorm.io/gorm"
)

type AddressRepository interface {
	FindByToken(user *entity.User, token string) error
}

type addressRepository struct {
	BaseRepository[entity.User]
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	r := &addressRepository{}
	r.db = db
	return r
}

func (r *addressRepository) FindByToken(user *entity.User, token string) error {
	return r.db.Where("token = ?", token).First(user).Error
}

func (r *addressRepository) CreateWithAddressCreditCard(user *entity.User, token string) error {
	return nil
}
