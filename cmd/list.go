package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list [status]",
	Short: "List tasks",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tasks := loadTasks()
		statusFilter := ""
		if len(args) > 0 {
			statusFilter = args[0]
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "ID\tDescription\tStatus\tCreated At\tUpdated At")

		for _, task := range tasks {
			if statusFilter == "" || task.Status == statusFilter {
				fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\n",
					task.ID,
					task.Description,
					task.Status,
					task.CreatedAt.Format("2006-01-02 15:04:05"),
					task.UpdatedAt.Format("2006-01-02 15:04:05"),
				)
			}
		}

		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
