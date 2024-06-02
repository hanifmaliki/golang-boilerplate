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
