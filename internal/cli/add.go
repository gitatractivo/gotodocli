package cli


import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

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
		if taskDescription == "" {
			taskDescription = "No description"
		}

		priority, _ := cmd.Flags().GetInt("priority")
		category, _ := cmd.Flags().GetString("category")
		tags, _ := cmd.Flags().GetStringSlice("tags")
		// project, _ := cmd.Flags().GetString("project")
		// subtask, _ := cmd.Flags().GetString("subtask")
		// reminder, _ := cmd.Flags().GetString("reminder")
		//convert due date to time.Time
		
		dueDate, _ := cmd.Flags().GetString("due-date")
    timeOfDay, _ := cmd.Flags().GetString("time")
		
    // Parse the due date and time
    var dueDateParsed time.Time
    var err error
    
    if dueDate != "" {
        // Handle relative dates like "Today", "Tomorrow", etc.
        switch strings.ToLower(dueDate) {
        case "today":
            dueDateParsed = time.Now()
        case "tomorrow":
            dueDateParsed = time.Now().AddDate(0, 0, 1)
        case "next week":
            dueDateParsed = time.Now().AddDate(0, 0, 7)
        case "next month":
            dueDateParsed = time.Now().AddDate(0, 1, 0)
        case "next year":
            dueDateParsed = time.Now().AddDate(1, 0, 0)
        default:
            // Try to parse as a date string
            dueDateParsed, err = time.Parse("2006-01-02", dueDate)
            if err != nil {
                return fmt.Errorf("invalid date format: %w", err)
            }
        }
        
        // If time is provided, combine it with the date
        if timeOfDay != "" {
            // Parse time (handles formats like "3pm", "15:00", "3:00pm")
            timeOnly, err := parseTimeString(timeOfDay)
            if err != nil {
                return fmt.Errorf("invalid time format: %w", err)
            }
            
            // Combine date and time
            dueDateParsed = time.Date(
                dueDateParsed.Year(),
                dueDateParsed.Month(),
                dueDateParsed.Day(),
                timeOnly.Hour(),
                timeOnly.Minute(),
                0, 0,
                dueDateParsed.Location(),
            )
        }
    }
		var tagModels []models.Tag
		for _, tag := range tags {
				tagModels = append(tagModels, models.Tag{Name: tag})
		}
		task := models.Task{
			Title:       taskTitle,
			Description: &taskDescription,
			Completed:   false,
			Priority:    &priority,
			DueDate:     &dueDateParsed,
			Category:    &category,
			Tags:        tagModels,
			// Project:     &project,
			// Subtask:     &subtask,
			// Reminder:    &reminder,
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
	addCmd.Flags().IntP("priority", "p", 0, "Priority of the task")
	addCmd.Flags().StringP("due-date", "d", "", "Due date of the task")
	addCmd.Flags().StringP("category", "c", "", "Category of the task")
	addCmd.Flags().StringSliceP("tags", "t", []string{}, "Tags of the task")
	addCmd.Flags().StringP("project", "r", "", "Project of the task")
	addCmd.Flags().StringP("subtask", "s", "", "Subtask of the task")
	addCmd.Flags().StringP("reminder", "m", "", "Reminder of the task")
	//-h tag for time of day
	addCmd.Flags().StringP("time", "T", "", "Time of day of the task")
	rootCmd.AddCommand(addCmd)
}

func parseTimeString(timeStr string) (time.Time, error) {
    formats := []string{
        "3pm",
        "3:04pm",
        "15:04",
        "3PM",
        "3:04PM",
    }
    
    for _, format := range formats {
        t, err := time.Parse(format, timeStr)
        if err == nil {
            return t, nil
        }
    }
    
    return time.Time{}, fmt.Errorf("could not parse time: %s", timeStr)
}