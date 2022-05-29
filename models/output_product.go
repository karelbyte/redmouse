package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type OuputProduct struct {
	ID                 uuid.UUID `gorm:"primaryKey;type:varchar(36)" json:"id"`
	OutputID           uuid.UUID `gorm:"not null"`;
	Output             Output
	ProductVariationID uuid.UUID `gorm:"not null"`;
	ProductVariation   ProductVariation
	Quantity           int32          `gorm:"type:int;not null" json:"quantity" binding:"required"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"-"`
}

func (ouputProduct *OuputProduct) BeforeCreate(tx *gorm.DB) (err error) {
	ouputProduct.ID = uuid.New()
	return
}
