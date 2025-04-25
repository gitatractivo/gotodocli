
package cli

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serverStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check server status",
	Long:  `Check if the API server is running.`,
	Run: func(cmd *cobra.Command, args []string) {
		if isServerRunning() {
			fmt.Println("Server is running")
		} else {
			fmt.Println("Server is not running")
		}
	},
}

// isServerRunning checks if the server is running
func isServerRunning() bool {
	pidFile := viper.GetString("server_pid_file")

	// Check if PID file exists
	if _, err := os.Stat(pidFile); os.IsNotExist(err) {
		return false
	}

	// Read PID from file
	data, err := os.ReadFile(pidFile)
	if err != nil {
		return false
	}

	// Parse PID
	pid, err := strconv.Atoi(string(data))
	if err != nil {
		return false
	}

	fmt.Printf("Server PID: %d\n", pid)
	
	// Check if process exists
	process, err := os.FindProcess(pid)
	if err != nil {
		os.Remove(pidFile)
		return false
	}

	// Platform-specific check
	if runtime.GOOS == "windows" {
		// On Windows, FindProcess only succeeds if the process exists
		return true
	} else {
		// On Unix, use signal 0 to check process existence
		err = process.Signal(syscall.Signal(0))
		return err == nil
	}
}

func init() {
	serverCmd.AddCommand(serverStatusCmd)
}

