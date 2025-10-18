package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the scheduleX daemon",
	Long:  `Starts the scheduleX daemon process that runs in the background.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start command called")
		// TODO: Add logic to start the scheduler daemon
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
