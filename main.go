package main

import (
	"bufio"
	"example/taskManager/controllers"
	"example/taskManager/data"
	"example/taskManager/router"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Initialize the task service
	taskService := data.NewTaskService()

	// Create a new task controller with the task service
	taskController := controllers.NewTaskController(taskService)

	// Set up the router with the task controller
	r := router.SetupRouter(taskController)

	// Start the server in a goroutine so it doesn't block the main thread
	go func() {
		r.Run()
	}()

	// If you need to run input validation, you can call a separate function for that
	runInputValidation()
}

func runInputValidation() {
	year := inputValidation("Enter a year between 1980-2017:", 1980, 2017)
	month := inputValidation("Enter a month between 1-12:", 1, 12)
	day := inputValidation("Enter a date between 1-31:", 1, 31)
	hour := inputValidation("Enter an hour between 0-23:", 0, 23)
	minute := inputValidation("Enter a minute between 0-59:", 0, 59)
	second := inputValidation("Enter a second between 0-59:", 0, 59)

	id := inputIDValidation("Enter a numeric ID:")
	item := inputItemValidation("Enter an item (letters only):")
	status := inputStatusValidation("Enter a status (letters only):")

	userTime := time.Date(year, time.Month(month), day, hour, minute, second, 0, time.Local)
	fmt.Println("Due date:", userTime)
	fmt.Println("ID:", id)
	fmt.Println("Item:", item)
	fmt.Println("Status:", status)
}

func inputValidation(prompt string, min, max int) int {
	for {
		fmt.Println(prompt)
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		if input == "" {
			fmt.Println("Error: Input can't be empty!")
			continue
		}
		input = strings.TrimSpace(input)
		value, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error: Invalid input, please enter a number.")
			continue
		}
		if value < min || value > max {
			fmt.Printf("Error: Value must be between %d and %d.\n", min, max)
			continue
		}
		return value
	}
}

func inputIDValidation(prompt string) int {
	for {
		fmt.Println(prompt)
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		if input == "" {
			fmt.Println("Error: ID can't be empty!")
			continue
		}
		input = strings.TrimSpace(input)
		if _, err := strconv.Atoi(input); err != nil {
			fmt.Println("Error: ID must be a number.")
			continue
		}
		value, _ := strconv.Atoi(input)
		return value
	}
}

func inputItemValidation(prompt string) string {
	for {
		var input string
		fmt.Println(prompt)
		reader := bufio.NewReader(os.Stdin)
		input, _ = reader.ReadString('\n')
		if input == "" {
			fmt.Println("Error: Item can't be empty!")
			continue
		}
		input = strings.TrimSpace(input)
		if !isLettersOnly(input) {
			fmt.Println("Error: Item must contain only letters.")
			continue
		}
		return input
	}
}

func inputStatusValidation(prompt string) string {
	for {
		var input string
		fmt.Println(prompt)
		reader := bufio.NewReader(os.Stdin)
		input, _ = reader.ReadString('\n')
		if input == "" {
			fmt.Println("Error: Status can't be empty!")
			continue
		}
		input = strings.TrimSpace(input)
		if !isLettersOnly(input) {
			fmt.Println("Error: Status must contain only letters.")
			continue
		}
		return input
	}
}

func isLettersOnly(s string) bool {
	return regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(s)
}
