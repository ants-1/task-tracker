package cmd

import (
	"fmt"
	"strconv"

	"github.com/ants-1/task-tracker/service"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update existing task.",
	Long:  `Update existing task in tasks.json file with new data.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid task ID:", args[0])
			return
		}

		newDescription := args[1]

		err = service.EditTask(id, newDescription)
		if err != nil {
			fmt.Printf("Error updating task: %v\n", err)
			return
		}
		fmt.Printf("Task %d updated successfully!\n", id)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
