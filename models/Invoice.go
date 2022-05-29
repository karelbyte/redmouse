package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Invoice struct {
	ID        uuid.UUID `gorm:"primaryKey;type:varchar(36)" json:"id"`
	UserID    uuid.UUID `gorm:"not null"`;
	User      User
	ClientID  uuid.UUID `gorm:"not null"`;
	Client    Client
	Code      string         `gorm:"type:varchar(10);not null" json:"code" binding:"required"`
	Note      string         `gorm:"type:varchar(255)" json:"note"`
	Status    bool           `gorm:"type:tinyint" json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (invoice *Invoice) BeforeCreate(tx *gorm.DB) (err error) {
	invoice.ID = uuid.New()
	return
}
