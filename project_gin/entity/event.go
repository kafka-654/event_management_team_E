package entity

import (
	"time"
)

type Event struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Date     time.Time
	Location string
}
