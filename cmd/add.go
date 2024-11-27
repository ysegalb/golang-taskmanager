package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"taskmanager/app"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add \"task description\"",
	Short: "Add a task to the task manager",
	Long:  `Add a task to the task manager`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("please provide a task description")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		repository, err := app.NewFileTaskRepository("tasks.json")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		service := app.NewTaskService(repository)

		task, err := service.AddTask(args[0])
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Printf("Task added successfully (ID: %d)\n", task.ID)
	},
}
