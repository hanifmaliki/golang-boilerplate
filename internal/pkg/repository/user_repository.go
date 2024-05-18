package repository

import (
	"context"
	"math"

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
	baseQuery := db

	if search := request.Search; search != "" {
		search = "%" + search + "%"
		db = db.Where("name LIKE ? OR email LIKE ?", search, search)
	}
	if len(request.CompanyID) > 0 {
		db = db.Where("company_id = ?", request.CompanyID)
	}
	if len(request.RoleID) > 0 {
		db = db.Where("id IN (?)", baseQuery.Table("user_roles").Select("user_id").
			Where("role_id IN ?", request.RoleID))
	}

	// New Session Methods for handle contaminated SQL queries
	db = db.Session(&gorm.Session{})

	// Init main query
	result := db

	// Pagination
	pagination := &pkg_model.Pagination{}
	if query.Page > 0 && query.PageSize > 0 {
		pagination.Page = query.Page
		pagination.PageSize = query.PageSize

		var totalItems int64
		resultCount := db.Model(&entity.User{}).Count(&totalItems)
		if resultCount.Error != nil {
			return nil, nil, resultCount.Error
		}
		pagination.TotalItems = totalItems
		pagination.TotalPages = int(math.Ceil(float64(totalItems) / float64(query.PageSize)))

		result = result.Offset((query.Page - 1) * query.PageSize).Limit(query.PageSize)
	} else {
		pagination = nil
	}

	for _, expand := range query.Expand {
		result = result.Preload(expand)
	}

	if query.SortBy != "" {
		result = result.Order(query.SortBy)
	}

	err := result.Find(datas).Error
	return datas, pagination, err
}
