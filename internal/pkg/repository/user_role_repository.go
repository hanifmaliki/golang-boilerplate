package repository

import (
	"github.com/hanifmaliki/go-boilerplate/internal/pkg/entity"
	"gorm.io/gorm"
)

type UserRoleRepository interface {
	BaseRepository[entity.UserRole]
	Test() (*entity.UserRole, error)
}

type userRoleRepository struct {
	baseRepository[entity.UserRole]
}

func NewUserRoleRepository(db *gorm.DB) UserRoleRepository {
	return &userRoleRepository{
		baseRepository: baseRepository[entity.UserRole]{db: db},
	}
}

func (r *userRoleRepository) Test() (*entity.UserRole, error) {
	var userRole entity.UserRole
	return &userRole, nil
}
