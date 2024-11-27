package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"taskmanager/app"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list [status]",
	Short: "List all tasks or tasks by status",
	Long:  `List all tasks or filter by status (todo, in-progress, done)`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return fmt.Errorf("accepts at most 1 arg")
		}
		if len(args) == 1 && args[0] != "todo" && args[0] != "in-progress" && args[0] != "done" {
			return fmt.Errorf("invalid status: %s (must be todo, in-progress, or done)", args[0])
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

		var tasks []app.Task
		if len(args) == 1 {
			status := app.Status(args[0])
			tasks, err = service.ListTasksByStatus(status)
		} else {
			tasks, err = service.ListTasks()
		}

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		if len(tasks) == 0 {
			fmt.Println("No tasks found")
			return
		}

		fmt.Println("Tasks:")
		for _, task := range tasks {
			fmt.Printf("[%d] %s (Status: %s)\n", task.ID, task.Description, task.Status)
		}
	},
}
