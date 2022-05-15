package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID          uuid.UUID      `gorm:"primaryKey;type:varchar(36)" json:"id"`
	Names       string         `gorm:"type:varchar(255)" json:"name"`
	Email       string         `gorm:"type:varchar(255)" json:"email"`
	Password    string         `gorm:"type:varchar(255)" json:"password"`
	Status      bool           `gorm:"type:tinyint" json:"status"`
	Description string         `gorm:"type:varchar(255)" json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	return
}
