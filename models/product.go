package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID               uuid.UUID `gorm:"primaryKey;type:varchar(36)" json:"id"`
	Description      string    `gorm:"type:varchar(500)" json:"description"`
	Iva              int64 `gorm:"type:int" son:"iva"`
	MeasureID        uuid.UUID
	Measure          Measure
	CategoryID       uuid.UUID
	Category         Category
	ProviderID       uuid.UUID
	Provider         Provider
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
}

func (product *Product) BeforeCreate(tx *gorm.DB) (err error) {
	product.ID = uuid.New()
	return
}
