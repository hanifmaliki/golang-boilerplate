package repository

import (
	"context"

	"github.com/hanifmaliki/golang-boilerplate/internal/pkg/entity"
	pkg_model "github.com/hanifmaliki/golang-boilerplate/pkg/model"

	"gorm.io/gorm"
)

type CompanyRepository interface {
	BaseRepository[entity.Company]
	Find(ctx context.Context, request *entity.Company, query *pkg_model.Query) ([]*entity.Company, error)
}

type companyRepository struct {
	baseRepository[entity.Company]
}

func NewCompanyRepository(db *gorm.DB) CompanyRepository {
	return &companyRepository{
		baseRepository: baseRepository[entity.Company]{db: db},
	}
}

func (r *companyRepository) Find(ctx context.Context, request *entity.Company, query *pkg_model.Query) ([]*entity.Company, error) {
	var datas []*entity.Company
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
