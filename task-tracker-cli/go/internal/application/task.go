package application

import (
	"taskcli/internal/domain"
	"taskcli/internal/ports"
)

// TaskService holds use-cases. No knowledge of file/JSON
type TaskService struct{ repo ports.TaskRepository }

func NewTaskService(r ports.TaskRepository) *TaskService {
	return &TaskService{repo: r}
}

func nextID(tasks []domain.Task) int {
	max := 0
	for _, t := range tasks {
		if t.ID > max {
			max = t.ID
		}
	}

	return max + 1
}

func (s *TaskService) Add(description string) (*domain.Task, error) {
	tasks, err := s.repo.Load()
	if err != nil {
		return nil, err
	}

	id := nextID(tasks)

	task, err := domain.NewTask(id, description)
	if err != nil {
		return nil, err
	}

	tasks = append(tasks, *task)
	if err := s.repo.Save(tasks); err != nil {
		return nil, err
	}

	return task, nil
}

func (s *TaskService) Update(id int, desc string) error {
	tasks, err := s.repo.Load()
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].ID == id {
			if err := tasks[i].UpdateDescription(desc); err != nil {
				return err
			}

			return s.repo.Save(tasks)
		}
	}

	return &domain.NotFoundError{Msg: "task not found"}
}

func (s *TaskService) Delete(id int) error {
	tasks, err := s.repo.Load()
	if err != nil {
		return err
	}

	idx := -1
	for i := range tasks {
		if tasks[i].ID == id {
			idx = i
			break
		}
	}

	if idx == -1 {
		return &domain.NotFoundError{Msg: "task not found"}
	}

	tasks = append(tasks[:idx], tasks[idx+1:]...)

	return s.repo.Save(tasks)
}
