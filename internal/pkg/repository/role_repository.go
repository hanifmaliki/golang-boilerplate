package repository

import (
	"context"

	"github.com/hanifmaliki/golang-boilerplate/internal/pkg/entity"
	pkg_model "github.com/hanifmaliki/golang-boilerplate/pkg/model"

	"gorm.io/gorm"
)

type RoleRepository interface {
	BaseRepository[entity.Role]
	Find(ctx context.Context, request *entity.Role, query *pkg_model.Query) ([]*entity.Role, error)
}

type roleRepository struct {
	baseRepository[entity.Role]
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{
		baseRepository: baseRepository[entity.Role]{db: db},
	}
}

func (r *roleRepository) Find(ctx context.Context, request *entity.Role, query *pkg_model.Query) ([]*entity.Role, error) {
	var datas []*entity.Role
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
