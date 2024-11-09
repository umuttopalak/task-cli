package storage

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/umuttopalak/task-cli/src/domain"
	"github.com/umuttopalak/task-cli/src/repository"
)

var taskNotFound = errors.New("Task Not Found")

type JSONTaskRepository struct {
	FilePath string
	tasks    []*domain.Task
}

func NewJSONTaskRepository(filePath string) repository.TaskRepository {
	return &JSONTaskRepository{
		FilePath: filePath,
		tasks:    []*domain.Task{},
	}
}

func (j *JSONTaskRepository) AddTask(task *domain.Task) error {
	if _, err := j.LoadTasks(); err != nil {
		return err
	}
	task.ID = j.generateNewTaskID()
	j.tasks = append(j.tasks, task)
	return j.Save()
}

func (j *JSONTaskRepository) DeleteByID(id int) error {
	if _, err := j.LoadTasks(); err != nil {
		return err
	}
	for i, task := range j.tasks {
		if task.ID == id {
			j.tasks = append(j.tasks[:i], j.tasks[i+1:]...)
			if err := j.Save(); err != nil {
				return err
			}
			return nil

		}
	}

	return taskNotFound
}

func (j *JSONTaskRepository) FindByID(id int) (*domain.Task, error) {
	if _, err := j.LoadTasks(); err != nil {
		return nil, err
	}

	for _, t := range j.tasks {
		if t.ID == id {
			return t, nil
		}
	}
	return nil, taskNotFound
}

func (j *JSONTaskRepository) LoadTasks() ([]*domain.Task, error) {
	if _, err := os.Stat(j.FilePath); os.IsNotExist(err) {
		j.tasks = []*domain.Task{}
		return j.tasks, nil
	}

	data, err := os.ReadFile(j.FilePath)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &j.tasks); err != nil {
		return nil, err
	}

	return j.tasks, nil
}

func (j *JSONTaskRepository) Save() error {
	bytes, err := json.MarshalIndent(j.tasks, "", "  ")
	if err != nil {
		return err
	}
	file, err := os.OpenFile(j.FilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(bytes)
	return err
}

func (j *JSONTaskRepository) generateNewTaskID() int {
	maxID := 0
	for _, t := range j.tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}

	return maxID + 1
}
