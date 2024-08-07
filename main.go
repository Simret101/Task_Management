package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"example/taskManager/console"
	"example/taskManager/router"
)

func clearScreen() {
	c := exec.Command("clear") // or "cls" for Windows
	c.Stdout = os.Stdout
	c.Run()
}

func printTitle(text string) {
	fmt.Printf("\033[1;34m%s\033[0m\n", text)
}

func printOption(text string) {
	fmt.Printf("\033[1;34m%s\033[0m\n", text)
}

func main() {
	clearScreen()
	printTitle("Task Management System")
	printOption("1. API Mode")
	printOption("2. Console Mode")

	reader := bufio.NewReader(os.Stdin)
	mode, _ := reader.ReadString('\n')
	mode = strings.TrimSpace(mode)

	switch mode {
	case "1":
		fmt.Println("Starting API mode...")
		r := router.SetupRouter()
		r.Run(":8080")

	case "2":
		fmt.Println("Starting console mode...")
		console.StartConsoleApp()
	default:
		fmt.Println("Invalid mode. Please choose 1 or 2.")
	}
}

