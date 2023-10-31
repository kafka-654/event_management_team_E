package login

import (
	"project/entity"
	"time"

	// "github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

// CreateLogin creates a new login record in the database.
func CreateLogin(db *gorm.DB, username string, password string) error {
	// db := database.GetDB() // Implement GetDB method in database package

	login := &entity.Login{Username: username, Password: password}
	member := &entity.Member{Name: username}
	// db.Create(&entity.Login{Username: username, Password: password})

	if err := db.Create(login).Error; err != nil {
		return err
	}
	if err := db.Create(member).Error; err != nil {
		return err
	}
	return nil
}

// GetLogin retrieves a login by username and password.
func GetLogin(db *gorm.DB, username string, password string) (*entity.Login, error) {

	login := entity.Login{}
	if err := db.First(&login, "Username = ? AND Password = ?", username, password).Error; err != nil {
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	return &login, nil
}

// UpdateLogin updates an existing login record in the database.
func UpdateLogin(db *gorm.DB, username string, newPassword string) error {

	if err := db.Model(&entity.Login{}).Where("username = ?", username).Update("password", newPassword).Error; err != nil {
		return err
	}
	return nil
}

// DeleteLogin deletes a login record from the database.
// func DeleteLogin(db *gorm.DB, username string) error {

//		if err := db.Where("Username = ?", username).Delete(&entity.Login{}).Error; err != nil {
//			return err
//		}
//		if err := db.Where("Name = ?", username).Delete(&entity.Member{}).Error; err != nil {
//			return err
//		}
//		return nil
//	}
func DeleteLogin(db *gorm.DB, username string) error {

	if err := db.Model(&entity.Login{}).Where("Username = ?", username).Update("deleted_at", time.Now().Format("2006-01-02 15:04:05")).Error; err != nil {
		return err
	}
	if err := db.Model(&entity.Member{}).Where("Name = ?", username).Update("deleted_at", time.Now().Format("2006-01-02 15:04:05")).Error; err != nil {
		return err
	}
	return nil
}
