package application

import (
	"taskcli/internal/domain"
	"testing"
)

// in-memory repo for tests
type memRepo struct {
	tasks []domain.Task
}

func (m *memRepo) Load() ([]domain.Task, error) {
	return append([]domain.Task(nil), m.tasks...), nil
}

func (m *memRepo) Save(t []domain.Task) error {
	m.tasks = append([]domain.Task(nil), t...)
	return nil
}

func TestAddTask(t *testing.T) {
	repo := &memRepo{}
	svc := NewTaskService(repo)

	task, err := svc.Add("Buy tomato")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if task.ID != 1 {
		t.Errorf("expected ID 1, got %d", task.ID)
	}
	if len(repo.tasks) != 1 {
		t.Errorf("expected 1 task saved")
	}
}

func TestUpdateTask(t *testing.T) {
	repo := &memRepo{
		tasks: []domain.Task{
			{ID: 1, Description: "Buy tomato", Status: domain.StatusTodo},
		},
	}
	svc := NewTaskService(repo)
	err := svc.Update(1, "Buy 2kg tomato")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if repo.tasks[0].Description != "Buy 2kg tomato" {
		t.Errorf("description not updated")
	}
}

func TestUpdateNotFound(t *testing.T) {
	repo := &memRepo{}
	svc := NewTaskService(repo)

	err := svc.Update(42, "Am I exist?")
	if err == nil {
		t.Fatalf("expected error")
	}

	if _, ok := err.(*domain.NotFoundError); !ok {
		t.Fatalf("expected NotFoundError")
	}
}

func TestMarkInProgress(t *testing.T) {
	repo := &memRepo{
		tasks: []domain.Task{
			{ID: 1, Description: "A", Status: domain.StatusTodo},
		},
	}
	svc := NewTaskService(repo)

	if err := svc.MarkInProgress(1); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if repo.tasks[0].Status != domain.StatusInProgress {
		t.Errorf("expected status %q, got %q", domain.StatusInProgress, repo.tasks[0].Status)
	}
}

func TestMarkDone(t *testing.T) {
	repo := &memRepo{
		tasks: []domain.Task{
			{ID: 1, Description: "A", Status: domain.StatusTodo},
		},
	}
	svc := NewTaskService(repo)

	if err := svc.MarkDone(1); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if repo.tasks[0].Status != domain.StatusDone {
		t.Errorf("expected status %q, got %q", domain.StatusDone, repo.tasks[0].Status)
	}
}

func TestDelete(t *testing.T) {
	repo := &memRepo{
		tasks: []domain.Task{
			{ID: 1, Description: "A", Status: domain.StatusTodo},
			{ID: 2, Description: "B", Status: domain.StatusTodo},
		},
	}
	svc := NewTaskService(repo)

	if err := svc.Delete(1); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(repo.tasks) != 1 {
		t.Fatalf("expected 1 task remaining, got %d", len(repo.tasks))
	}
	if repo.tasks[0].ID != 2 {
		t.Errorf("expected remaining task ID 2, got %d", repo.tasks[0].ID)
	}
}

func TestDeleteNotFound(t *testing.T) {
	repo := &memRepo{
		tasks: []domain.Task{{ID: 1, Description: "A", Status: domain.StatusTodo}},
	}
	svc := NewTaskService(repo)

	err := svc.Delete(999)
	if err == nil {
		t.Fatalf("expected error")
	}
	if _, ok := err.(*domain.NotFoundError); !ok {
		t.Fatalf("expected NotFoundError, got %T", err)
	}
}

func TestListAllSortedByID(t *testing.T) {
	repo := &memRepo{
		tasks: []domain.Task{
			{ID: 3, Description: "C", Status: domain.StatusTodo},
			{ID: 1, Description: "A", Status: domain.StatusDone},
			{ID: 2, Description: "B", Status: domain.StatusInProgress},
		},
	}
	svc := NewTaskService(repo)

	got, err := svc.List(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(got) != 3 {
		t.Fatalf("expected 3 tasks, got %d", len(got))
	}
	if got[0].ID != 1 || got[1].ID != 2 || got[2].ID != 3 {
		t.Errorf("expected IDs [1,2,3], got [%d,%d,%d]", got[0].ID, got[1].ID, got[2].ID)
	}
}

func TestListFiltered(t *testing.T) {
	repo := &memRepo{
		tasks: []domain.Task{
			{ID: 1, Description: "A", Status: domain.StatusDone},
			{ID: 2, Description: "B", Status: domain.StatusTodo},
			{ID: 3, Description: "C", Status: domain.StatusDone},
		},
	}
	svc := NewTaskService(repo)

	filter := domain.StatusDone
	got, err := svc.List(&filter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(got) != 2 {
		t.Fatalf("expected 2 done tasks, got %d", len(got))
	}
	if got[0].Status != domain.StatusDone || got[1].Status != domain.StatusDone {
		t.Errorf("expected all statuses %q", domain.StatusDone)
	}
	if got[0].ID != 1 || got[1].ID != 3 {
		t.Errorf("expected done IDs [1,3], got [%d,%d]", got[0].ID, got[1].ID)
	}
}
