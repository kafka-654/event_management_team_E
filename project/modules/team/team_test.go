package team

import (
	"project/database"
	"testing"
)

func TestTeamCRUD(t *testing.T) {
	t.Run("TestCreateTeam", testCreateTeam)
	// t.Run("TestGetTeam", testGetTeam)
	t.Run("TestUpdateTeam", testUpdateTeam)
	t.Run("TestDeleteTeam", testDeleteTeam)
}

func testCreateTeam(t *testing.T) {
	db, _ := database.InitDB()
	// Test CreatePerson function
	err := CreateTeam(db, "John Doe", 5)
	if err != nil {
		t.Errorf("Error creating team: %v", err)
	}
}

func testGetTeam(t *testing.T) {

	db, _ := database.InitDB()  // Test GetPerson function
	team, err := GetTeam(db, 1) // Replace with a valid person ID
	if err != nil {
		t.Errorf("Error getting person: %v", err)
	}
	if team == nil {
		t.Errorf("team not found")
	}
}

func testUpdateTeam(t *testing.T) {
	db, _ := database.InitDB()
	// Test UpdatePerson function
	err := UpdateTeam(db, 1, "Jane Doe") // Replace with a valid person ID
	if err != nil {
		t.Errorf("Error updating person: %v", err)
	}
}

func testDeleteTeam(t *testing.T) {
	// Test DeletePerson function
	db, _ := database.InitDB()
	err := DeleteTeam(db, 1) // Replace with a valid person ID
	if err != nil {
		t.Errorf("Error deleting person: %v", err)
	}
}
