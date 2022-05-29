package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type (
	APIUser struct {
		ID          uuid.UUID `json:"id"`
		Names       string    `json:"names"`
		Email       string    `json:"email"`
		Status      bool      `json:"status"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	User struct {
		ID          uuid.UUID      `gorm:"primaryKey;type:varchar(36)" json:"id"`
		Names       string         `gorm:"type:varchar(255);not null" json:"names" binding:"required"`
		Email       string         `gorm:"type:varchar(255);not null" json:"email" binding:"required"`
		Password    string         `gorm:"type:varchar(255);not null" json:"password" binding:"required"`
		Status      bool           `gorm:"type:tinyint" json:"status"`
		Description string         `gorm:"type:varchar(255)" json:"description"`
		CreatedAt   time.Time      `json:"created_at"`
		UpdatedAt   time.Time      `json:"updated_at"`
		DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	}
)

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	user.Password, _ = HashPassword(user.Password)
	return
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
