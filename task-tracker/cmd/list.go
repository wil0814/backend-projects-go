package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wil0814/backend-projects-go/task-tracker/internal/model"
	"github.com/wil0814/backend-projects-go/task-tracker/internal/task"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list [status]",
	Short: "List tasks",
	Long:  `List all tasks or filter tasks by their status (e.g., done, todo, in-progress).`,
	Example: `  # List all tasks
  task-cli list

  # List tasks with status "done"
  task-cli list done

  # List tasks with status "todo"
  task-cli list todo

  # List tasks with status "in-progress"
  task-cli list in-progress`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Listing all tasks:")
			tasks, err := task.ListTasks(nil)
			if err != nil {
				fmt.Printf("Failed to list tasks: %v\n", err)
				return
			}
			for _, task := range tasks {
				fmt.Printf("%d. %s [%s]\n", task.ID, task.Description, task.Status)
			}
		} else {
			status := args[0]
			switch status {
			case "done":
				fmt.Println("Listing tasks with status: done")
				status := model.TaskStatusDone
				tasks, err := task.ListTasks(&status)
				if err != nil {
					fmt.Printf("Failed to list tasks: %v\n", err)
					return
				}
				if len(tasks) == 0 {
					fmt.Println("No tasks with status: done")
				}
				for _, task := range tasks {
					fmt.Printf("%d. %s [%s]\n", task.ID, task.Description, task.Status)
				}
			case "todo":
				fmt.Println("Listing tasks with status: todo")
				status := model.TaskStatusTodo
				tasks, err := task.ListTasks(&status)
				if err != nil {
					fmt.Printf("Failed to list tasks: %v\n", err)
					return
				}
				if len(tasks) == 0 {
					fmt.Println("No tasks with status: todo")
				}
				for _, task := range tasks {
					fmt.Printf("%d. %s [%s]\n", task.ID, task.Description, task.Status)
				}
			case "in-progress":
				fmt.Println("Listing tasks with status: in-progress")
				status := model.TaskStatusInProgress
				tasks, err := task.ListTasks(&status)
				if err != nil {
					fmt.Printf("Failed to list tasks: %v\n", err)
					return
				}
				if len(tasks) == 0 {
					fmt.Println("No tasks with status: in-progress")
				}
				for _, task := range tasks {
					fmt.Printf("%d. %s [%s]\n", task.ID, task.Description, task.Status)
				}
			default:
				fmt.Printf("Invalid status: %s\n", status)
			}
		}
	},
}

func init() {
	listCmd.DisableFlagsInUseLine = true
	rootCmd.AddCommand(listCmd)
}
