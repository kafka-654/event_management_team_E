package entity

import "github.com/jinzhu/gorm"

type Member struct {
	gorm.Model
	Name string
}
