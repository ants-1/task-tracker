package cmd

import (
	"fmt"
	"strconv"

	"github.com/ants-1/task-tracker/service"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete existing task by ID.",
	Long:  `Delete existing task by ID from tasks.json file.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid task ID:", args[0])
			return
		}

		err = service.DeleteTask(id)
		if err != nil {
			fmt.Printf("Error deleting task: %v\n", err)
			return
		}
		fmt.Printf("Task %d deleted successfully!\n", id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
