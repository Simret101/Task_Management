package router

import (
	"example/taskManager/controllers"

	"github.com/gin-gonic/gin"
)

// initializes the router and sets up the API routes.
func SetupRouter() *gin.Engine {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	// Define the routes and their corresponding controller functions
	r.GET("/tasks", controllers.GetAllTasks)       // to retrieve all tasks
	r.GET("/tasks/:id", controllers.GetTaskByID)   // to retrieve a specific task by ID
	r.POST("/tasks", controllers.CreateTask)       // to create a new task
	r.PUT("/tasks/:id", controllers.UpdateTask)    // to update an existing task by ID
	r.DELETE("/tasks/:id", controllers.DeleteTask) // to delete a task by ID

	return r
}
