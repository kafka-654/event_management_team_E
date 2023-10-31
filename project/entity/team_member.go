package entity

import "github.com/jinzhu/gorm"

type TeamMember struct {
	gorm.Model
	TeamID   uint
	MemberID uint
}
