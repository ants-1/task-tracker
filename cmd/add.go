package cmd

import (
	"fmt"
	"strings"

	"github.com/ants-1/task-tracker/service"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Long:  `Add a new task to tasks.json `,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		description := strings.Join(args, " ")

		newTask := service.Task{
			Description: description,
			Status:      "todo",
		}

		service.CreateTaskFile()

		err := service.AddTask(newTask)
		if err != nil {
			fmt.Println("Error adding task:", err)
			return
		}

		fmt.Println("Task added sucessfully!")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
