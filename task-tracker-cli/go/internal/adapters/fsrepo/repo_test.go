package fsrepo

import (
	"path/filepath"
	"taskcli/internal/domain"
	"testing"
)

func TestRepoSaveAndLoad(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "tasks.json")

	repo, err := New(path)
	if err != nil {
		t.Fatalf("failed to create repo: %v", err)
	}

	tasks := []domain.Task{
		{ID: 1, Description: "Buy tomato", Status: domain.StatusTodo},
	}

	if err := repo.Save(tasks); err != nil {
		t.Fatalf("save failed: %v", err)
	}

	loaded, err := repo.Load()
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}

	if len(loaded) != 1 || loaded[0].ID != 1 {
		t.Fatalf("unexpected loaded tasks: %+v", loaded)
	}
}

func TestRepoEmptyFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "tasks.json")

	repo, err := New(path)
	if err != nil {
		t.Fatalf("failed to create repo: %v", err)
	}

	tasks, err := repo.Load()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tasks) != 0 {
		t.Fatalf("expected empty tasks")
	}
}
