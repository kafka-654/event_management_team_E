package entity

import (
	"time"
)

type Person struct {
	ID        uint `gorm:"primaryKey"`
	FirstName string
	LastName  string
	CreatedAt time.Time
}
