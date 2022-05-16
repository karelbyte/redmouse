package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Inventory struct {
	ID                 uuid.UUID `gorm:"primaryKey;type:varchar(36)" json:"id"`
	WarehouseID        uuid.UUID
	Warehouse          Warehouse
	ProductVariationID uuid.UUID
	ProductVariation   ProductVariation
	Quantity           int64          `gorm:"type:int" json:"price"`
	MinInStock         int64          `gorm:"type:int" json:"min_in_stock"`
	Status             bool           `gorm:"type:tinyint" json:"status"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"-"`
}

func (inventory *Inventory) BeforeCreate(tx *gorm.DB) (err error) {
	inventory.ID = uuid.New()
	return
}
