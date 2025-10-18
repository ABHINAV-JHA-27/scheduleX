package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "scheduleX",
	Short: "A CLI-based task scheduler.",
	Long: `scheduleX is a lightweight, cron-like job scheduler
that runs tasks defined in your configuration.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI: %s", err)
		os.Exit(1)
	}
}
