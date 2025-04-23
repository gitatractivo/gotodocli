package cli

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gitatractivo/gotodocli/internal/models"
)

// App represents the CLI application
type App struct {
	Version   string
	Commit    string
	BuildDate string
}
const (
	apiURL = "http://localhost:8080/api/v1"
	pidFile = "/tmp/todo-server.pid"
)


// NewApp creates a new CLI application also take version, commit and build date as parameters
func NewApp(version, commit, buildDate string) *App {
	return &App{
		Version:   version,
		Commit:    commit,
		BuildDate: buildDate,
	}
}

func (a *App) Run(args []string) error {
if len(args) < 2 {
		fmt.Println("Expected 'add', 'list', 'done', 'server' subcommands")
		return errors.New("no command provided. insufficient arguments")
	}

	switch args[1] {
	case "add":
		return a.handleAddTask(args[2:])
	case "list":
		return a.handleListTasks()
	case "done":
		return a.handleMarkTaskAsDone(args[2:])
	case "server":
		if len(args) < 3 {
			return errors.New("expected 'start' or 'stop' after 'server'")
		}
		switch args[2] {
		case "start":
			return a.handleStartServer()
		case "stop":
			return a.handleStopServer()
		case "status":
			return a.handleServerStatus()
		default:
			return errors.New("invalid server command. expected 'start', 'stop', or 'status'")
		}
	case "--help", "-h":
		return a.printHelp()
	case "--version", "-v":
		return a.printVersion()
	default:
		return errors.New("invalid command. expected 'add', 'list', 'done', or 'server'")
	}

}

func (a *App) printHelp() error {
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

func (a *App) printVersion() error {
	fmt.Printf("Todo CLI v%s (commit: %s, built: %s)\n", a.Version, a.Commit, a.BuildDate)
	return nil
}

func (a *App) handleAddTask(args []string) error {
	if len(args) < 1 {
		return errors.New("Missing task description. Usage: gotodo add <task>")
	}
	taskTitle := args[0]
	taskDescription := strings.Join(args[1:], " ")
	task :=models.Task{
		Title: taskTitle,
		Description: taskDescription,
		Completed: false,
	}
	//convert task to json 

	jsonData, err := json.Marshal(task)
	if err != nil {
		return fmt.Errorf("error preparing task data: %w", err)
	}
	resp, err := http.Post(apiURL+"/tasks","application/json",bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error sending request to server: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body,_:=io.ReadAll(resp.Body)
		return fmt.Errorf("failed to create task: %s", string(body))
	}

	var createdTask models.Task
	if err := json.NewDecoder(resp.Body).Decode(&createdTask); err != nil {
		return fmt.Errorf("error parsing response: %w", err)
	}
	fmt.Println("ID\tCompleted\tTitle")
	fmt.Printf("%d\t%t\t%s\n", createdTask.ID, createdTask.Completed, createdTask.Title)
	fmt.Println("Task created successfully")
	return nil
		
	
}

func (a *App) handleListTasks() error {
	fmt.Println("Listing all tasks")
	return nil
}

func (a *App) handleMarkTaskAsDone(args []string) error {
	if len(args) < 1 {
		return errors.New("Missing task ID. Usage: gotodo done <id>")
	}
	id := args[0]
	fmt.Println("Marking task as done:", id)
	return nil
}
