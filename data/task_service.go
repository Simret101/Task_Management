package data

import (
	"errors"
	"example/taskManager/models"
)

// TaskService handles the business logic and data manipulation for tasks.
type TaskService struct {
	tasks map[string]*models.Task
}

// NewTaskService creates and returns a new TaskService with an initialized in-memory task store.
func NewTaskService() *TaskService {
	return &TaskService{
		tasks: make(map[string]*models.Task),
	}
}

// GetTasks retrieves all tasks from the in-memory store.
func (ts *TaskService) GetTasks() []*models.Task {
	var tasks []*models.Task
	for _, task := range ts.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

// GetTask retrieves a task by its ID from the in-memory store.
func (ts *TaskService) GetTask(id string) (*models.Task, error) {
	task, ok := ts.tasks[id] // Check if the task exists in the in-memory store.
	if !ok {
		return nil, errors.New("task not found")
	}
	return task, nil
}

// CreateTask adds a new task to the in-memory store.
func (ts *TaskService) CreateTask(task *models.Task) (*models.Task, error) {
	ts.tasks[task.ID] = task
	return task, nil
}

// UpdateTask updates an existing task in the in-memory store.
func (ts *TaskService) UpdateTask(id string, updatedTask *models.Task) (*models.Task, error) {
	task, ok := ts.tasks[id]
	if !ok {
		return nil, errors.New("task not found")
	}
	task.Title = updatedTask.Title
	task.Description = updatedTask.Description
	task.Duedate = updatedTask.Duedate
	task.Status = updatedTask.Status
	return task, nil
}

// DeleteTask removes a task by its ID from the in-memory store.
func (ts *TaskService) DeleteTask(id string) error {
	_, ok := ts.tasks[id]
	if !ok {
		return errors.New("task not found")
	}
	delete(ts.tasks, id)
	return nil
}
