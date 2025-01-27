# Task Tracker CLI

A simple task tracker CLI written in Go.

## Installation

To install the task-tracker CLI, run the following command:

```bash
go install backend-products-go/task-tracker
```

## Usage

To use the task-tracker CLI, run the following command:

```bash
task-tracker
```

This will display the help message for the task-tracker CLI.

## Commands

The task-tracker CLI supports the following commands:

- `add`: Add a new task
- `delete`: Delete a specific task
- `list`: List tasks
- `mark-done`: Mark a task as done
- `mark-in-progress`: Mark a task as in progress
- `update`: Update an existing task

For example, to add a new task, run the following command:

```bash
task-tracker add "Buy groceries"
```

This will add a new task with the description "Buy groceries" to the task tracker.

## Environment Variables

The task-tracker CLI uses the following environment variables:

- `TASKS_FILE_PATH`: The path to the file where the tasks are stored. If not set, the default path is the current directory.

## Contributing

Contributions are welcome! If you find a bug or have a suggestion, please open an issue or submit a pull request.