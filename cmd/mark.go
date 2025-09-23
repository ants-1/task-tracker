package cmd

import (
	"fmt"
	"strconv"

	"github.com/ants-1/task-tracker/service"
	"github.com/spf13/cobra"
)

func runMark(_ *cobra.Command, args []string, status string) {
	if len(args) < 1 {
		fmt.Println("Please provide a task ID")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid task ID:", args[0])
		return
	}

	err = service.EditTaskStatus(id, status)
	if err != nil {
		fmt.Printf("Error updating task: %v\n", err)
		return
	}

	fmt.Printf("Task %d marked as %s sucessfully\n", id, status)
}

var markInProgressCmd = &cobra.Command{
	Use:   "mark-in-progress [task ID]",
	Short: "Mark a task as in-progress",
	Run: func(cmd *cobra.Command, args []string) {
		runMark(cmd, args, "in-progress")
	},
}

var markDoneCmd = &cobra.Command{
	Use:   "mark-done [taskID]",
	Short: "Mark a task as done",
	Run: func(cmd *cobra.Command, args []string) {
		runMark(cmd, args, "done")
	},
}

var markTodoCmd = &cobra.Command{
	Use:   "mark-todo [task ID]",
	Short: "Mark a task as todo",
	Run: func(cmd *cobra.Command, args []string) {
		runMark(cmd, args, "todo")
	},
}

func init() {
	rootCmd.AddCommand(markInProgressCmd)
	rootCmd.AddCommand(markDoneCmd)
	rootCmd.AddCommand(markTodoCmd)
}
