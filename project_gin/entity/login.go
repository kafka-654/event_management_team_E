package entity

import (
	"time"
)

type Login struct {
	ID        uint `gorm:"primaryKey"`
	Username  string
	Password  string
	CreatedAt time.Time
}
