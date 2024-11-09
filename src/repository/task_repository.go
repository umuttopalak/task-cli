package repository

import "github.com/umuttopalak/task-cli/src/domain"

type TaskRepository interface {
	LoadTasks() ([]*domain.Task, error)
	Save() error
	AddTask(task *domain.Task) error
	FindByID(id int) (*domain.Task, error)
	DeleteByID(id int) error
}
