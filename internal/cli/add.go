package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gitatractivo/gotodocli/internal/cli/utils"
	"github.com/gitatractivo/gotodocli/internal/models"
	"github.com/olekukonko/tablewriter"
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
		//convert to table
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Completed", "Title"})
		table.Append([]string{
			fmt.Sprintf("%d", createdTask.ID),
			utils.GetStatusEmoji(createdTask.Completed),
			createdTask.Title,
		})
		table.Render()
		fmt.Println("Task created successfully")
		
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

