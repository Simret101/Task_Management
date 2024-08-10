package router

import (
	"example/taskManager/controllers"

	_ "example/taskManager/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//@title Task API
//@version 1.0
//@description This is an api for task management system
//@contact.name Simret Belete
//@contact.url https://github.com/Simret101
//@contact.email semretb4@gmail.com

// @host localhost:8080
// @BasePath /api/v1
//// @Router /tasks [get]
// SetupRouter creates the router and initializes the services, repositories, use cases, and controllers
// It sets up the routes for the API endpoints

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	tasks := v1.Group("/tasks")
	{
		
		tasks.GET("", ginSwagger.WrapHandler(swaggerFiles.Handler), controllers.GetAllTasks)
		tasks.GET("/:id", controllers.GetTaskByID)
		tasks.POST("", controllers.CreateTask)
		tasks.PUT("/:id", controllers.UpdateTask)
		tasks.DELETE("/:id", controllers.DeleteTask)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}

