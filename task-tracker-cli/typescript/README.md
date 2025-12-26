# 📝 Task Tracker CLI - TypeScript Implementation

A command-line task management application built in TypeScript using Clean Architecture principles.

---

## 🎯 Features

- ✅ Add, update, and delete tasks
- 📋 List tasks with filtering by status
- 🏗️ Clean Architecture (Hexagonal/Ports & Adapters)
- 💾 JSON file-based persistence
- ✅ Comprehensive test coverage
- 🔧 Type-safe with TypeScript

---

## 🏗️ Architecture

```
src/
├── domain/          # Business entities and errors
├── ports/           # Interfaces (repository contracts)
├── services/        # Use cases and business rules
├── adapters/        # External implementations (file storage)
├── cli/             # Command-line interface logic
├── infra/           # Infrastructure utilities
└── utils/           # Shared utilities
```

**Design Pattern:** Hexagonal Architecture (Ports & Adapters)  
**Storage:** JSON file (`tasks.json`)

---

## 🚀 Getting Started

### Prerequisites

- Node.js 18+
- npm or yarn

### Installation

```bash
# Clone and navigate to the project
cd task-tracker-cli/typescript

# Install dependencies
npm install
# or
yarn install

# Build and link the CLI globally
npm run link
# or build separately:
npm run build
npm link
```

### Usage

```bash
# Add a new task
task-cli-ts add "Buy groceries"

# Update a task
task-cli-ts update 1 "Buy groceries and cook dinner"

# Mark task as in-progress
task-cli-ts mark-in-progress 1

# Mark task as done
task-cli-ts mark-done 1

# Delete a task
task-cli-ts delete 1

# List all tasks
task-cli-ts list

# List by status
task-cli-ts list done
task-cli-ts list todo
task-cli-ts list in-progress
```

---

## 🧪 Testing

```bash
# Run all tests
npm test

# Run tests in watch mode
npm run test:watch

# Compile tests only
npm run test:compile
```

---

## 📦 Project Structure

```
.
├── src/
│   ├── app.ts               # Application entry point
│   ├── domain/              # Task entity and errors
│   ├── ports/               # Repository interface
│   ├── services/            # Task service (use cases)
│   ├── adapters/
│   │   └── fs/              # File system repository
│   ├── cli/                 # CLI argument parsing and routing
│   ├── infra/               # Infrastructure helpers
│   └── utils/               # Time utilities
├── tests/
│   ├── domain/              # Domain entity tests
│   ├── services/            # Service tests
│   └── mocks/               # Mock implementations
├── package.json
├── tsconfig.json
└── README.md
```

---

## 🧰 Technologies

- **Language:** TypeScript 5.6+
- **Runtime:** Node.js 18+
- **Architecture:** Clean Architecture / Hexagonal
- **Storage:** JSON file
- **Testing:** Node.js native test runner
- **Build:** TypeScript Compiler (tsc)

---

## 📚 Key Concepts Demonstrated

- ✅ Clean Architecture / Hexagonal Architecture
- ✅ Dependency Inversion Principle
- ✅ Repository Pattern
- ✅ Domain-Driven Design (DDD) basics
- ✅ Type safety with TypeScript
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

> **Note:** Make sure you've built and linked the CLI first by running `npm run link` (or `npm run build && npm link`).

```bash
$ task-cli-ts add "Learn TypeScript"
Task added successfully (ID: 1)

$ task-cli-ts add "Build CLI app"
Task added successfully (ID: 2)

$ task-cli-ts list
[1] Learn TypeScript (todo)
[2] Build CLI app (todo)

$ task-cli-ts mark-in-progress 1
Task marked as in-progress

$ task-cli-ts mark-done 1
Task marked as done

$ task-cli-ts list
[1] Learn TypeScript (done)
[2] Build CLI app (todo)
```

---

## 🧑‍💻 Development

### Building

```bash
npm run build       # Compile TypeScript to dist/
```

### Link Globally

```bash
npm run link        # Build and link globally (makes task-cli-ts available)
npm run unlink      # Unlink the global command
```

### Development Mode

```bash
npm run dev -- add "Test task"  # Run with ts-node (no build needed)
# or
npm start add "Test task"       # Run compiled version
```

### Scripts

- `npm run build` - Compile TypeScript
- `npm run link` - Build and link CLI globally
- `npm run unlink` - Remove global link
- `npm start` - Run the compiled app (requires args)
- `npm run dev` - Run with ts-node
- `npm test` - Run tests
- `npm run test:watch` - Run tests in watch mode
- `npm run lint` - Run linter (placeholder)

---

## 📜 License

This project is part of the Backend Engineering Lab learning repository.
