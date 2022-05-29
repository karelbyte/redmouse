package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type InvoiceProduct struct {
	ID                 uuid.UUID `gorm:"primaryKey;type:varchar(36)" json:"id"`
	InvoiceID          uuid.UUID
	Invoice            Invoice
	ProductVariationID uuid.UUID
	ProductVariation   ProductVariation
	Quantity           int32          `gorm:"type:varchar(10)" json:"quantity" binding:"required"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"-"`
}

func (invoiceProduct *InvoiceProduct) BeforeCreate(tx *gorm.DB) (err error) {
	invoiceProduct.ID = uuid.New()
	return
}
