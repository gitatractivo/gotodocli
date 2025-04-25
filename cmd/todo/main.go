package main

import (
	"fmt"
	"os"

	"github.com/gitatractivo/gotodocli/configs"
	"github.com/gitatractivo/gotodocli/internal/api"
	"github.com/gitatractivo/gotodocli/internal/cli"

)

var (
	Version   = "dev"
	Commit    = ""
	BuildDate = ""
)



func main() {
	configs.InitConfig(Version, Commit, BuildDate)
	if len(os.Args) >= 3 && os.Args[1] == "server" && os.Args[2] == "run" {

		if err := startServer(); err != nil {
			fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
			os.Exit(1)
		}
		return
	}
	cli.Execute()
}


func startServer() error {
	pid := os.Getpid()
	err := os.WriteFile(configs.GetConfig().ServerPidFile, []byte(fmt.Sprintf("%d", pid)), 0644)

	if err != nil {
		return fmt.Errorf("failed to write PID file: %w", err)
	}

	fmt.Printf("Starting Todo API server (version %s)\n", Version)
	fmt.Printf("Server PID: %d\n", pid)

	server := api.NewServer(configs.GetConfig().ServerPort)
	if err := server.Run(); err != nil {
		return fmt.Errorf("server error: %w", err)
	}

	return nil

}
