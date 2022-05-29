package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type ProductProvider struct {
	ID         uuid.UUID `gorm:"primaryKey;type:varchar(36)" json:"id"`
	ProductID  uuid.UUID `gorm:"not null"`;
	Product    Product
	ProviderID uuid.UUID `gorm:"not null"`;
	Provider   Provider
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

func (productProvider *ProductProvider) BeforeCreate(tx *gorm.DB) (err error) {
	productProvider.ID = uuid.New()
	return
}
