package main

import (
	"fmt"
	"os"
	"github.com/gitatractivo/gotodocli/internal/cli"
)


var (
	Version = "dev"
	Commit = ""
	BuildDate = ""
)

func main() {
	fmt.Printf("Todo CLI v%s (commit: %s, built at: %s)\n", Version, Commit, BuildDate)

	app := cli.NewApp(Version, Commit, BuildDate)
	if err := app.Run(os.Args); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	
}