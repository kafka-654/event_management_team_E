// main.go

package main

import (
	"project_gin/controllers"
	"project_gin/database"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := database.InitDB(); err != nil {
		panic("Error initializing the database: " + err.Error())
	}
	defer database.CloseDB()

	r := gin.Default()

	// RESTful API for Login
	r.POST("/logins", controllers.CreateLogin)
	r.GET("/logins", controllers.GetLogins)
	r.GET("/logins/:id", controllers.GetLogin)
	r.PUT("/logins/:id", controllers.UpdateLogin)
	r.DELETE("/logins/:id", controllers.DeleteLogin)

	// RESTful API for Person
	r.POST("/people", controllers.CreatePerson)
	r.GET("/people", controllers.GetPeople)
	r.GET("/people/:id", controllers.GetPerson)
	r.PUT("/people/:id", controllers.UpdatePerson)
	r.DELETE("/people/:id", controllers.DeletePerson)

	// RESTful API for Event
	r.POST("/events", controllers.CreateEvent)
	r.GET("/events", controllers.GetEvents)
	r.GET("/events/:id", controllers.GetEvent)
	r.PUT("/events/:id", controllers.UpdateEvent)
	r.DELETE("/events/:id", controllers.DeleteEvent)

	r.Run(":8080")
}
