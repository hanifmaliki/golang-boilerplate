package repository

import (
	"context"

	"github.com/hanifmaliki/golang-boilerplate/internal/pkg/entity"

	"gorm.io/gorm"
)

type CompanyRepository interface {
	BaseRepository[entity.Company]
	Test(ctx context.Context) (*entity.Company, error)
}

type companyRepository struct {
	baseRepository[entity.Company]
}

func NewCompanyRepository(db *gorm.DB) CompanyRepository {
	return &companyRepository{
		baseRepository: baseRepository[entity.Company]{db: db},
	}
}

func (r *companyRepository) Test(ctx context.Context) (*entity.Company, error) {
	var company entity.Company
	return &company, nil
}
