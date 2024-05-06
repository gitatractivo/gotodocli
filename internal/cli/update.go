// add tablewriter
package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gitatractivo/gotodocli/internal/models"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)


var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a task",
	Long:  `Update a task by ID.`,
	
	RunE: func(cmd *cobra.Command, args []string) error {
		apiURL := viper.GetString("api_base_url")


		fmt.Println("Sending request to server...", apiURL+"/tasks/"+args[0])

		resp, err := http.Get(apiURL + "/tasks/" + args[0])
		if err != nil {
			return fmt.Errorf("error connecting to server: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			return fmt.Errorf("server error (status %d): %s", resp.StatusCode, string(body))
		}

		var task models.Task
		err = json.NewDecoder(resp.Body).Decode(&task)
		if err != nil {
			return fmt.Errorf("error decoding response: %w", err)
		}


		
		

		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringP("title", "t", "", "The new title of the task")
	updateCmd.Flags().StringP("description", "d", "", "The new description of the task")
	updateCmd.Flags().StringP("due-date", "D", "", "The new due date of the task")
	updateCmd.Flags().StringP("priority", "p", "", "The new priority of the task")
	updateCmd.Flags().StringP("project", "P", "", "The new project of the task")
	updateCmd.Flags().StringP("tags", "T", "", "The new tags of the task")
	

}