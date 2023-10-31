package events

import (
	"project/entity"
	"time"

	"gorm.io/gorm"
)

// CreateEvent creates a new event in the database.
func CreateEvent(db *gorm.DB, name string, id uint) error {
	event := &entity.Event{Name: name, TeamId: id}
	if err := db.Create(event).Error; err != nil {
		return err
	}
	return nil
}

// GetEvent retrieves an event by ID.
func GetEvent(db *gorm.DB, id uint) (*entity.Event, error) {

	var event entity.Event
	if err := db.First(&event, id).Error; err != nil {
		if err != nil {
			return nil, nil
		}
		return nil, err
	}
	return &event, nil
}

// UpdateEvent updates an existing event in the database.
func UpdateEvent(db *gorm.DB, id uint, name string) error {

	if err := db.Model(&entity.Event{}).Where("id = ?", id).Update("Name", name).Error; err != nil {
		return err
	}
	return nil
}

// DeleteEvent deletes an event by ID.
// func DeleteEvent(db *gorm.DB, id uint) error {

//		if err := db.Where("id = ?", id).Delete(&entity.Event{}).Error; err != nil {
//			return err
//		}
//		return nil
//	}
func DeleteEvent(db *gorm.DB, id uint) error {

	if err := db.Model(&entity.Event{}).Where("id = ?", id).Update("deleted_at", time.Now().Format("2006-01-02 15:04:05")).Error; err != nil {
		return err
	}
	return nil
}
