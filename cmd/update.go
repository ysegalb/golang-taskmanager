package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"taskmanager/app"
)

func init() {
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update [id] [description]",
	Short: "Update a task description",
	Long:  `Update the description of an existing task by its ID`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Error: invalid task ID: %v\n", err)
			return
		}

		repository, err := app.NewFileTaskRepository("tasks.json")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		service := app.NewTaskService(repository)

		task, err := service.UpdateTask(id, args[1])
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Printf("Task updated successfully (ID: %d)\n", task.ID)
	},
}
