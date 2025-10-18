package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all scheduled jobs",
	Long:  `Connects to the scheduler and prints a table of all currently loaded jobs and their schedules.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list command called")
		// TODO: Add logic to list jobs
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
