package commands

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gitatractivo/gotodocli/configs"
	"github.com/gitatractivo/gotodocli/internal/cli/utils"
	"github.com/gitatractivo/gotodocli/internal/models"
)

// HandleAddTask handles the 'add' command
func HandleAddTask(args []string) error {
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
	
	jsonData, err := json.Marshal(task)
	if err != nil {
		return fmt.Errorf("error preparing task data: %w", err)
	}
	
	apiURL := configs.GetConfig().ApiBaseUrl
	fmt.Println("Sending request to server...", apiURL+"/tasks")
	
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
	fmt.Printf("%d\t%s\t%s\n", createdTask.ID, utils.GetStatusEmoji(createdTask.Completed), createdTask.Title)
	fmt.Println("--\t---------\t-----")
	fmt.Println("Task created successfully")
	
	return nil
}