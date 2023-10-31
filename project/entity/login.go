package entity

import "github.com/jinzhu/gorm"

// Login represents a login entity.
type Login struct {
	gorm.Model
	Username string
	Password string
}
