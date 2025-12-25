.PHONY: build test test-verbose test-coverage clean run install fmt vet help

# Binary name
BINARY_NAME=task-tracker-cli-go
MAIN_PATH=./cmd/taskcli

# Build the application
build:
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BINARY_NAME) $(MAIN_PATH)
	@echo "Build complete!"

# Run tests
test:
	@echo "Running tests..."
	@go test ./...

# Run tests with verbose output
test-verbose:
	@echo "Running tests (verbose)..."
	@go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	@go test -cover ./...
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -f $(BINARY_NAME)
	@rm -f coverage.out coverage.html
	@echo "Clean complete!"

# Install dependencies
install:
	@echo "Installing dependencies..."
	@go mod download
	@go mod tidy

# Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# Run go vet
vet:
	@echo "Running go vet..."
	@go vet ./...

# Build and run
run: build
	@./$(BINARY_NAME) $(ARGS)

# Default target
help:
	@echo "Available targets:"
	@echo "  make build         - Build the application"
	@echo "  make test          - Run tests"
	@echo "  make test-verbose  - Run tests with verbose output"
	@echo "  make test-coverage - Run tests with coverage report"
	@echo "  make clean         - Remove build artifacts"
	@echo "  make install       - Install dependencies"
	@echo "  make fmt           - Format code"
	@echo "  make vet           - Run go vet"
	@echo "  make run ARGS='...'- Build and run (pass args with ARGS='...')"
	@echo "  make help          - Show this help message"
