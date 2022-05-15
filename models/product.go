package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID               uuid.UUID `gorm:"primaryKey;type:varchar(36)" json:"id"`
	Description      string    `gorm:"type:varchar(500)" json:"description"`
	MeasureID        int
	Measure          Measure
	CategoryID       int
	Category         Category
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
}

func (product *Product) BeforeCreate(tx *gorm.DB) (err error) {
	product.ID = uuid.New()
	return
}
