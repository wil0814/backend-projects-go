package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// markInProgressCmd represents the markInProgress command
var markInProgressCmd = &cobra.Command{
	Use:   "mark-in-progress [taskID]",
	Short: "Mark a task as in progress",
	Long:  `Use this command to mark a specific task as "in progress" in your task tracker by providing its ID.`,
	Example: `  # Mark task with ID 1 as in progress
  task-cli mark-in-progress 1`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskID := args[0]
		fmt.Printf("Task with ID %s has been marked as in progress.\n", taskID)
	},
}

func init() {
	markInProgressCmd.DisableFlagsInUseLine = true
	rootCmd.AddCommand(markInProgressCmd)
}
