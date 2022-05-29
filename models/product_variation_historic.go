package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type ProductVariationHistoric struct {
	ID                 uuid.UUID `gorm:"primaryKey;type:varchar(36)" json:"id"`
	ProductVariationID uuid.UUID `gorm:"not null"`;
	ProductVariation   ProductVariation
	DocumentID         uuid.UUID `gorm:"not null"`;
	DocumentType       string
	Quantity           int64          `gorm:"type:int;not null" json:"quantity" binding:"required"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"-"`
}

func (productVariationHistoric *ProductVariationHistoric) BeforeCreate(tx *gorm.DB) (err error) {
	productVariationHistoric.ID = uuid.New()
	return
}
