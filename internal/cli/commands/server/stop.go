package server

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/gitatractivo/gotodocli/configs"
)

// HandleStopServer handles the 'server stop' command
func HandleStopServer() error {
	pidFile := configs.GetConfig().ServerPidFile
	
	// Read the PID file
	data, err := os.ReadFile(pidFile)
	if err != nil {
		return errors.New("server not running or PID file not found")
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
}