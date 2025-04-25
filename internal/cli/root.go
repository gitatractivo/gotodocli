package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile     string
	showVersion bool = false

	rootCmd = &cobra.Command{
		Use:   "gotodo",
		Short: "A simple CLI todo application",
		Long:  `A CLI application for managing your todo list with local server capabilities.`,
		Run: func(cmd *cobra.Command, args []string) {
			if showVersion {
				printVersion()
				return
			}
			// If no subcommand is provided, show help
			cmd.Help()
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "Print version information")

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gotodo.yaml)")
	rootCmd.PersistentFlags().String("api-url", "", "API base URL")
	rootCmd.PersistentFlags().String("server-port", "", "Server port")

	// Bind flags to viper
	viper.BindPFlag("api_base_url", rootCmd.PersistentFlags().Lookup("api-url"))
	viper.BindPFlag("server_port", rootCmd.PersistentFlags().Lookup("server-port"))

	// If config file is specified via flag, use it
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

}

func printVersion() {
	version := viper.GetString("version")
	buildDate := viper.GetString("build_date")

	fmt.Printf("Todo CLI v%s ( built: %s)\n", version, buildDate)
}
