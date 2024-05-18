package repository

import (
	"context"

	"github.com/hanifmaliki/golang-boilerplate/internal/pkg/entity"
	pkg_model "github.com/hanifmaliki/golang-boilerplate/pkg/model"

	"gorm.io/gorm"
)

type CreditCardRepository interface {
	BaseRepository[entity.CreditCard]
	Find(ctx context.Context, request *entity.CreditCard, query *pkg_model.Query) ([]*entity.CreditCard, error)
}

type creditCardRepository struct {
	baseRepository[entity.CreditCard]
}

func NewCreditCardRepository(db *gorm.DB) CreditCardRepository {
	return &creditCardRepository{
		baseRepository: baseRepository[entity.CreditCard]{db: db},
	}
}

func (r *creditCardRepository) Find(ctx context.Context, request *entity.CreditCard, query *pkg_model.Query) ([]*entity.CreditCard, error) {
	var datas []*entity.CreditCard
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
