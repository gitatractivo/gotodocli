package server

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/gitatractivo/gotodocli/configs"
)

// HandleStartServer handles the 'server start' command
func HandleStartServer() error {
	// Check if server is already running
	if IsServerRunning() {
		return errors.New("server is already running")
	}

	pidFile := configs.GetConfig().ServerPidFile

	// Start server as a background process
	cmd := exec.Command(os.Args[0], "server", "run")

	// Detach process from terminal
	cmd.Stdout = nil
	cmd.Stderr = nil
	cmd.Stdin = nil

	// Start the process
	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	// Write PID to file
	err = os.WriteFile(pidFile, []byte(fmt.Sprintf("%d", cmd.Process.Pid)), 0644)
	if err != nil {
		return fmt.Errorf("server started but failed to write PID file: %w", err)
	}

	fmt.Println("Server started successfully in background")

	// Wait for server to be ready
	port := configs.GetConfig().ServerPort
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
}