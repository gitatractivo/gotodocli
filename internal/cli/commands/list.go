package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gitatractivo/gotodocli/configs"
	"github.com/gitatractivo/gotodocli/internal/cli/utils"
	"github.com/gitatractivo/gotodocli/internal/models"
)

// HandleListTasks handles the 'list' command
func HandleListTasks() error {
	apiURL := configs.GetConfig().ApiBaseUrl
	fmt.Println("Sending request to server...", apiURL+"/tasks")
	
	resp, err := http.Get(apiURL + "/tasks")
	if err != nil {
		return fmt.Errorf("error connecting to server: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("server error (status %d): %s", resp.StatusCode, string(body))
	}

	var tasks []models.Task
	err = json.NewDecoder(resp.Body).Decode(&tasks)
	if err != nil {
		return fmt.Errorf("error parsing server response: %w", err)
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return nil
	}

	fmt.Println("ID\tCompleted\tTitle")
	fmt.Println("--\t---------\t-----")
	
	for _, task := range tasks {
		fmt.Printf("%d\t%s\t\t%s\n", task.ID, utils.GetStatusEmoji(task.Completed), task.Title)
	}

	return nil
}