
# Task Management API Documentation

This document provides details about the Task Management API endpoints, including how to interact with the API to manage tasks. For more comprehensive API documentation, including detailed descriptions, examples, and schema, refer to the Swagger and Postman documentation links provided below.

## API Documentation Links

- **Swagger Documentation:** [Swagger UI](http://localhost:8080/swagger/index.html#/)
- **Postman Documentation:** [Postman Collection](https://documenter.getpostman.com/view/37289771/2sA3rwNZnx)

#
This structure gives users quick access to the basic endpoint details while also directing them to the more detailed Swagger and Postman documentation.
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


