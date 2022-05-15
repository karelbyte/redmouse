package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type ProductVariationPrice struct {
	ID                 uuid.UUID `gorm:"primaryKey;type:varchar(36)" json:"id"`
	Code               string    `gorm:"type:varchar(255)" json:"code"`
	ShortDescription   string    `gorm:"type:varchar(255)" json:"name"`
	Price              float64   `gorm:"type:decimal(12,2)" json:"price"`
	ProductVariationID uuid.UUID
	ProductVariation   ProductVariation
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"-"`
}

func (productVariationPrice *ProductVariationPrice) BeforeCreate(tx *gorm.DB) (err error) {
	productVariationPrice.ID = uuid.New()
	return
}
