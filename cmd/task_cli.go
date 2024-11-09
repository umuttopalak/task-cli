package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/umuttopalak/task-cli/src/cli"
	"github.com/umuttopalak/task-cli/src/service"
	"github.com/umuttopalak/task-cli/src/storage"
)

func Execute() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error: Unable to get the home directory")
		os.Exit(1)
	}

	tasksFilePath := filepath.Join(homeDir, ".task-tracker.json")

	repo := storage.NewJSONTaskRepository(tasksFilePath)

	taskService := service.NewTaskService(repo)

	cliHandler := cli.NewCLIHandler(taskService)

	cliHandler.Run(os.Args)
}
