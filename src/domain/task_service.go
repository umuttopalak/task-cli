package domain

type TaskService interface {
	AddTask(description string) (*Task, error)
	UpdateTask(id int, description string) error
	DeleteTask(id int) error
	SetStatus(id int, status TaskStatus) error
	ListAllTasks() ([]*Task, error)
	ListTasksByStatus(status TaskStatus) ([]*Task, error)
}
