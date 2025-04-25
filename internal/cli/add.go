
package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gitatractivo/gotodocli/internal/models"
	"github.com/gitatractivo/gotodocli/internal/cli/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addCmd = &cobra.Command{
	Use:   "add [title] [description]",
	Short: "Add a new task",
	Long:  `Add a new task to your todo list.`,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
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
		
		apiURL := viper.GetString("api_base_url")
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
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

