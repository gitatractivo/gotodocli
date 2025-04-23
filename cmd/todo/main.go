package main

import (
	"fmt"
	"os"

	"github.com/gitatractivo/gotodocli/internal/api"
	"github.com/gitatractivo/gotodocli/internal/cli"
)


var (
	Version = "dev"
	Commit = ""
	BuildDate = ""
)

func main() {
	fmt.Printf("Todo CLI v%s (commit: %s, built at: %s)\n", Version, Commit, BuildDate)


	if len(os.Args) >= 3 && os.Args[1] == "server" && os.Args[2] == "run" {
		// This mode is only called internally by the CLI's handleStartServer method
		fmt.Printf("Starting Todo API server (version %s)\n", Version)
		
		// Create and run the server using the implementation in server.go
		server := api.NewServer("8080")
		if err := server.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
			os.Exit(1)
		}
		return
	}
	app := cli.NewApp(Version, Commit, BuildDate)
	if err := app.Run(os.Args); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	
}