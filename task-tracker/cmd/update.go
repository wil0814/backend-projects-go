package cmd

import (
	"backend-products-go/task-tracker/internal/task"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update [taskID] [task name]",
	Short: "Update an existing task",
	Long:  `Use this command to update the details of an existing task in your task tracker.`,
	Example: `  # Update task with ID 1 to "New Task Name"
  task-cli update 1 "New Task Name"`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		taskID := args[0]
		taskName := args[1]

		taskIDInt64, err := strconv.ParseInt(taskID, 10, 64)
		if err != nil {
			fmt.Printf("Invalid task ID: %s\n", taskID)
			return
		}
		err = task.UpdateTask(taskIDInt64, taskName)
		if err != nil {
			fmt.Printf("Failed to update task: %v\n", err)
			return
		}
		fmt.Printf("Task with ID %s has been updated.\n", taskID)
	},
}

func init() {
	updateCmd.DisableFlagsInUseLine = true
	rootCmd.AddCommand(updateCmd)
}
