package repository

import (
	"context"

	"github.com/hanifmaliki/go-boilerplate/internal/pkg/entity"

	"gorm.io/gorm"
)

type RoleRepository interface {
	BaseRepository[entity.Role]
	Test(ctx context.Context) (*entity.Role, error)
}

type roleRepository struct {
	baseRepository[entity.Role]
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{
		baseRepository: baseRepository[entity.Role]{db: db},
	}
}

func (r *roleRepository) Test(ctx context.Context) (*entity.Role, error) {
	var role entity.Role
	return &role, nil
}
