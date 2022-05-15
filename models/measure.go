package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Measure struct {
	ID        uuid.UUID `gorm:"primaryKey;type:varchar(36)" json:"id"`
	Description string  `gorm:"type:varchar(255)" json:"description"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (measure *Measure) BeforeCreate(tx *gorm.DB) (err error) {
	measure.ID = uuid.New()
	return
}
