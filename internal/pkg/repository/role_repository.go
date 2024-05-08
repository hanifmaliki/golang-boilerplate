package repository

import (
	"github.com/hanifmaliki/go-boilerplate/internal/pkg/entity"
	"gorm.io/gorm"
)

type RoleRepository interface {
	BaseRepository[entity.Role]
	Test() (*entity.Role, error)
}

type roleRepository struct {
	baseRepository[entity.Role]
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{
		baseRepository: baseRepository[entity.Role]{db: db},
	}
}

func (r *roleRepository) Test() (*entity.Role, error) {
	var role entity.Role
	return &role, nil
}
