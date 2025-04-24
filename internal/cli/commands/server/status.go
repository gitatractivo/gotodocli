package server

import "fmt"

// HandleServerStatus handles the 'server status' command
func HandleServerStatus() error {
	if IsServerRunning() {
		fmt.Println("Server is running")
	} else {
		fmt.Println("Server is not running")
	}
	return nil
}