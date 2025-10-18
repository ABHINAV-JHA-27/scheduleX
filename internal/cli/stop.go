package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops the scheduleX daemon",
	Long:  `Finds the running scheduleX daemon and sends it a signal to gracefully shut down.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stop command called")
		// TODO: Add logic to stop the scheduler daemon
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
