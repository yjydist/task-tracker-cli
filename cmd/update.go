package cmd

import "github.com/spf13/cobra"

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a task",
	Long:  "Update a task",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {

	},
}
