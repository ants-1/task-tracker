package cmd

import (
	"fmt"

	"github.com/ants-1/task-tracker/service"
	"github.com/spf13/cobra"
)

var (
	updateID          int
	updateDescription string
	updateStatus      string
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update existing task.",
	Long:  `Update existing task in tasks.json file with new data.`,
	Run: func(cmd *cobra.Command, args []string) {
		if updateID == 0 {
			fmt.Println("Please provide a valid task ID with --id")
		}

		err := service.EditTask(updateID, updateDescription, updateStatus)
		if err != nil {
			fmt.Printf("Error updating task: %v\n", err)
		}
		fmt.Printf("Task %d updated sucessfully!\n", updateID)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().IntVarP(&updateID, "id", "i", 0, "ID of the task to update")
	updateCmd.Flags().StringVarP(&updateDescription, "desc", "d", "", "New description of the task")
	updateCmd.Flags().StringVarP(&updateStatus, "status", "s", "", "New status of the task (todo, in-progress, done)")

	updateCmd.MarkFlagRequired("id")
}
