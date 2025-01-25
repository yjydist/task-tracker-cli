// Package cmd
// Description: This file contains the version command for the Task Tracker CLI
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Task Tracker",
	Long:  "Print the version number of Task Tracker",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Task Tracker v0.1")
	},
}
