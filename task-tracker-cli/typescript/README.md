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

# Build the project
npm run build
```

### Usage

```bash
# Add a new task
npm start add "Buy groceries"

# Update a task
npm start update 1 "Buy groceries and cook dinner"

# Mark task as in-progress
npm start mark-in-progress 1

# Mark task as done
npm start mark-done 1

# Delete a task
npm start delete 1

# List all tasks
npm start list

# List by status
npm start list done
npm start list todo
npm start list in-progress
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

```bash
$ npm start add "Learn TypeScript"
Task added successfully (ID: 1)

$ npm start add "Build CLI app"
Task added successfully (ID: 2)

$ npm start list
[1] Learn TypeScript (todo)
[2] Build CLI app (todo)

$ npm start mark-in-progress 1
Task marked as in-progress

$ npm start mark-done 1
Task marked as done

$ npm start list
[1] Learn TypeScript (done)
[2] Build CLI app (todo)
```

---

## 🧑‍💻 Development

### Building

```bash
npm run build       # Compile TypeScript to dist/
```

### Development Mode

```bash
npm run dev         # Run with ts-node (no build needed)
```

### Scripts

- `npm run build` - Compile TypeScript
- `npm start` - Run the compiled app
- `npm run dev` - Run with ts-node
- `npm test` - Run tests
- `npm run test:watch` - Run tests in watch mode
- `npm run lint` - Run linter (placeholder)

---

## 📜 License

This project is part of the Backend Challenge learning repository.
