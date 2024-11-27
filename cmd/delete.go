package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"taskmanager/app"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a task",
	Long:  `Delete an existing task by its ID`,
	Args:  cobra.ExactArgs(1),
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

		if err := service.DeleteTask(id); err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Printf("Task deleted successfully (ID: %d)\n", id)
	},
}
