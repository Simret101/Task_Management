package main

import (
	"bufio"
	"example/taskManager/console"
	_ "example/taskManager/docs"
	"example/taskManager/router"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// clearScreen clears the console screen.
func clearScreen() {
	// Command to clear the terminal screen
	c := exec.Command("clear") // Use "cls" for Windows
	c.Stdout = os.Stdout
	c.Run()
}

// printTitle prints a title with a specific color.
func printTitle(text string) {
	// Print text in bold blue
	fmt.Printf("\033[1;34m%s\033[0m\n", text)
}

// printOption prints an option with a specific color.
func printOption(text string) {
	// Print text in blue
	fmt.Printf("\033[1;34m%s\033[0m\n", text)
}

func main() {
	// Clear the screen and print the main menu
	clearScreen()
	printTitle("Task Management System")
	printOption("1. API Mode")
	printOption("2. Console Mode")

	// Read the user's choice
	reader := bufio.NewReader(os.Stdin)
	mode, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	mode = strings.TrimSpace(mode)

	// Handle the user's choice
	switch mode {
	case "1":
		fmt.Println("Starting API mode...")
		r := router.SetupRouter()
		if err := r.Run(); err != nil {
			fmt.Println("Error starting API server:", err)
		}

	case "2":
		fmt.Println("Starting console mode...")
		console.StartConsoleApp()
	default:
		fmt.Println("Invalid mode. Please choose 1 or 2.")
	}
}


