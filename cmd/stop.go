package cmd

import (
	"github.com/chigaji/sysadmin/internal/monitor"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the monitoring process",
	Run: func(cmd *cobra.Command, args []string) {

		cfg := monitor.NewConfig(0, 0)
		mon := monitor.NewMonitor(cfg)

		mon.Stop()

	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
