package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Output struct {
	ID                        uuid.UUID `gorm:"primaryKey;type:varchar(36)" json:"id"`
	UserID                    uuid.UUID
	User                      User
	WarehouseID               uuid.UUID
	Warehouse                 Warehouse
	Code                      string                     `gorm:"type:varchar(10)" json:"code" binding:"required"`
	Note                      string                     `gorm:"type:varchar(255)" json:"note"`
	Status                    bool                       `gorm:"type:tinyint" json:"status"`
	ProductVariationHistorics []ProductVariationHistoric `gorm:"polymorphic:Document;"`
	CreatedAt                 time.Time                  `json:"created_at"`
	UpdatedAt                 time.Time                  `json:"updated_at"`
	DeletedAt                 gorm.DeletedAt             `gorm:"index" json:"-"`
}

func (output *Output) BeforeCreate(tx *gorm.DB) (err error) {
	output.ID = uuid.New()
	return
}
