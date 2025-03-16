package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := strconv.Atoi(args[0])
		tasks := loadTasks()
		newTasks := []Task{}

		found := false
		for _, task := range tasks {
			if task.ID == id {
				found = true
			} else {
				newTasks = append(newTasks, task)
			}
		}

		if !found {
			fmt.Printf("Task with ID %d not found\n", id)
			return
		}

		saveTasks(newTasks)
		fmt.Println("Task deleted successfully")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
