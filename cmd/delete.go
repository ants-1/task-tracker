package cmd

import (
	"fmt"

	"github.com/ants-1/task-tracker/service"
	"github.com/spf13/cobra"
)

var deleteID int

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete existing task by ID.",
	Long:  `Delete existing task by ID from tasks.json file.`,
	Run: func(cmd *cobra.Command, args []string) {
		if deleteID == 0 {
			fmt.Println("Please provide a valid task ID with --id")
		}

		err := service.DeleteTask(deleteID)
		if err != nil {
			fmt.Printf("Error deleting task: %v\n", err)
			return
		}
		fmt.Printf("Task %d deleted successfully!\n", deleteID)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().IntVarP(&deleteID, "id", "i", 0, "ID of the task to delete")
	deleteCmd.MarkFlagRequired("id")
}
