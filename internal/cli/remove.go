package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove [job-name]",
	Short: "Removes a job from the schedule",
	Long:  `Finds a job by its name and removes it from the schedule.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		jobName := args[0]
		fmt.Printf("remove command called for job: %s\n", jobName)
		// TODO: Add logic to remove the job
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
