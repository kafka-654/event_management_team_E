package events

import (
	"project/database"
	"testing"
)

func TestEventsCRUD(t *testing.T) {
	t.Run("TestCreateEvent", testCreateEvent)
	// t.Run("TestGetEvent", testGetEvent)
	t.Run("TestUpdateEvent", testUpdateEvent)
	t.Run("TestDeleteEvent", testDeleteEvent)
}

func testCreateEvent(t *testing.T) {
	// Test CreateEvent function
	db, _ := database.InitDB()
	if err := CreateEvent(db, "Conference", 1); err != nil {
		t.Errorf("Error creating event: %v", err)
	}
}

func testGetEvent(t *testing.T) {
	// Test GetEvent function
	db, _ := database.InitDB()
	event, err := GetEvent(db, 1) // Replace with a valid event ID
	if err != nil {
		t.Errorf("Error getting event: %v", err)
	}
	if event == nil {
		t.Errorf("Event not found")
	}
}

func testUpdateEvent(t *testing.T) {
	db, _ := database.InitDB()
	// Test UpdateEvent function
	err := UpdateEvent(db, 1, "Workshop") // Replace with a valid event ID
	if err != nil {
		t.Errorf("Error updating event: %v", err)
	}
}

func testDeleteEvent(t *testing.T) {
	db, _ := database.InitDB()
	// Test DeleteEvent function
	err := DeleteEvent(db, 1) // Replace with a valid event ID
	if err != nil {
		t.Errorf("Error deleting event: %v", err)
	}
}
