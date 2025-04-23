# Todo CLI Application

A command-line interface (CLI) for managing todo tasks with a Gin-backed REST API and SQLite database.

## Features

- Add, list, and mark tasks as done via CLI
- REST API for integration with other applications
- SQLite database for persistence
- Production-grade project structure

## Getting Started

### Prerequisites

- Go 1.18 or later
- Make

### Installation

1. Clone the repository
   ```
   git clone https://github.com/yourusername/todo-app.git
   cd todo-app
   ```

2. Initialize the project
   ```
   make init
   ```

3. Build the application
   ```
   make build
   ```

### Usage

```
# Add a new task
./bin/todo add "Buy groceries"

# List all tasks
./bin/todo list

# Mark a task as done
./bin/todo done 1
```

## Development

- Run tests: `make test`
- Generate coverage report: `make coverage`
- Run linter: `make lint`
- Check code formatting: `make fmt`

## License

This project is licensed under the MIT License - see the LICENSE file for details.
EOF
