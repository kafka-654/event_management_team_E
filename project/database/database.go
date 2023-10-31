package database

import (
	"project/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	var err error

	dsn := "root:Surya@123@tcp(127.0.0.1:3306)/sys"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	// AutoMigrate creates tables for your entities
	db.AutoMigrate(&entity.Department{})
	db.AutoMigrate(&entity.TeamMember{})
	db.AutoMigrate(&entity.Member{})
	db.AutoMigrate(&entity.Team{})
	db.AutoMigrate(&entity.Login{})
	db.AutoMigrate(&entity.Event{})

	// Set up database relationships here

	return db, nil
}

//	func GetDB() *gorm.DB {
//		return db
//	}

// func CloseDB() {
// 	if db != nil {
// 		db.Close()
// 	}
// }
