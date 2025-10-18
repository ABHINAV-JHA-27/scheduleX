package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Shows the status of the scheduler daemon",
	Long:  `Checks if the scheduleX daemon is running and reports its status and uptime.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("status command called")
		// TODO: Add logic to check daemon status
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
