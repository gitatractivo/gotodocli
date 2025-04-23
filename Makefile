.PHONY: all build clean test lint run migrate docker-build help init vet fmt check

# go parameters
BINARY_NAME = go-api
MAIN_PACKAGE=./cmd/todo
BINARY_PATH= ./bin/$(BINARY_NAME)
GO_FILES=$(shell find . -name "*.go" -type f)
GO_PACKAGES=$(shell go list ./... | grep -v /vendor/)

#BUILD PARAMETERS
VERSION=$(shell git describe --tags --always --dirty)
COMMIT=$(shell git rev-parse HEAD)
BUILD_DATE=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.Commit=$(COMMIT) -X main.BuildDate=$(BUILD_DATE)"

# docker parameters
IMAGE_NAME=go-api
IMAGE_TAG=$(VERSION)

# Test parameters
COVERAGE_DIR=coverage
COVERAGE_PROFILE=$(COVERAGE_DIR)/coverage.out
COVERAGE_HTML=$(COVERAGE_DIR)/coverage.html

all: check test build

help:
	@echo "Available commands:"
	@echo "  make init       - Initialize the project and install dependencies"
	@echo "  make build      - Build the binary"
	@echo "  make run        - Run the application"
	@echo "  make test       - Run tests"
	@echo "  make coverage   - Run tests with coverage report"
	@echo "  make lint       - Run linter"
	@echo "  make vet        - Run go vet"
	@echo "  make fmt        - Run go fmt"
	@echo "  make check      - Run all checks (fmt, vet, lint)"
	@echo "  make clean      - Remove build artifacts"
	@echo "  make migrate    - Run database migrations"
	@echo "  make docker-build - Build Docker image"


init: 
	@echo "Initializing project..."
	go mod tidy
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/tools/cmd/goimports@latest
	mkdir -p bin
	mkdir -p $(COVERAGE_DIR)
	@echo "Project initialized successfully"

build: 
	@echo "Building $(BINARY_NAME)..."
	go build $(LDFLAGS) -o $(BINARY_PATH) $(MAIN_PACKAGE)
	@echo "Build completed successfully"

run: Build
	@echo "Running $(BINARY_NAME)..."
	$(BINARY_PATH)

test:
	@echo "Running tests..."
	go test -v -cover -coverprofile=$(COVERAGE_PROFILE) $(GO_PACKAGES)
	@echo "Tests completed successfully"

coverage:
	@echo "Running tests with coverage..."
	mkdir -p ${COVERAGE_DIR}
	go test -coverprofile=${COVERAGE_PROFILE} ${GO_PACKAGES}
	go tool cover -html=${COVERAGE_PROFILE} -o ${COVERAGE_HTML}
	@echo "Coverage report generated at ${COVERAGE_HTML}"


lint:
	@echo "Running lint..."
	golangci-lint run --fix ./...
	@echo "Lint completed successfully"


vet:
	@echo "Running vet..."
	go vet $(GO_PACKAGES)
	@echo "Vet completed successfully"


fmt:
	@echo "Running fmt..."
	go fmt $(GO_PACKAGES)
	@echo "Fmt completed successfully"


check: fmt vet lint
	@echo "All checks completed successfully"


clean:
	@echo "Cleaning up..."
	rm -rf bin $(COVERAGE_DIR)
	@echo "Cleanup completed successfully"


migrate:
	@echo "Running migrations..."
	go run $(MAIN_PACKAGE)/migrations/migrate.go up
	@echo "Migrations completed successfully"


docker-build:
	@echo "Building Docker image..."
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) .
	@echo "Docker image built successfully"



	