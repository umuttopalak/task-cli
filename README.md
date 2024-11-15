# Task Tracker CLI

> https://roadmap.sh/projects/task-tracker <br>
> This project is part of [roadmap.sh](https://roadmap.sh) and helps you build a command-line tool for task tracking and management. By working on this project, you will learn about CLI development, data storage, and command handling in Go. Completing this project will enhance your understanding of building real-world tools with Go.

## Features

- **Add a Task**: Create a new task with a description and unique ID.
- **Update a Task**: Edit the description of an existing task.
- **Delete a Task**: Remove a task by its ID.
- **Change Status**: Mark tasks as `todo`, `in-progress`, or `done`.
- **View Tasks**: List tasks based on their status (all, done, in-progress, todo).

## Requirements

- **Command-line Interface**: Must accept commands from the terminal.
- **JSON Storage**: Saves tasks in a `task-tracker.json` file.
- **No External Libraries**: Should use only standard libraries.
  
## Task Properties

Each task is stored with these properties:
- `id`: Unique identifier.
- `description`: A short text description.
- `status`: Task status (`todo`, `in-progress`, `done`).
- `createdAt`: Timestamp for creation.
- `updatedAt`: Timestamp for the latest update.

## Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/umuttopalak/task-cli.git
   cd task-cli
   ```

2. Build the application:
   ```bash
   go build -o task-cli
   ```

3. Move `task-cli` to a directory in your PATH (e.g., `/usr/local/bin`):
   ```bash
   sudo mv task-cli /usr/local/bin/
   ```

## Usage

1. **Add a Task**
   ```bash
   task-cli add "Description of the task"
   ```

2. **Update a Task**
   ```bash
   task-cli update <task_id> "New description"
   ```

3. **Delete a Task**
   ```bash
   task-cli delete <task_id>
   ```

4. **Mark as In-Progress**
   ```bash
   task-cli mark-in-progress <task_id>
   ```

5. **Mark as Done**
   ```bash
   task-cli mark-done <task_id>
   ```

6. **List Tasks**
   - All tasks:
     ```bash
     task-cli list
     ```
   - Filtered by status:
     ```bash
     task-cli list todo
     task-cli list in-progress
     task-cli list done
     ```

## Development

- **Directory Structure**: Organize code by `domain`, `repository`, and `cli` layers.
- **Testing**: Test each function independently to verify expected results.

## License

This project is licensed under the MIT License.
