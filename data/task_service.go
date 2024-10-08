package data

import (
	"errors"
	"example/taskManager/models"
	"sync"
)

var (
	tasks  = []models.Task{}
	lastID = 0
	mu     sync.Mutex
)

// GetAllTasks retrieves all tasks from the in-memory storage.
func GetAllTasks() []models.Task {
	mu.Lock()
	defer mu.Unlock()
	return append([]models.Task(nil), tasks...) 
}

// GetTaskByID retrieves a specific task by its ID.
func GetTaskByID(id int) (*models.Task, error) {
	mu.Lock()
	defer mu.Unlock()
	for _, task := range tasks {
		if task.ID == id {
			return &task, nil
		}
	}
	return nil, errors.New("task not found")
}

// CreateTask adds a new task to the in-memory storage.
func CreateTask(task *models.Task) {
	mu.Lock()
	defer mu.Unlock()
	lastID++
	task.ID = lastID
	tasks = append(tasks, *task)
}

// UpdateTask modifies an existing task.
func UpdateTask(id int, updatedTask *models.Task) error {
	mu.Lock()
	defer mu.Unlock()
	for i, task := range tasks {
		if task.ID == id {
			updatedTask.ID = id
			tasks[i] = *updatedTask
			return nil
		}
	}
	return errors.New("task not found")
}

// DeleteTask removes a task from the in-memory storage.
func DeleteTask(id int) error {
	mu.Lock()
	defer mu.Unlock()
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}



