package models

// Task represents a task with its ID, title, description, due date, and status.
type Task struct {
	ID          int    `json:"id"`          // Unique identifier for the task
	Title       string `json:"title"`       // Title of the task
	Description string `json:"description"` // Detailed description of the task
	DueDate     string `json:"due_date"`    // Due date for the task in YYYY-MM-DD format
	Status      string `json:"status"`      // Status of the task ( "completed", "inprogress", "started")
}

