package database

import (
	"project_gin/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// InitDB initializes the MySQL database
func InitDB() error {
	dsn := "user:password@tcp(your-mysql-host:3306)/your-mysql-database?parseTime=true"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	db = d
	db.AutoMigrate(&entity.Login{})
	db.AutoMigrate(&entity.Person{})
	db.AutoMigrate(&entity.Event{})

	return nil
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return db
}

// CloseDB closes the database connection
func CloseDB() {
	sqlDB, _ := db.DB()
	sqlDB.Close()
}
