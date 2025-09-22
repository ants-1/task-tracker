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
	Run: func(cmd *cobra.Command, args []string) {
		service.CreateTaskFile()

		tasks, err := service.LoadTasks()
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			return
		}

		for _, t := range tasks {
			fmt.Printf("%d. %s (%s)\n", t.ID, t.Description, t.Status)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
