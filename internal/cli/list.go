// add tablewriter
package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gitatractivo/gotodocli/internal/cli/utils"
	"github.com/gitatractivo/gotodocli/internal/models"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  `List all of your todo tasks.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiURL := viper.GetString("api_base_url")
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

		table := tablewriter.NewWriter(os.Stdout)

		table.SetHeader([]string{"ID", "Completed", "Title"})

		for _, task := range tasks {
			// fmt.Printf("%d\t%s\t\t%s\n", task.ID,
			table.Append([]string{
				fmt.Sprintf("%d", task.ID),
				utils.GetStatusEmoji(task.Completed),
				task.Title,
			})
		}

		table.Render()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
