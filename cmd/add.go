package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your task list",
	Long:  "Add a new task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add called")
	},
}
