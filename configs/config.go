// configs/config.go
package configs

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort    string
	ApiBaseUrl    string
	ServerPidFile string
	LogFile       string
	Version       string
	Commit        string
	BuildDate     string
}

var config *Config

// InitConfig now uses Viper to load configuration
func InitConfig(version, commit, builddate string) {
	// Set defaults in Viper
	viper.SetDefault("server_port", "8080")
	viper.SetDefault("api_base_url", "http://localhost:8080/v1")
	viper.SetDefault("server_pid_file", "/tmp/todo-server.pid")
	viper.SetDefault("log_file", "/tmp/todo-server.log")
	viper.SetDefault("version", version)
	viper.SetDefault("commit", commit)
	viper.SetDefault("build_date", builddate)

	home, err := os.UserHomeDir()
	if err == nil {
		// Search config in home directory with name ".gotodo" (without extension)
		viper.AddConfigPath(home)
		viper.SetConfigName(".gotodo")
	}

	// Also check the current working directory for a config file
	viper.AddConfigPath(".")

	// Read environment variables that match
	viper.AutomaticEnv()

	
	// Create config from Viper values
	config = &Config{
		ServerPort:    viper.GetString("server_port"),
		ApiBaseUrl:    viper.GetString("api_base_url"),
		ServerPidFile: viper.GetString("server_pid_file"),
		LogFile:       viper.GetString("log_file"),
		Version:       viper.GetString("version"),
		Commit:        viper.GetString("commit"),
		BuildDate:     viper.GetString("build_date"),
	}
}

func GetConfig() *Config {
	if config == nil {
		InitConfig("dev", "dev", "dev")
	}
	return config
}