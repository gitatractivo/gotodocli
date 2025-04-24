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
	fmt.Printf("Todo CLI v%s (commit: %s, built at: %s)\n", Version, Commit, BuildDate)

	if len(os.Args) >= 3 && os.Args[1] == "server" && os.Args[2] == "run" {

		if err := startServer(); err != nil {
			fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
			os.Exit(1)
		}
		return
	}
	app := cli.NewApp()
	if err := app.Run(os.Args); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
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
