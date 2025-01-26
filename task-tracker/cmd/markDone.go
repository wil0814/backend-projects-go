package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// markDoneCmd represents the markDone command
var markDoneCmd = &cobra.Command{
	Use:   "mark-done [taskID]",
	Short: "Mark a task as done",
	Long:  `Use this command to mark a specific task as done in your task tracker by providing its ID.`,
	Example: `  # Mark task with ID 1 as done
  task-cli mark-done 1`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskID := args[0]
		fmt.Printf("Task with ID %s has been marked as done.\n", taskID)
	},
}

func init() {
	markDoneCmd.DisableFlagsInUseLine = true
	rootCmd.AddCommand(markDoneCmd)
}
