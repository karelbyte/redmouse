package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Size struct {
	ID        uuid.UUID `gorm:"primaryKey;type:varchar(36)" json:"id"`
	Description string  `gorm:"type:varchar(255)" json:"description" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (size *Size) BeforeCreate(tx *gorm.DB) (err error) {
	size.ID = uuid.New()
	return
}
