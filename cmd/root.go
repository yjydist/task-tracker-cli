// Package cmd
// Description: This file contains the root command for the Task Tracker CLI
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "task-tracker",
	Short: "Task Tracker is a CLI for managing your tasks",
	Long:  "Task Tracker is a CLI for managing your tasks. It is a simple CLI that allows you to add, list, and complete tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Task Tracker, type --help for usage")
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
	RootCmd.AddCommand(listCmd)
	RootCmd.AddCommand(updateCmd)
	RootCmd.AddCommand(versionCmd)
}
