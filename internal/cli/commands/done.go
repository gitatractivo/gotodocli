package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gitatractivo/gotodocli/configs"
	"github.com/gitatractivo/gotodocli/internal/cli/utils"
	"github.com/gitatractivo/gotodocli/internal/models"
)

// HandleMarkTaskAsDone handles the 'done' command
func HandleMarkTaskAsDone(args []string) error {
	if len(args) < 1 {
		return errors.New("missing task ID. usage: todo done <id>")
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid task ID: %w", err)
	}

	apiURL := configs.GetConfig().ApiBaseUrl
	fmt.Println("Sending request to server...", apiURL+"/tasks/done/"+strconv.Itoa(id))

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/tasks/done/%d", apiURL, id), nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

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

	fmt.Printf("%d\t%s\t\t%s\n", updatedTask.ID, utils.GetStatusEmoji(updatedTask.Completed), updatedTask.Title)
	fmt.Println("--\t---------\t-----")
	fmt.Println("Task marked as done")
	
	return nil
}