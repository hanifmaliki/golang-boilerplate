package repository

import (
	"context"

	"github.com/hanifmaliki/golang-boilerplate/internal/pkg/entity"

	"gorm.io/gorm"
)

type CreditCardRepository interface {
	BaseRepository[entity.CreditCard]
	Test(ctx context.Context) (*entity.CreditCard, error)
}

type creditCardRepository struct {
	baseRepository[entity.CreditCard]
}

func NewCreditCardRepository(db *gorm.DB) CreditCardRepository {
	return &creditCardRepository{
		baseRepository: baseRepository[entity.CreditCard]{db: db},
	}
}

func (r *creditCardRepository) Test(ctx context.Context) (*entity.CreditCard, error) {
	var creditCard entity.CreditCard
	return &creditCard, nil
}
