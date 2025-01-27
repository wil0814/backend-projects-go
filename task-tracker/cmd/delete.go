package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/wil0814/backend-projects-go/task-tracker/internal/task"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [task ID]",
	Short: "Delete a specific task",
	Long:  `Use this command to delete a specific task from your task tracker by providing its ID.`,
	Example: `  # Delete the task with ID 1
  task-cli delete 1`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskID := args[0]
		taskIDInt64, err := strconv.ParseInt(taskID, 10, 64)
		if err != nil {
			fmt.Printf("Invalid task ID: %s\n", taskID)
			return
		}
		err = task.DeleteTask(taskIDInt64)
		if err != nil {
			fmt.Printf("Failed to deleting task: %s\n", err)
			return
		}
		fmt.Printf("Task with ID %s has been deleted.\n", taskID)
	},
}

func init() {
	deleteCmd.DisableFlagsInUseLine = true
	rootCmd.AddCommand(deleteCmd)
}
