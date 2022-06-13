package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Provider struct {
	ID          uuid.UUID `gorm:"primaryKey;type:varchar(36)" json:"id"`
	Names       string    `gorm:"type:varchar(255);not null" json:"names"`
	Description string    `gorm:"type:varchar(255)" json:"description"`
	Address     string    `gorm:"type:varchar(255)" json:"adress"`
	Rfc         string    `gorm:"type:varchar(10)" json:"rfc"`
	Email1      string    `gorm:"type:varchar(100);not null" json:"email1"`
	Email2      string    `gorm:"type:varchar(100)" json:"email2"`
	Phone1      string    `gorm:"type:varchar(20)" json:"phone1"`
	Phone2      string    `gorm:"type:varchar(20)" json:"phone2"`
	Products    []Product `json:"-"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (provider *Provider) BeforeCreate(tx *gorm.DB) (err error) {
	provider.ID = uuid.New()
	return
}
