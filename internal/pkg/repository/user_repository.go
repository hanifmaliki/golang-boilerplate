package repository

import (
	"github.com/hanifmaliki/go-boilerplate/internal/pkg/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByToken(user *entity.User, token string) error
}

type userRepository struct {
	BaseRepository[entity.User]
}

func NewUserRepository(db *gorm.DB) UserRepository {
	r := &userRepository{}
	r.db = db
	return r
}

func (r *userRepository) FindByToken(user *entity.User, token string) error {
	return r.db.Where("token = ?", token).First(user).Error
}
