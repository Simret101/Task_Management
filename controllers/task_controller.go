package controllers

import (
	"example/taskManager/data"
	"example/taskManager/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllTasks godoc
// @Summary Get all tasks
// @Description Get a list of all tasks
// @Tags tasks
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Task
// @Router /tasks [get]
func GetAllTasks(c *gin.Context) {
	tasks := data.GetAllTasks()
	c.JSON(http.StatusOK, tasks)
}

// GetTaskByID godoc
// @Summary Retrieve a specific task
// @Description Get a task by its ID
// @Tags Tasks
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} models.Task
// @Router /tasks/{id} [get]
func GetTaskByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "invalid task ID"})
		return
	}
	task, err := data.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": "task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

// CreateTask godoc
// @Summary Create a new task
// @Description Add a new task to the system
// @Tags Tasks
// @Accept json
// @Produce json
// @Param task body models.Task true "Task to create"
// @Success 201 {object} models.Task
// @Router /tasks [post]
func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "invalid input"})
		return
	}
	data.CreateTask(&task)
	c.JSON(http.StatusCreated, task)
}

// UpdateTask godoc
// @Summary Update an existing task
// @Description Modify the details of an existing task by its ID
// @Tags Tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Param task body models.Task true "Updated task details"
// @Success 200 {object} models.Task
// @Router /tasks/{id} [put]
func UpdateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "invalid task ID"})
		return
	}
	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "invalid input"})
		return
	}
	if err := data.UpdateTask(id, &updatedTask); err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": "task not found"})
		return
	}
	c.JSON(http.StatusOK, updatedTask)
}

// DeleteTask godoc
// @Summary Delete a task
// @Description Remove a task from the system by its ID
// @Tags Tasks
// @Param id path int true "Task ID"
// @Success 204 "No Content"
// @Router /tasks/{id} [delete]
func DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "invalid task ID"})
		return
	}
	if err := data.DeleteTask(id); err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": "task not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}


