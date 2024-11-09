package service

import (
	"github.com/umuttopalak/task-cli/src/domain"
	"github.com/umuttopalak/task-cli/src/repository"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) domain.TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (s *TaskService) AddTask(description string) (*domain.Task, error) {
	task := domain.NewTask(description)
	err := s.repo.AddTask(task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *TaskService) UpdateTask(id int, description string) error {
	task, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	task.Update(description)
	return s.repo.Save()
}

func (s *TaskService) SetStatus(id int, status domain.TaskStatus) error {
	task, err := s.repo.FindByID(id)
	if err != nil {
		return nil
	}
	task.SetStatus(domain.InProgress)
	return s.repo.Save()
}

func (s *TaskService) ListAllTasks() ([]*domain.Task, error) {
	return s.repo.LoadTasks()
}

func (s *TaskService) ListTasksByStatus(status domain.TaskStatus) ([]*domain.Task, error) {
	tasks, err := s.repo.LoadTasks()
	if err != nil {
		return nil, err
	}
	var filteredTasks []*domain.Task
	for _, t := range tasks {
		if t.Status == status {
			filteredTasks = append(filteredTasks, t)
		}
	}

	return filteredTasks, nil
}

func (s *TaskService) DeleteTask(id int) error {
	if err := s.repo.DeleteByID(id); err != nil {
		return err
	}
	return nil
}
