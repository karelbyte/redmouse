package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type ProductVariation struct {
	ID               uuid.UUID `gorm:"primaryKey;type:varchar(36)" json:"id"`
	Code             string    `gorm:"type:varchar(255);not null" json:"code" binding:"required"`
	ShortDescription string    `gorm:"type:varchar(255);not null" json:"name" binding:"required"`
	Description      string    `gorm:"type:varchar(500)" json:"description"`
	ProductID        uuid.UUID `gorm:"not null"`
	Product          Product
	MeasureID        uuid.UUID `gorm:"not null"`
	Measure          Measure
	CategoryID       uuid.UUID `gorm:"not null"`
	Category         Category
	SizeID           uuid.UUID `gorm:"not null"`
	Size             Size
	ColorID          uuid.UUID `gorm:"not null"`
	Color            Color
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
}

func (productVariation *ProductVariation) BeforeCreate(tx *gorm.DB) (err error) {
	productVariation.ID = uuid.New()
	return
}
