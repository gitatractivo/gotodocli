package cli

import (
	"errors"
	"fmt"
	
	"github.com/gitatractivo/gotodocli/internal/cli/commands"
	"github.com/gitatractivo/gotodocli/internal/cli/commands/server"
)

// App represents the CLI application
type App struct {
}



// NewApp creates a new CLI application also take version, commit and build date as parameters
func NewApp() *App {
	return &App{}
}

func (a *App) Run(args []string) error {
	if len(args) < 2 {
		fmt.Println("Expected 'add', 'list', 'done', 'server' subcommands")
		return errors.New("no command provided. insufficient arguments")
	}

	switch args[1] {
	case "add":
		return commands.HandleAddTask(args[2:])
	case "list":
		return commands.HandleListTasks()
	case "done":
		return commands.HandleMarkTaskAsDone(args[2:])
	case "server":
		if len(args) < 3 {
			return errors.New("expected 'start' or 'stop' after 'server'")
		}
		switch args[2] {
		case "start":
			return server.HandleStartServer()
		case "stop":
			return server.HandleStopServer()
		case "status":
			return server.HandleServerStatus()
		default:
			return errors.New("invalid server command. expected 'start', 'stop', or 'status'")
		}
	case "--help", "-h":
		return commands.PrintHelp()
	case "--version", "-v":
		return commands.PrintVersion()
	default:
		return errors.New("invalid command. expected 'add', 'list', 'done', or 'server'")
	}

}
