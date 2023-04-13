package entity

import (
	"time"
)

// name convention using singular, will be plural when migrate
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:not null;unique;type:varchar(20)`
	Products  []Product
	CreatedAt time.Time
	UpdatedAt time.Time
}
