
package cli

import (
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Manage the API server",
	Long:  `Start, stop, or check the status of the API server.`,
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
