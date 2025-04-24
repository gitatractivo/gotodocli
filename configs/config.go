package configs


type Config struct {
	ServerPort string
	ApiBaseUrl string
	ServerPidFile string
	LogFile string
	Version string
	Commit string
	BuildDate string
}

var config *Config

//here input version, commit, builddate
func InitConfig(version, commit, builddate string) {
	config = &Config{
		ServerPort: "8080",
		ApiBaseUrl: "http://localhost:8080/v1",
		ServerPidFile: "/tmp/todo-server.pid",
		LogFile: "/tmp/todo-server.log",
		Version: version,
		Commit: commit,
		BuildDate: builddate,
	}
}

func GetConfig() *Config {
	if config == nil {
		InitConfig("dev", "dev", "dev")
	}
	return config
}


	