package fsrepo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"taskcli/internal/domain"
	"taskcli/internal/ports"
)

// Repo implements TaskRepository against a JSON file on disk.
// Standard library only. Use os, io, encoding/json, os.CreateTemp + os.Rename for atomic-ish writes.

var _ ports.TaskRepository = (*Repo)(nil)

type Repo struct{ path string }

// New creates a file-backed repo at filename
// filename is interpreted relative to the current working directory if not absolute
func New(filename string) (*Repo, error) {
	if filename == "" {
		filename = "task.json"
	}

	abs, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}

	if err := ensureFile(abs); err != nil {
		return nil, err
	}

	return &Repo{path: abs}, nil
}

func (r *Repo) Load() ([]domain.Task, error) {
	f, err := os.Open(r.path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	// Allow empty file to represent no tasks
	if len(b) == 0 {
		return []domain.Task{}, nil
	}

	var tasks []domain.Task
	if err := json.Unmarshal(b, &tasks); err != nil {
		return nil, fmt.Errorf("corrupted JSON is in %s: %w", r.path, err)
	}

	if tasks == nil {
		return []domain.Task{}, nil
	}

	return tasks, nil
}

func (r *Repo) Save(tasks []domain.Task) error {
	dir := filepath.Dir(r.path)

	tmp, err := os.CreateTemp(dir, "task-*.json")
	if err != nil {
		return err
	}
	tmpPath := tmp.Name()

	// if we fail anywhere, clean up temp file
	cleanup := func(closeErr error) error {
		_ = tmp.Close()
		_ = os.Remove(tmpPath)
		return closeErr
	}

	enc := json.NewEncoder(tmp)
	enc.SetIndent("", " ")

	if err := enc.Encode(tasks); err != nil {
		return cleanup(err)
	}

	// Best-effort flush to disk (especially useful on crashes)
	if err := tmp.Sync(); err != nil {
		return cleanup(err)
	}
	if err := tmp.Close(); err != nil {
		_ = os.Remove(tmpPath)
		return err
	}

	// on Windows, os.Rename fails if destination exists
	// Best-effort remove the old file first
	_ = os.Remove(r.path)

	if err := os.Rename(tmpPath, r.path); err != nil {
		return err
	}

	return nil

}

func ensureFile(path string) error {
	info, err := os.Stat(path)
	if err == nil {
		if info.IsDir() {
			return errors.New("tasks path is directory, expected a file")
		}
		return nil
	}

	if !os.IsNotExist(err) {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}

	return os.WriteFile(path, []byte(""), 0o644)
}
