package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "task-cli",
	Short: "A Task Tracker CLI took",
	Long:  `Task tracker is a project used to track and manage your tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		tasksFilePath := os.Getenv("TASKS_FILE_PATH")

		if tasksFilePath == "" {
			fmt.Println("Welcome to Task CLI!")
			fmt.Println("The environment variable 'TASKS_FILE_PATH' is not set.")
			fmt.Println("By default, tasks.json will be created in the current directory.")
			fmt.Println("You can set the 'TASKS_FILE_PATH' environment variable to specify a custom directory for tasks.json.")
		} else {
			fmt.Printf("Welcome to Task CLI!\nThe tasks.json file will be saved in: %s\n", tasksFilePath)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
