package cmd

import (
	"fmt"

	"github.com/chigaji/sysadmin/internal/monitor"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var thresholdsCmd = &cobra.Command{
	Use:   "thresholds",
	Short: "Set CPU and Memory thresholds",
	Run: func(cmd *cobra.Command, args []string) {
		// read thresholds from viper
		cpuThreshold := viper.GetFloat64("cpu_threshold")
		memoryThreshold := viper.GetFloat64("memory_threshold")

		//Initialize monitor with configuration
		cfg := monitor.NewConfig(cpuThreshold, memoryThreshold)
		mon := monitor.NewMonitor(cfg)

		mon.SetRules(cpuThreshold, memoryThreshold)

		fmt.Printf("Setting cpu threshold: %.2f%%\n", cpuThreshold)
		fmt.Printf("Setting memory threshold: %.2f%%\n", memoryThreshold)
	},
}

func init() {
	rootCmd.AddCommand(thresholdsCmd)
}
