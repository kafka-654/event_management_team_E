package main

import (
	"log"
	"project/database"
	"project/modules/events"
	"project/modules/login"
	"project/modules/team"

	"gorm.io/gorm"
)

func main() {
	// Initialize the database
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Error initializing the database: %v", err)
	}
	loginExample(db)
	teamExample(db)
	eventsExample(db)

	// loginn, err := login.GetLogin(db, "usernamee", "1")
	// if err != nil {
	// 	log.Fatalf("Error creating login: %v", err)
	// }

	// // log.Fatalf(loginn)
	// fmt.Println("Login Succeful", loginn)

	// Perform CRUD operations or other actsions as needed
	// Example usage:

	// Close the database connection when done
	// database.CloseDB()
}

func loginExample(db *gorm.DB) {
	// Example usage of login module
	// You can call CreateLogin, GetLogin, UpdateLogin, and DeleteLogin functions here
	// Call the CreateLogin function with valid parameters
	if err := login.CreateLogin(db, "usernamee", "passwordd"); err != nil {
		log.Fatalf("Error creating login: %v", err)
	}
	if err := login.CreateLogin(db, "name", "pass"); err != nil {
		log.Fatalf("Error creating login: %v", err)
	}
	if err := login.CreateLogin(db, "name", "pass"); err != nil {
		log.Fatalf("Error creating login: %v", err)
	}
	if err := login.CreateLogin(db, "name1", "pass"); err != nil {
		log.Fatalf("Error creating login: %v", err)
	}
	if err := login.DeleteLogin(db, "usernamee"); err != nil {
		log.Fatalf("Error creating login: %v", err)
	}
	if err := login.UpdateLogin(db, "usernamee", "1"); err != nil {
		log.Fatalf("Error creating login: %v", err)
	}
}

func teamExample(db *gorm.DB) {
	// Example usage of team module
	// You can call CreateTeam, GetTeam, UpdateTeam, and DeleteTeam functions here
	if err := team.CreateTeam(db, "Eagle", 1); err != nil {
		log.Fatalf("Error creating login: %v", err)
	}
	if err := team.DeleteTeam(db, 2); err != nil {
		log.Fatalf("Error creating login: %v", err)
	}

}

func eventsExample(db *gorm.DB) {
	// Example usage of events module
	// You can call CreateEvent, GetEvent, UpdateEvent, and DeleteEvent functions here
	if err := events.CreateEvent(db, "shadi", 1); err != nil {
		log.Fatalf("Error creating login: %v", err)
	}
	if err := events.CreateEvent(db, "Pelli", 2); err != nil {
		log.Fatalf("Error creating login: %v", err)
	}
	if err := events.CreateEvent(db, "sangeet", 5); err != nil {
		log.Fatalf("Error creating login: %v", err)
	}
	if err := events.UpdateEvent(db, 1, "nikka"); err != nil {
		log.Fatalf("Error creating login: %v", err)
	}
	if err := events.DeleteEvent(db, 10); err != nil {
		log.Fatalf("Error creating login: %v", err)
	}
}
