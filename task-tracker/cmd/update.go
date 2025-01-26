package cmd

import (
	"fmt"

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

		fmt.Printf("Task with ID %s has been updated to: %s\n", taskID, taskName)
	},
}

func init() {
	updateCmd.DisableFlagsInUseLine = true
	rootCmd.AddCommand(updateCmd)
}
