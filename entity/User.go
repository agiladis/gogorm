package entity

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"prmaryKey"`
	Email     string `gorm:not null;unique;type:varchar(20)`
	Products  []Product
	CreatedAt time.Time
	UpdatedAt time.Time
}
