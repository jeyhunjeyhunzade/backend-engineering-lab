# 📝 Task Tracker CLI - Go Implementation

A command-line task management application built in Go using Clean Architecture principles.

---

## 🎯 Features

- ✅ Add, update, and delete tasks
- 📋 List tasks with filtering by status
- 🏗️ Clean Architecture (Hexagonal/Ports & Adapters)
- 💾 JSON file-based persistence
- ✅ Comprehensive test coverage

---

## 🏗️ Architecture

```
internal/
├── domain/          # Business entities and logic
├── ports/           # Interfaces (repository contracts)
├── application/     # Use cases and business rules
└── adapters/        # External implementations (file storage)
```

**Design Pattern:** Hexagonal Architecture (Ports & Adapters)  
**Storage:** JSON file (`tasks.json`)

---

## 🚀 Getting Started

### Prerequisites

- Go 1.21+

### Installation

```bash
# Clone and navigate to the project
cd task-tracker-cli/go

# Install dependencies
go mod download

# Build the binary
go build -o task-tracker-cli-go ./cmd/taskcli
```

### Usage

```bash
# Add a new task
./task-tracker-cli-go add "Buy groceries"

# Update a task
./task-tracker-cli-go update 1 "Buy groceries and cook dinner"

# Mark task as in-progress
./task-tracker-cli-go mark-in-progress 1

# Mark task as done
./task-tracker-cli-go mark-done 1

# Delete a task
./task-tracker-cli-go delete 1

# List all tasks
./task-tracker-cli-go list

# List by status
./task-tracker-cli-go list done
./task-tracker-cli-go list todo
./task-tracker-cli-go list in-progress
```

---

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...
```

---

## 📦 Project Structure

```
.
├── cmd/
│   └── taskcli/           # Application entry point
├── internal/
│   ├── domain/            # Task entity and business logic
│   ├── ports/             # Repository interface
│   ├── application/       # Task service (use cases)
│   └── adapters/
│       └── fsrepo/        # File system repository implementation
├── go.mod
├── tasks.json             # Data storage (generated)
└── README.md
```

---

## 🧰 Technologies

- **Language:** Go 1.21+
- **Architecture:** Clean Architecture / Hexagonal
- **Storage:** JSON file
- **Testing:** Go standard testing package

---

## 📚 Key Concepts Demonstrated

- ✅ Clean Architecture / Hexagonal Architecture
- ✅ Dependency Inversion Principle
- ✅ Repository Pattern
- ✅ Domain-Driven Design (DDD) basics
- ✅ Unit Testing
- ✅ CLI application development

---

## 🔄 Task States

Tasks can be in one of three states:

- **todo** - Initial state
- **in-progress** - Task is being worked on
- **done** - Task is completed

---

## 📝 Example

```bash
$ ./task-tracker-cli-go add "Learn Go"
Task added successfully (ID: 1)

$ ./task-tracker-cli-go add "Build CLI app"
Task added successfully (ID: 2)

$ ./task-tracker-cli-go list
[1] Learn Go (todo)
[2] Build CLI app (todo)

$ ./task-tracker-cli-go mark-in-progress 1
Task marked as in-progress

$ ./task-tracker-cli-go mark-done 1
Task marked as done

$ ./task-tracker-cli-go list
[1] Learn Go (done)
[2] Build CLI app (todo)
```

---

## 🧑‍💻 Development

### Building

```bash
make build    # or: go build -o task-tracker-cli-go ./cmd/taskcli
```

### Running Tests

```bash
make test     # or: go test ./...
```

---

## 📜 License

This project is part of the Backend Engineering Lab learning repository.
