Here's a GitHub `README.md` for the Console Task Manager application:

# Console Task Manager

The Console Task Manager is a command-line application written in Go for managing tasks. It provides functionality to add, view, update, and delete tasks, as well as mark tasks as complete. Tasks are stored in a JSON file for persistence.

## Features
- **Add Task**: Add a new task with details including title, description, due date, and status.
- **View Tasks**: Display all tasks in a tabular format for easy review.
- **Get Task by ID**: Retrieve and view a specific task by its ID.
- **Update Task**: Modify the details of an existing task.
- **Remove Task**: Delete a task from the list by its ID.
- **Mark Task as Complete**: Update the status of a task to 'completed'.

## Installation
To get started with the Console Task Manager, follow these steps:

1. **Clone the repository**:
   ```sh
   git clone [https://github.com/Simret101/Task_Management]
   ```
2. **Navigate to the project directory**:
   ```sh
   cd task-manager
   ```
3. **Install dependencies**:
   ```sh
   go mod tidy
   ```

## Usage
1. **Run the application**:
   ```sh
   go run main.go
   ```
2. **Interact with the application** by selecting options from the main menu:
   ```
   Task List
   1. Add Task.
   2. View all Tasks.
   3. Get Task by ID.
   4. Update Task.
   5. Remove Task.
   6. Mark Task as Complete.
   7. Exit
   Enter your choice:
   ```
---

## Contributing
Contributions are welcome! Please open an issue or submit a pull request if you'd like to contribute to this project.

## Contact
For any questions or feedback, please reach out to [semretb74@gmail.com]



