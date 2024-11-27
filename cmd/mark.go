package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"taskmanager/app"
)

func init() {
	rootCmd.AddCommand(markInProgressCmd)
	rootCmd.AddCommand(markDoneCmd)
}

var markInProgressCmd = &cobra.Command{
	Use:   "mark-in-progress [id]",
	Short: "Mark a task as in progress",
	Long:  `Mark an existing task as in progress by its ID`,
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

		if err := service.MarkTaskInProgress(id); err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Printf("Task marked as in progress (ID: %d)\n", id)
	},
}

var markDoneCmd = &cobra.Command{
	Use:   "mark-done [id]",
	Short: "Mark a task as done",
	Long:  `Mark an existing task as done by its ID`,
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

		if err := service.MarkTaskDone(id); err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Printf("Task marked as done (ID: %d)\n", id)
	},
}
