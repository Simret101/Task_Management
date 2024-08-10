package models

// Task represents a task with its ID, title, description, due date, and status.
type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	Status      string `json:"status"` // Valid statuses might be "completed", "inprogress", "started"
}

// ErrorResponse represents a generic error response format
type ErrorResponse struct {
	Error string `json:"error"`
}


