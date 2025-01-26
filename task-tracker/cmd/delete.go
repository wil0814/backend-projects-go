package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
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
		fmt.Printf("Task with ID %s has been deleted.\n", taskID)
	},
}

func init() {
	deleteCmd.DisableFlagsInUseLine = true
	rootCmd.AddCommand(deleteCmd)
}
