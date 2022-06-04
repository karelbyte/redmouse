package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID          uuid.UUID      `gorm:"primaryKey;type:varchar(36)" json:"id"`
	Description string         `gorm:"type:varchar(500);not null" json:"description" binding:"required"`
	Iva         int64          `gorm:"type:int;" json:"iva" binding:"required"`
	MeasureID   uuid.UUID      `gorm:"not null" json:"measure_id" binding:"required"`
	Measure     Measure        `gorm:"foreignKey:MeasureID;" json:"measure"`
	CategoryID  uuid.UUID      `gorm:"not null" json:"category_id" binding:"required"`
	Category    Category       `gorm:"foreignKey:CategoryID;" json:"category"`
	ProviderID  uuid.UUID      `gorm:"not null" json:"provider_id" binding:"required"`
	Provider    Provider       `gorm:"foreignKey:ProviderID;" json:"provider"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (product *Product) BeforeCreate(tx *gorm.DB) (err error) {
	product.ID = uuid.New()
	return
}
