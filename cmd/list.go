package cmd

import "github.com/spf13/cobra"

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List your tasks",
	Long:  "List your tasks",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
