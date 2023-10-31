package team

import (
	"project/entity"
	"time"

	// "github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

// CreatePerson creates a new person in the database.
func CreateTeam(db *gorm.DB, name string, departmentID uint) error {

	team := &entity.Team{Name: name, DepartmentID: departmentID}
	if err := db.Create(team).Error; err != nil {
		return err
	}
	return nil
}

// GetPerson retrieves a person by ID.
func GetTeam(db *gorm.DB, id uint) (*entity.Team, error) {

	var team entity.Team
	if err := db.First(&team, id).Error; err != nil {
		if err != nil {
			return nil, nil
		}
		return nil, err
	}
	return &team, nil
}

// UpdatePerson updates an existing person in the database.
func UpdateTeam(db *gorm.DB, id uint, name string) error {

	if err := db.Model(&entity.Team{}).Where("id = ?", id).Update("name", name).Error; err != nil {
		return err
	}
	return nil
}

// DeletePerson deletes a person by ID.
// func DeleteTeam(db *gorm.DB, id uint) error {

// 	if err := db.Where("id = ?", id).Delete(&entity.Team{}).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

func DeleteTeam(db *gorm.DB, id uint) error {

	if err := db.Model(&entity.Team{}).Where("id = ?", id).Update("deleted_at", time.Now().Format("2006-01-02 15:04:05")).Error; err != nil {
		return err
	}
	return nil
}
