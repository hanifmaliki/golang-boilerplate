package repository

import (
	"context"

	"github.com/hanifmaliki/golang-boilerplate/internal/pkg/entity"
	pkg_model "github.com/hanifmaliki/golang-boilerplate/pkg/model"

	"gorm.io/gorm"
)

type UserRoleRepository interface {
	BaseRepository[entity.UserRole]
	Find(ctx context.Context, request *entity.UserRole, query *pkg_model.Query) ([]*entity.UserRole, error)
}

type userRoleRepository struct {
	baseRepository[entity.UserRole]
}

func NewUserRoleRepository(db *gorm.DB) UserRoleRepository {
	return &userRoleRepository{
		baseRepository: baseRepository[entity.UserRole]{db: db},
	}
}

func (r *userRoleRepository) Find(ctx context.Context, request *entity.UserRole, query *pkg_model.Query) ([]*entity.UserRole, error) {
	var datas []*entity.UserRole
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
