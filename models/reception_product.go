package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type ReceptionProduct struct {
	ID                 uuid.UUID `gorm:"primaryKey;type:varchar(36)" json:"id"`
	ReceptionID        uuid.UUID
	Reception          Reception
	ProductVariationID uuid.UUID
	ProductVariation   ProductVariation
	Quantity           int32          `gorm:"type:varchar(10)" json:"quantity"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"-"`
}

func (receptionProduct *ReceptionProduct) BeforeCreate(tx *gorm.DB) (err error) {
	receptionProduct.ID = uuid.New()
	return
}
