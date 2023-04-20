package entity

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// User represent the model for an user
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"not null;unique;type:varchar(20)"`
	Products  []Product
	CreatedAt time.Time
	UpdatedAt time.Time
}

// HOOK
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.Email == "" {
		return errors.New("email must be filled")
	}

	return nil
}
