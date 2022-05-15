package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Color struct {
	ID          uuid.UUID      `gorm:"primaryKey;type:varchar(36)" json:"id"`
	Description string         `gorm:"type:varchar(255)" json:"description"`
	Tone        string         `gorm:"type:varchar(10)" json:"tone"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (color *Color) BeforeCreate(tx *gorm.DB) (err error) {
	color.ID = uuid.New()
	return
}
