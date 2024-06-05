package repository

import (
	"context"
	"math"
	"strings"

	"github.com/hanifmaliki/golang-boilerplate/internal/pkg/entity"
	"github.com/hanifmaliki/golang-boilerplate/internal/pkg/model"
	pkg_model "github.com/hanifmaliki/golang-boilerplate/pkg/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	BaseRepository[entity.User]
	Find(ctx context.Context, request *model.GetUserRequest, query *pkg_model.Query) ([]*entity.User, *pkg_model.Pagination, error)
}

type userRepository struct {
	baseRepository[entity.User]
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		baseRepository: baseRepository[entity.User]{db: db},
	}
}

func (r *userRepository) Find(ctx context.Context, request *model.GetUserRequest, query *pkg_model.Query) ([]*entity.User, *pkg_model.Pagination, error) {
	var datas []*entity.User
	db := r.db.WithContext(ctx)

	// Apply filters
	if request.Search != "" {
		search := "%" + strings.ToLower(request.Search) + "%"
		db = db.Where("LOWER(name) LIKE ? OR LOWER(email) LIKE ?", search, search)
	}
	if len(request.CompanyID) > 0 {
		db = db.Where("company_id = ?", request.CompanyID)
	}
	if len(request.RoleID) > 0 {
		db = db.Where("id IN (?)", r.db.Table("user_roles").Select("user_id").
			Where("role_id IN ?", request.RoleID))
	}

	// New Session Methods for handle contaminated SQL queries
	db = db.Session(&gorm.Session{})

	// Pagination setup
	var pagination *pkg_model.Pagination
	if query.Page > 0 && query.PageSize > 0 {
		var totalItems int64
		if err := db.Model(&entity.User{}).Count(&totalItems).Error; err != nil {
			return nil, nil, err
		}

		pagination = &pkg_model.Pagination{
			Page:       query.Page,
			PageSize:   query.PageSize,
			TotalItems: totalItems,
			TotalPages: int(math.Ceil(float64(totalItems) / float64(query.PageSize))),
		}

		db = db.Offset((query.Page - 1) * query.PageSize).Limit(query.PageSize)
	}

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
		return nil, nil, err
	}

	return datas, pagination, nil
}
