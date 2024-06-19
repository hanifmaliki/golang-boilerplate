package model

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	// gorm.Model

	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index" swaggertype:"primitive,string"`

	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
	DeletedBy string `json:"deleted_dy"`
}

func (m *Base) GetID() uint {
	return m.ID
}

func (m *Base) SetCreatedBy(by string) {
	m.CreatedBy = by
}

func (m *Base) SetUpdatedBy(by string) {
	m.UpdatedBy = by
}
