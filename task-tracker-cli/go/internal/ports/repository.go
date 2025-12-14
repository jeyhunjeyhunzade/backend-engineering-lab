package ports

import "taskcli/internal/domain"

// TaskRepository abstract persistance
type TaskRepository interface {
	Load() ([]domain.Task, error)
	Save([]domain.Task) error
}
