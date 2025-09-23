package cmd

import (
	"fmt"

	"github.com/ants-1/task-tracker/service"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  `List all tasks stored in tasks.json. If the file foes not exist, it will be created.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		service.CreateTaskFile()

		var tasks []service.Task
		var err error

		if len(args) == 0 {
			tasks, err = service.LoadTasks()
		} else {
			tasks, err = service.FilteredTasks(args[0])
		}

		if err != nil {
			fmt.Println("Error loading tasks:", err)
			return
		}

		if len(tasks) < 1 {
			fmt.Println("Tasks list it empty")
			return
		}

		for _, t := range tasks {
			fmt.Printf("%d. %s (%s), created: %s, updated: %s\n",
				t.ID,
				t.Description,
				t.Status,
				t.CreatedAt,
				t.UpdatedAt,
			)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
