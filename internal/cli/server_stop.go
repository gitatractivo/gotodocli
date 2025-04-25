
package cli

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serverStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the API server",
	Long:  `Stop the running API server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		pidFile := viper.GetString("server_pid_file")
		
		// Read the PID file
		data, err := os.ReadFile(pidFile)
		if err != nil {
			return fmt.Errorf("server not running or PID file not found")
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
	},
}

func init() {
	serverCmd.AddCommand(serverStopCmd)
}

