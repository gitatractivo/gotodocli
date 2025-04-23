package cli

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/gitatractivo/gotodocli/internal/models"
)

// App represents the CLI application
type App struct {
	Version   string
	Commit    string
	BuildDate string
}

const (
	apiURL  = "http://localhost:8080/api/v1"
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
	task := models.Task{
		Title:       taskTitle,
		Description: taskDescription,
		Completed:   false,
	}
	//convert task to json

	jsonData, err := json.Marshal(task)
	if err != nil {
		return fmt.Errorf("error preparing task data: %w", err)
	}
	resp, err := http.Post(apiURL+"/tasks", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error sending request to server: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to create task: %s", string(body))
	}

	var createdTask models.Task
	if err := json.NewDecoder(resp.Body).Decode(&createdTask); err != nil {
		return fmt.Errorf("error parsing response: %w", err)
	}
	fmt.Println("ID\tCompleted\tTitle")
	fmt.Println("--\t---------\t-----")

	fmt.Printf("%d\t%t\t%s\n", createdTask.ID, getStatusEmoji(createdTask.Completed), createdTask.Title)
	fmt.Println("--\t---------\t-----")

	fmt.Println("Task created successfully")
	return nil

}

func (a *App) handleListTasks() error {
	// Send GET request to API
	resp, err := http.Get(apiURL + "/tasks")
	if err != nil {
		return fmt.Errorf("error connecting to server: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("server error (status %d): %s", resp.StatusCode, string(body))
	}

	// Parse response
	var tasks []models.Task
	err = json.NewDecoder(resp.Body).Decode(&tasks)
	if err != nil {
		return fmt.Errorf("error parsing server response: %w", err)
	}

	// Display tasks
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return nil
	}

	fmt.Println("ID\tCompleted\tTitle")

	fmt.Println("--\t---------\t-----")
	for _, task := range tasks {

		fmt.Printf("%d\t[%s]\t\t%s\n", task.ID, getStatusEmoji(task.Completed), task.Title)
	}

	return nil
}

func (a *App) handleMarkTaskAsDone(args []string) error {
	if len(args) < 1 {
		return errors.New("missing task ID. usage: todo done <id>")
	}

	// Parse ID
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid task ID: %w", err)
	}

	// Create HTTP client with PUT method
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/tasks/%d/done", apiURL, id), nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error connecting to server: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("server error (status %d): %s", resp.StatusCode, string(body))
	}

	fmt.Println("ID\tCompleted\tTitle")
	fmt.Println("--\t---------\t-----")

	var updatedTask models.Task
	if err := json.NewDecoder(resp.Body).Decode(&updatedTask); err != nil {
		return fmt.Errorf("error parsing response: %w", err)
	}

	fmt.Printf("%d\t[%s]\t\t%s\n", updatedTask.ID, getStatusEmoji(updatedTask.Completed), updatedTask.Title)
	fmt.Println("--\t---------\t-----")

	fmt.Println("Task marked as done")
	return nil
}

func (a *App) handleStartServer() error {
	// Check if server is already running
	if isServerRunning() {
		return errors.New("server is already running")
	}

	// Start server as a background process
	cmd := exec.Command(os.Args[0], "server", "run")

	// Detach process from terminal
	cmd.Stdout = nil
	cmd.Stderr = nil
	cmd.Stdin = nil

	// Start the process
	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	// Write PID to file
	err = os.WriteFile(pidFile, []byte(fmt.Sprintf("%d", cmd.Process.Pid)), 0644)
	if err != nil {
		return fmt.Errorf("server started but failed to write PID file: %w", err)
	}

	fmt.Println("Server started successfully in background")

	// Wait for server to be ready
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		resp, err := http.Get("http://localhost:8080/health")
		if err == nil && resp.StatusCode == http.StatusOK {
			fmt.Println("Server is ready")
			return nil
		}
	}

	fmt.Println("Server started but not responding yet, it may still be initializing")
	return nil
}

func (a *App) handleStopServer() error {
	// Read the PID file
	data, err := os.ReadFile(pidFile)
	if err != nil {
		return errors.New("server not running or PID file not found")
	}

	// Parse PID
	pid, err := strconv.Atoi(string(data))
	if err != nil {
		return fmt.Errorf("invalid PID in file: %w", err)
	}

	// Find the process
	process, err := os.FindProcess(pid)
	if err != nil {
		return fmt.Errorf("process not found: %w", err)
	}

	// Send kill signal
	err = process.Signal(os.Interrupt) // Send SIGINT for graceful shutdown
	if err != nil {
		return fmt.Errorf("failed to stop server: %w", err)
	}

	// Remove PID file
	os.Remove(pidFile)

	fmt.Println("Server stopped successfully")
	return nil
}

func (a *App) handleServerStatus() error {
	if isServerRunning() {
		fmt.Println("Server is running")
	} else {
		fmt.Println("Server is not running")
	}
	return nil
}

// Helper function to check if server is running
func isServerRunning() bool {
	// Check if PID file exists
	if _, err := os.Stat(pidFile); os.IsNotExist(err) {
		return false
	}

	// Read PID from file
	data, err := os.ReadFile(pidFile)
	if err != nil {
		return false
	}

	// Parse PID
	pid, err := strconv.Atoi(string(data))
	if err != nil {
		return false
	}

	// Check if process exists
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}

	// On Unix systems, FindProcess always succeeds, so we need to send signal 0
	// to check if process exists
	err = process.Signal(os.Signal(nil))
	return err == nil
}

func getStatusEmoji(status bool) string {
	if status {
		return "✅"
	}
	return "❌"
}
