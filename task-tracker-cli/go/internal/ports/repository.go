package ports

import "taskcli/internal/domain"

// TaskRepository abstract persistence
type TaskRepository interface {
	Load() ([]domain.Task, error)
	Save([]domain.Task) error
}
