package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type QuotationProduct struct {
	ID                 uuid.UUID `gorm:"primaryKey;type:varchar(36)" json:"id"`
	QuotationID        uuid.UUID
	Quotation          Quotation
	ProductVariationID uuid.UUID
	ProductVariation   ProductVariation
	Quantity           int32          `gorm:"type:varchar(10)" json:"quantity" binding:"required"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"-"`
}

func (quotationProduct *QuotationProduct) BeforeCreate(tx *gorm.DB) (err error) {
	quotationProduct.ID = uuid.New()
	return
}
