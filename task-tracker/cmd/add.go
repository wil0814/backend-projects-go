package cmd

import (
	"fmt"

	"github.com/wil0814/backend-projects-go/task-tracker/internal/task"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [task name]",
	Short: "Add a new task",
	Long:  `Use this command to add a new task to your task tracker.`,
	Example: `  # Add a task named "Buy groceries"
  task-cli add "Buy groceries"`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskName := args[0]
		id, err := task.AddTask(taskName)
		if err != nil {
			fmt.Printf("Failed to add task: %v\n", err)
			return
		}
		fmt.Printf("Task added successfully (ID: %d)\n", id)
	},
}

func init() {
	addCmd.DisableFlagsInUseLine = true
	rootCmd.AddCommand(addCmd)
}
