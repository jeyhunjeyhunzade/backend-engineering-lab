package domain

import (
	"fmt"
	"strings"
	"time"
)

type TaskStatus string

const (
	StatusTodo       TaskStatus = "todo"
	StatusInProgress TaskStatus = "in-progress"
	StatusDone       TaskStatus = "done"
)

// Task is aggregate root
type Task struct {
	ID          int        `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   string     `json:"createdAt"`
	UpdatedAt   string     `json:"updatedAt"`
}

// Domain errors
// Use this types so the CLI can map to exit codes

// NotFound Error
type NotFoundError struct{ Msg string }

func (e *NotFoundError) Error() string { return e.Msg }

// Validation Error
type ValidationError struct{ Msg string }

func (e *ValidationError) Error() string { return e.Msg }

// ParseStatus parses a string into a TaskStatus
func ParseStatus(s string) (TaskStatus, error) {
	switch s {
	case string(StatusTodo):
		return StatusTodo, nil
	case string(StatusInProgress):
		return StatusInProgress, nil
	case string(StatusDone):
		return StatusDone, nil
	default:
		return "", &ValidationError{Msg: fmt.Sprintf("invalid status %q (expected: todo|in-progress|done)", s)}
	}
}

// NowIso returns current UTC time in ISO-8601 / RFC3339 format.
func NowIso() string {
	return time.Now().UTC().Format(time.RFC3339)
}

// NewTask is Task aggregate constructor.
// It enforces initial invariants
func NewTask(id int, desc string) (*Task, error) {
	if id <= 0 {
		return nil, &ValidationError{Msg: "id must be positive!"}
	}
	if strings.TrimSpace(desc) == "" {
		return nil, &ValidationError{Msg: "description cannot be empty"}
	}

	now := NowIso()

	return &Task{
		ID:          id,
		Description: desc,
		Status:      StatusTodo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

// UpdateDescription changes the task description and refreshes its updatedAt
func (t *Task) UpdateDescription(desc string) error {
	if strings.TrimSpace(desc) == "" {
		return &ValidationError{Msg: "description cannot be empty"}
	}

	t.Description = desc
	t.UpdatedAt = NowIso()
	return nil
}

// MarkInProgress transition task in progress if allowed
func (t *Task) MarkInProgress() error {
	if t.Status == StatusDone {
		return &ValidationError{Msg: "cannot mark done task as in progress"}
	}

	t.Status = StatusInProgress
	t.UpdatedAt = NowIso()
	return nil
}

// MarkDone transition the task to done
func (t *Task) MarkDone() error {
	if t.Status == StatusDone {
		return nil // idempotent
	}

	t.Status = StatusDone
	t.UpdatedAt = NowIso()
	return nil
}
