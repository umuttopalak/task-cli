package cli

import (
	"fmt"
	"strconv"

	"github.com/umuttopalak/task-cli/src/domain"
)

type CLIHandler struct {
	taskService domain.TaskService
}

func NewCLIHandler(taskService domain.TaskService) *CLIHandler {
	return &CLIHandler{
		taskService: taskService,
	}
}

func (h *CLIHandler) Run(args []string) {
	if len(args) < 2 {
		h.printUsage()
		return
	}

	command := args[1]
	switch command {
	case "add":
		h.handleAdd(args)
	case "update":
		h.handleUpdate(args)
	case "delete":
		h.handleDelete(args)
	case "mark-in-progress":
		h.handleMarkInProgress(args)
	case "mark-done":
		h.handleMarkDone(args)
	case "list":
		h.handleList(args)
	default:
		fmt.Println("Unknown command:", command)
		h.printUsage()
	}
}

func (h *CLIHandler) handleAdd(args []string) {
	if len(args) < 3 {
		fmt.Println("Usage: task-cli add <description>")
		return
	}
	description := args[2]
	task, err := h.taskService.AddTask(description)
	if err != nil {
		fmt.Println("Error adding task:", err)
	} else {
		fmt.Printf("Task added successfully (ID: %d)\n", task.ID)
	}
}

func (h *CLIHandler) handleUpdate(args []string) {
	if len(args) < 4 {
		fmt.Println("Usage: tasks-cli update <id> <description>")
		return
	}
	id, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Invalid task ID")
		return
	}
	description := args[3]
	err = h.taskService.UpdateTask(id, description)
	if err != nil {
		fmt.Println("Error updating task:", err)
	} else {
		fmt.Println("Task updated successfully")
	}
}

func (h *CLIHandler) handleDelete(args []string) {
	if len(args) < 3 {
		fmt.Println("Usage: task-cli delete <id>")
		return
	}

	id, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Invalid task ID")
		return
	}

	if err := h.taskService.DeleteTask(id); err != nil {
		fmt.Println("Error deleting task:", err)
	} else {
		fmt.Println("Task deleted successfully")
	}
}

func (h *CLIHandler) handleMarkInProgress(args []string) {
	if len(args) < 3 {
		fmt.Println("Usage: task-cli mark-in-progress <id>")
		return
	}
	id, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Invalid task ID")
		return
	}
	if err := h.taskService.SetStatus(id, domain.InProgress); err != nil {
		fmt.Println("Error marking task as in progress", err)
	} else {
		fmt.Println("Task market as in progress")
	}
}

func (h *CLIHandler) handleMarkDone(args []string) {
	if len(args) < 3 {
		fmt.Println("Usage: task-cli mark-done <id>")
		return
	}
	id, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Invalid task ID")
		return
	}
	err = h.taskService.SetStatus(id, domain.Done)
	if err != nil {
		fmt.Println("Error marking task as done:", err)
	} else {
		fmt.Println("Task marked as done")
	}
}

func (h *CLIHandler) handleList(args []string) {
	var tasks []*domain.Task
	var err error
	if len(args) == 2 {
		tasks, err = h.taskService.ListAllTasks()
	} else {
		statusStr := args[2]
		var status domain.TaskStatus
		switch statusStr {
		case "todo":
			status = domain.Todo
		case "in-progress":
			status = domain.InProgress
		case "done":
			status = domain.Done
		default:
			fmt.Println("Unknown task status:", statusStr)
			h.printUsage()
			return
		}
		tasks, err = h.taskService.ListTasksByStatus(status)
	}

	if err != nil {
		fmt.Println("Error listing tasks:", err)
		return
	}
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
	} else {
		for _, t := range tasks {
			fmt.Println(t)
		}
	}
}

func (h *CLIHandler) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\rtask-cli add <description>")
	fmt.Println("\rtask-cli update <id> <description>")
	fmt.Println("\rtask-cli delete <id>")
	fmt.Println("\rtask-cli mark-in-progress <id>")
	fmt.Println("\rtask-cli mark-done <id>")
	fmt.Println("\rtask-cli list [todo|in-progress|done]")
}
