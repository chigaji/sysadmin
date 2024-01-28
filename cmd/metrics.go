package cmd

import (
	"github.com/chigaji/sysadmin/internal/monitor"
	"github.com/spf13/cobra"
)

var metricsCmd = &cobra.Command{
	Use:   "metrics",
	Short: "View real-time metrics",
	Run: func(cmd *cobra.Command, args []string) {

		cfg := monitor.NewConfig(0, 0)
		mon := monitor.NewMonitor(cfg)

		//start monitoring and display real-time metrics
		cpuUsageCh, memoryUsageCh, errorCh := mon.Start()

		mon.PrintMetrics(cpuUsageCh, memoryUsageCh, errorCh)

		// fmt.Println("View real-time metrics..")
	},
}

func init() {
	rootCmd.AddCommand(metricsCmd)
}
