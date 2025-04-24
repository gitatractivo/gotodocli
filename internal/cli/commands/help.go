package commands

import "fmt"

// PrintHelp prints the help message
func PrintHelp() error {
	fmt.Println("Usage: todo <command> [options]")
	fmt.Println("Commands:")
	fmt.Println("  add <task>        Add a new task")
	fmt.Println("  list              List all tasks")
	fmt.Println("  done <id>         Mark a task as done")
	fmt.Println("  server start      Start the API server in the background")
	fmt.Println("  server stop       Stop the running API server")
	fmt.Println("  server status     Check if the server is running")
	fmt.Println("  --help, -h        Show this help message")
	fmt.Println("  --version, -v     Show version information")
	return nil
}