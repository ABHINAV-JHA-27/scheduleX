package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [job-name] [cron-schedule] [command]",
	Short: "Adds a new job to the schedule",
	Long: `Adds a new job. Requires a unique name, a valid cron string, and the command to be executed.
Example:
scheduleX add backup-db "0 2 * * *" /usr/bin/backup-script.sh`,
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		jobName := args[0]
		schedule := args[1]
		command := args[2]
		fmt.Printf("add command called for job: %s, schedule: %s, command: %s\n", jobName, schedule, command)
		// TODO: Add logic to save the new job
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
