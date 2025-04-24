package commands

import (
	"fmt"

	"github.com/gitatractivo/gotodocli/configs"
)

// PrintVersion prints the version information
func PrintVersion() error {
	config := configs.GetConfig()
	fmt.Printf("Todo CLI v%s (commit: %s, built: %s)\n", config.Version, config.Commit, config.BuildDate)
	return nil
}