package cli

import (
	"errors"
	"fmt"
)

// App represents the CLI application
type App struct {
	Version string
	Commit string
	BuildDate string
}

// NewApp creates a new CLI application also take version, commit and build date as parameters
func NewApp(version, commit, buildDate string) *App {
	return &App{
		Version: version,
		Commit: commit,
		BuildDate: buildDate,
	}
}


func (a *App) Run(args []string) error {

	if len(args) == 1 && (args[0] == "-v" || args[0] == "--version") {
		// TODO: Get version from build flags
	
		fmt.Println("Version:", a.Version)
		return nil
	}else if len(args) == 1 && (args[0] == "-h" || args[0] == "--help") {
		fmt.Println("Help")
		return nil
		// return a.printHelp()
	}else if len(args)<2{
		fmt.Println("Expected 'add', 'list', or 'done' subcommands")
		return errors.New("No command provided. Insufficient arguments")
	}
	switch args[1] {
	case "add":
		return a.handleAddTask(args[2:])
	case "list":
		return a.handleListTasks()
	case "done":
		return a.handleMarkTaskAsDone(args[2:])
	default:
		return errors.New("Invalid command. Expected 'add', 'list', or 'done'")
	}

}


func (a *App) printHelp() error {
	// TODO: Add help text and add colors to the output with a good output table 
	fmt.Println("Usage: gotodo <command> [options]")
	fmt.Println("Commands:")
	fmt.Println("  add <task>    Add a new task")
	fmt.Println("  list         List all tasks")
	fmt.Println("  done <id>    Mark a task as done")
	return nil
}

func (a *App) handleAddTask(args []string) error {
	if(len(args)<1){
		return errors.New("Missing task description. Usage: gotodo add <task>")
	}
	task := args[0]
	fmt.Println("Adding task:", task)
	return nil
}

func (a *App) handleListTasks() error {
	fmt.Println("Listing all tasks")
	return nil
}

func (a *App) handleMarkTaskAsDone(args []string) error {
	if(len(args)<1){
		return errors.New("Missing task ID. Usage: gotodo done <id>")
	}
	id := args[0]
	fmt.Println("Marking task as done:", id)
	return nil
}