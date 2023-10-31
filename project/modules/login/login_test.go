package login

import (
	"project/database"
	"testing"
)

func TestLoginCRUD(t *testing.T) {
	t.Run("TestCreateLogin", testCreateLogin)
	t.Run("TestGetLogin", testGetLogin)
	t.Run("TestUpdateLogin", testUpdateLogin)
	t.Run("TestDeleteLogin", testDeleteLogin)
}

func testCreateLogin(t *testing.T) {
	db, _ := database.InitDB()
	// Test CreateLogin function
	err := CreateLogin(db, "testuser", "password")
	if err != nil {
		t.Errorf("Error creating login: %v", err)
	}
}

func testGetLogin(t *testing.T) {
	db, _ := database.InitDB()
	// Test GetLogin function
	exists, err := GetLogin(db, "testuser", "password")
	if err != nil {
		t.Errorf("Error getting login: %v", err)
	}
	if exists == nil {
		t.Errorf("Login not found")
	}
}

func testUpdateLogin(t *testing.T) {
	db, _ := database.InitDB()
	// Test UpdateLogin function
	err := UpdateLogin(db, "testuser", "newpassword")
	if err != nil {
		t.Errorf("Error updating login: %v", err)
	}
}

func testDeleteLogin(t *testing.T) {
	db, _ := database.InitDB()
	// Test DeleteLogin function
	err := DeleteLogin(db, "testuser")
	if err != nil {
		t.Errorf("Error deleting login: %v", err)
	}
}
