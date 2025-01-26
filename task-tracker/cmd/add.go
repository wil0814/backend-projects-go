package cmd

import (
	"fmt"

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
		fmt.Printf("Task added: %s\n", taskName)
	},
}

func init() {
	addCmd.DisableFlagsInUseLine = true
	rootCmd.AddCommand(addCmd)
}
