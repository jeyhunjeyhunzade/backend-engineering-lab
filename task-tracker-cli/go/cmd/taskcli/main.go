package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"taskcli/internal/adapters/fsrepo"
	"taskcli/internal/application"
	"taskcli/internal/domain"
)

const (
	ExitOk         = 0
	ExitGeneralErr = 1
	ExitUsage      = 2
	ExitNotFound   = 3
)

func main() {
	os.Exit(run(os.Args))
}

func run(args []string) int {
	if len(args) < 2 {
		printHelp()
		return ExitUsage
	}

	repo, err := fsrepo.New("tasks.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return ExitGeneralErr
	}
	svc := application.NewTaskService(repo)

	switch args[1] {
	case "help", "-h", "--help":
		printHelp()
		return ExitOk
	case "add":
		if len(args) < 3 {
			usage(`add "description"`)
			return ExitUsage
		}
		desc := strings.Join(args[2:], " ")
		t, err := svc.Add(desc)
		if err != nil {
			return handleError(err)
		}
		fmt.Printf("Task added successfully (ID: %d)\n", t.ID)
		return ExitOk
	case "update":
		fmt.Println("(stub) update is not implemented yet")
		return ExitGeneralErr
	case "delete":
		fmt.Println("(stub) delete is not implemented yet")
		return ExitGeneralErr
	case "mark-in-progress":
		fmt.Println("(stub) mark-in-progress is not implemented yet")
		return ExitGeneralErr
	case "mark-done":
		fmt.Println("(stub) mark-done is not implemented yet")
		return ExitGeneralErr
	case "list":
		fmt.Println("(stub) list is not implemented yet")
		return ExitGeneralErr
	default:
		fmt.Fprintf(os.Stderr, "unknown command %s\n", args[1])
		printHelp()
		return ExitUsage
	}
}

func parseId(s string) (int, error) {
	id, err := strconv.Atoi(s)
	if err != nil || id <= 0 {
		return 0, fmt.Errorf("invalid id: %q", s)
	}

	return id, nil
}

func usage(example string) {
	fmt.Fprintf(os.Stderr, "usage: task-tracker-cli-go %s\n", example)
}

func handleError(err error) int {
	// Map domain errors to exit codes
	var nf *domain.NotFoundError
	var ve *domain.ValidationError

	switch {
	case errors.As(err, &nf):
		fmt.Fprintln(os.Stderr, err.Error())
		return ExitNotFound
	case errors.As(err, &ve):
		fmt.Fprintln(os.Stderr, err.Error())
		return ExitUsage
	default:
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return ExitGeneralErr
	}
}

const help = `Task Tracker CLI (Go)

Usage:

  task-cli add "description"
  task-cli update <id> "new description"
  task-cli delete <id>
  task-cli mark-in-progress <id>
  task-cli mark-done <id>
  task-cli list [todo|in-progress|done]
`

func printHelp() {
	fmt.Println(help)
}
