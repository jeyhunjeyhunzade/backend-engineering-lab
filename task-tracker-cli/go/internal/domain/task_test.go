package domain

import "testing"

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
