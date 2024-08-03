package router

import (
	"example/taskManager/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(taskController *controllers.TaskController) *gin.Engine {
	r := gin.Default()

	tasks := r.Group("/tasks")
	{
		tasks.GET("", taskController.GetTasks)
		tasks.GET("/:id", taskController.GetTask)
		tasks.POST("", taskController.CreateTask)
		tasks.PUT("/:id", taskController.UpdateTask)
		tasks.DELETE("/:id", taskController.DeleteTask)
	}

	return r
}
