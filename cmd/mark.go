package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var markCmd = &cobra.Command{
	Use:   "mark-[status] [id]",
	Short: "Mark task status",
}

var markInProgressCmd = &cobra.Command{
	Use:   "in-progress [id]",
	Short: "Mark task as in progress",
	Args:  cobra.ExactArgs(1),
	Run:   markStatus("in-progress"),
}

var markDoneCmd = &cobra.Command{
	Use:   "done [id]",
	Short: "Mark task as done",
	Args:  cobra.ExactArgs(1),
	Run:   markStatus("done"),
}

func init() {
	markCmd.AddCommand(markInProgressCmd)
	markCmd.AddCommand(markDoneCmd)
	rootCmd.AddCommand(markCmd)
}

func markStatus(status string) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		id, _ := strconv.Atoi(args[0])
		tasks := loadTasks()

		found := false
		for i := range tasks {
			if tasks[i].ID == id {
				tasks[i].Status = status
				tasks[i].UpdatedAt = time.Now()
				found = true
				break
			}
		}

		if !found {
			fmt.Printf("Task with ID %d not found\n", id)
			return
		}

		saveTasks(tasks)
		fmt.Printf("Task marked as %s successfully\n", status)
	}
}
