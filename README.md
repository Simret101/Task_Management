# Task Management API

## Overview
The Task Management API is a RESTful service designed to manage tasks efficiently. It allows users to create, retrieve, update, and delete tasks. The API is built using the [Gin](https://github.com/gin-gonic/gin) framework for Go, and it supports JSON data format.

## Features
- **CRUD Operations**: Create, read, update, and delete tasks.
- **Task Attributes**: Each task includes a title, description, due date, and status.
- **Structured and Clean Code**: Organized using MVC architecture with clear separation of concerns.
- **Auto-Generated Documentation**: Swagger-based documentation for easy API exploration.
- **Testing**: Postman collection for testing the API endpoints.

## Table of Contents
- [Overview](#overview)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Swagger Documentation](#swagger-documentation)
- [Postman Collection](#postman-collection)
- [Contributing](#contributing)

## Installation
1. **Clone the Repository:**
    ```sh
    git clone [https://github.com/Simret101/Task_Management]
    cd task-manager-api
    ```
2. **Install Dependencies:**
    Ensure you have [Go](https://golang.org/) installed, then run:
    ```sh
    go mod tidy
    ```
3. **Set Up Environment Variables:**
   Create a `.env` file with the necessary environment variables such as database connection details.
   
4. **Run the API:**
    ```sh
    go run main.go
    ```
5. **Access the API:**
   The API will be running at `http://localhost:8080`.

## Usage
After starting the API, you can use tools like [Postman](https://www.postman.com/) or `curl` to interact with the API.

### Example Request:
```sh
curl -X POST http://localhost:8080/tasks \
-H "Content-Type: application/json" \
-d '{"title": "New Task", "description": "Description of the task", "due_date": "2024-08-31", "status": "Pending"}'
```

### Example Response:
```json
{
    "id": 1,
    "title": "New Task",
    "description": "Description of the task",
    "due_date": "2024-08-31",
    "status": "Pending"
}
```

## API Endpoints
The API exposes the following endpoints:

- **GET /tasks**: Retrieve all tasks.
- **GET /tasks/:id**: Retrieve a specific task by ID.
- **POST /tasks**: Create a new task.
- **PUT /tasks/:id**: Update a task by ID.
- **DELETE /tasks/:id**: Delete a task by ID.

For detailed information on each endpoint, refer to the [http://localhost:8080/swagger/index.html#/].

## Swagger Documentation
The API is documented using Swagger. You can explore the documentation and test endpoints directly in your browser.

- **Access Swagger UI**: `[http://localhost:8080/swagger/index.html#/]`

## Postman Collection
A Postman collection is provided for testing the API endpoints.

- **Access Postman Documentation**: [Postman Documentation][https://documenter.getpostman.com/view/37289771/2sA3rwNZnx]
- **Import Collection**: The collection can be imported into your Postman workspace for easy testing.

## Contributing
Contributions are welcome! If you have ideas for improvements or find any issues, please open a pull request or issue. 



