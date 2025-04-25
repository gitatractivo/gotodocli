package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gitatractivo/gotodocli/internal/models"
	"github.com/gitatractivo/gotodocli/internal/cli/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var doneCmd = &cobra.Command{
	Use:   "done [id]",
	Short: "Mark a task as done",
	Long:  `Mark a specific task as completed.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid task ID: %w", err)
		}

		apiURL := viper.GetString("api_base_url")
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
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}

