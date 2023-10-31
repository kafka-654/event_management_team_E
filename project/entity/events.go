package entity

import "github.com/jinzhu/gorm"

// Event represents an event entity.
type Event struct {
	gorm.Model
	Name   string
	TeamId uint
}
