
package cli

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serverStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the API server",
	Long:  `Start the API server in the background.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Check if server is already running
		if isServerRunning() {
			return fmt.Errorf("server is already running")
		}

		pidFile := viper.GetString("server_pid_file")

		// Start server as a background process
		execCmd := exec.Command(os.Args[0], "server", "run")

		// Detach process from terminal
		execCmd.Stdout = nil
		execCmd.Stderr = nil
		execCmd.Stdin = nil

		// Start the process
		err := execCmd.Start()
		if err != nil {
			return fmt.Errorf("failed to start server: %w", err)
		}

		// Write PID to file
		err = os.WriteFile(pidFile, []byte(fmt.Sprintf("%d", execCmd.Process.Pid)), 0644)
		if err != nil {
			return fmt.Errorf("server started but failed to write PID file: %w", err)
		}

		fmt.Println("Server started successfully in background")

		// Wait for server to be ready
		port := viper.GetString("server_port")
		healthURL := fmt.Sprintf("http://localhost:%s/health", port)
		
		for i := 0; i < 5; i++ {
			time.Sleep(500 * time.Millisecond)
			resp, err := http.Get(healthURL)
			if err == nil && resp.StatusCode == http.StatusOK {
				fmt.Println("Server is ready")
				return nil
			}
		}

		fmt.Println("Server started but not responding yet, it may still be initializing")
		return nil
	},
}

func init() {
	serverCmd.AddCommand(serverStartCmd)
}

