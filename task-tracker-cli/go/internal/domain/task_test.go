package domain

import (
	"errors"
	"testing"
)

func TestNewTask(t *testing.T) {
	task, err := NewTask(1, "Buy tomato")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if task.ID != 1 {
		t.Errorf("expected ID 1, got %d", task.ID)
	}
	if task.Status != StatusTodo {
		t.Errorf("expected status todo, got %s", task.Status)
	}
	if task.CreatedAt == "" || task.UpdatedAt == "" {
		t.Error("timestamp should be set")
	}
}

func TestUpdateDescription(t *testing.T) {
	task, _ := NewTask(1, "Buy tomato")
	err := task.UpdateDescription("Buy 1kg tomato")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if task.Description != "Buy 1kg tomato" {
		t.Errorf("expected updated description")
	}
}

func TestUpdateDescription_EmptyDescription_ShouldFail(t *testing.T) {
	task, _ := NewTask(1, "Buy tomato")
	originalDesc := task.Description

	err := task.UpdateDescription("")
	if err == nil {
		t.Fatalf("expected error for empty description")
	}

	var validationErr *ValidationError
	if !errors.As(err, &validationErr) {
		t.Fatalf("expected ValidationError, got %T: %v", err, err)
	}

	// Description should remain unchanged
	if task.Description != originalDesc {
		t.Errorf("description should remain %q, got %q", originalDesc, task.Description)
	}
}

func TestMarkInProgress(t *testing.T) {
	task, _ := NewTask(1, "Buy tomato")
	err := task.MarkInProgress()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if task.Status != StatusInProgress {
		t.Errorf("expected in-progress, got %s", task.Status)
	}
}

func TestMarkInProgress_FromDone_ShouldFail(t *testing.T) {
	task, _ := NewTask(1, "Buy tomato")
	_ = task.MarkDone()

	err := task.MarkInProgress()
	if err == nil {
		t.Fatalf("expected error when marking done task as in-progress")
	}

	var validationErr *ValidationError
	if !errors.As(err, &validationErr) {
		t.Fatalf("expected ValidationError, got %T: %v", err, err)
	}

	// Status should remain done
	if task.Status != StatusDone {
		t.Errorf("expected status to remain %q, got %q", StatusDone, task.Status)
	}
}

func TestMarkDoneIsIdempotent(t *testing.T) {
	task, _ := NewTask(1, "Buy tomato")
	_ = task.MarkDone()
	err := task.MarkDone() // second call

	if err != nil {
		t.Fatalf("mark-done should be idempotent")
	}
	if task.Status != StatusDone {
		t.Errorf("expected done status, but got %s", task.Status)
	}
}

func TestNewTask_InvalidID_ShouldFail(t *testing.T) {
	tests := []struct {
		name string
		id   int
	}{
		{"zero ID", 0},
		{"negative ID", -1},
		{"large negative ID", -999},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task, err := NewTask(tt.id, "Valid description")
			if err == nil {
				t.Fatalf("expected error for ID %d", tt.id)
			}

			var validationErr *ValidationError
			if !errors.As(err, &validationErr) {
				t.Fatalf("expected ValidationError, got %T: %v", err, err)
			}

			if task != nil {
				t.Errorf("expected nil task, got %+v", task)
			}
		})
	}
}

func TestNewTask_EmptyDescription_ShouldFail(t *testing.T) {
	tests := []struct {
		name        string
		description string
	}{
		{"empty string", ""},
		{"only spaces", "   "},
		{"only tabs", "\t\t"},
		{"only newlines", "\n\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task, err := NewTask(1, tt.description)
			if err == nil {
				t.Fatalf("expected error for description %q", tt.description)
			}

			var validationErr *ValidationError
			if !errors.As(err, &validationErr) {
				t.Fatalf("expected ValidationError, got %T: %v", err, err)
			}

			if task != nil {
				t.Errorf("expected nil task, got %+v", task)
			}
		})
	}
}
