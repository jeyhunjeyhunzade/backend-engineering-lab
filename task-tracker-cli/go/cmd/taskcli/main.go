package main

import (
	"fmt"
	"os"
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

	switch args[1] {
	case "help", "-h", "--help":
		printHelp()
		return ExitGeneralErr
	case "add":
		fmt.Println("(stub) add is not implemented yet")
		return ExitGeneralErr
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
