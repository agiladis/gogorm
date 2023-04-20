package entity

import (
	"time"
)

// name convention using singular, will be plural when migrate
type Product struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null;type:varchar(200)"`
	Brand     string `gorm:"not null;type:varchar(200)"`
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
