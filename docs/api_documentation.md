# Task Management API Documentation

## Endpoints

### GET /tasks
Retrieve a list of all tasks.
- **Response:**
  - Status: 200 OK
  - Body: JSON array of tasks

### GET /tasks/:id
Retrieve the details of a specific task.
- **Parameters:**
  - `id` (string): The ID of the task
- **Response:**
  - Status: 200 OK
  - Body: JSON object of the task
  - Status: 404 Not Found (if the task does not exist)

### POST /tasks
Create a new task.
- **Request Body:**
  - `title` (string): The title of the task
  - `description` (string): The description of the task
  - `due_date` (time): The due date of the task
  - `status` (string): The status of the task
- **Response:**
  - Status: 201 Created
  - Body: JSON object of the created task

### PUT /tasks/:id
Update a specific task.
- **Parameters:**
  - `id` (string): The ID of the task
- **Request Body:**
  - `title` (string): The new title of the task
  - `description` (string): The new description of the task
  - `due_date` (time): The new due date of the task.
  - `status` (string): The new status of the task
- **Response:**
  - Status: 200 OK
  - Body: JSON object of the updated task
  - Status: 404 Not Found (if the task does not exist)

### DELETE /tasks/:id
Delete a specific task.
- **Parameters:**
  - `id` (string): The ID of the task
- **Response:**
  - Status: 204 No Content
  - Status: 404 Not Found (if the task does not exist)
### API DOCUMENTION: https://documenter.getpostman.com/view/37289771/2sA3rwLu59