package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// var showingVersion bool = false
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `Print version information about the GoToDo CLI.`,
	Run: func(cmd *cobra.Command, args []string) {
		version := viper.GetString("version")
		// commit := viper.GetString("commit")
		buildDate := viper.GetString("build_date")

		fmt.Printf("Todo CLI v%s ( built: %s)\n", version, buildDate)
	},
}

func init() {
		rootCmd.AddCommand(versionCmd)



}
