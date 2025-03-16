package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [id] [new description]",
	Short: "Update a task description",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := strconv.Atoi(args[0])
		newDesc := args[1]

		tasks := loadTasks()
		found := false

		for i := range tasks {
			if tasks[i].ID == id {
				tasks[i].Description = newDesc
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
		fmt.Println("Task updated successfully")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
