package repository

import (
	"context"

	"github.com/hanifmaliki/golang-boilerplate/internal/pkg/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	BaseRepository[entity.User]
	Test(ctx context.Context) (*entity.User, error)
}

type userRepository struct {
	baseRepository[entity.User]
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		baseRepository: baseRepository[entity.User]{db: db},
	}
}

func (r *userRepository) Test(ctx context.Context) (*entity.User, error) {
	var user entity.User
	return &user, nil
}
