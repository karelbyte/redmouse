package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type ProductVariation struct {
	ID               uuid.UUID `gorm:"primaryKey;type:varchar(36)" json:"id"`
	Code             string    `gorm:"type:varchar(255)" json:"code" binding:"required"`
	ShortDescription string    `gorm:"type:varchar(255)" json:"name" binding:"required"`
	Description      string    `gorm:"type:varchar(500)" json:"description"`
	ProductID        uuid.UUID
	Product          Product
	MeasureID        uuid.UUID
	Measure          Measure
	CategoryID       uuid.UUID
	Category         Category
	SizeID           uuid.UUID
	Size             Size
	ColorID          uuid.UUID
	Color            Color
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
}

func (productVariation *ProductVariation) BeforeCreate(tx *gorm.DB) (err error) {
	productVariation.ID = uuid.New()
	return
}
