package domain

import (
	"fmt"
	"time"
)

type TaskStatus string

const (
	Todo       TaskStatus = "to-do"
	InProgress TaskStatus = "in-progress"
	Done       TaskStatus = "done"
)

type Task struct {
	ID          int
	Description string
	Status      TaskStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTask(description string) *Task {
	return &Task{
		Description: description,
		Status:      Todo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (t *Task) Update(description string) {
	t.Description = description
	t.UpdatedAt = time.Now()
}

func (t *Task) SetStatus(status TaskStatus) {
	t.Status = status
	t.UpdatedAt = time.Now()
}

func (t *Task) String() string {
	return fmt.Sprintf("ID: %d, Description: %s, Status: %s, Created At: %s, Updated At: %s\n",
		t.ID, t.Description, t.Status, t.CreatedAt.Format(time.RFC3339), t.UpdatedAt.Format(time.RFC3339))
}
