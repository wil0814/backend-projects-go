package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
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
		// 判斷是否提供了 status 參數
		if len(args) == 0 {
			// 沒有提供 status，列出所有任務
			fmt.Println("Listing all tasks:")
			// 模擬列出所有任務的邏輯
			fmt.Println("1. Task A [todo]")
			fmt.Println("2. Task B [done]")
			fmt.Println("3. Task C [in-progress]")
		} else {
			// 根據提供的 status 過濾任務
			status := args[0]
			switch status {
			case "done":
				fmt.Println("Listing tasks with status: done")
				// 模擬列出 "done" 任務
				fmt.Println("2. Task B [done]")
			case "todo":
				fmt.Println("Listing tasks with status: todo")
				// 模擬列出 "todo" 任務
				fmt.Println("1. Task A [todo]")
			case "in-progress":
				fmt.Println("Listing tasks with status: in-progress")
				// 模擬列出 "in-progress" 任務
				fmt.Println("3. Task C [in-progress]")
			default:
				// 無效的 status，提示錯誤
				fmt.Printf("Invalid status: %s. Valid statuses are: done, todo, in-progress.\n", status)
			}
		}
	},
}

func init() {
	listCmd.DisableFlagsInUseLine = true
	rootCmd.AddCommand(listCmd)
}
