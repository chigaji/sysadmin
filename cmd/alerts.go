package cmd

import (
	"fmt"

	"github.com/chigaji/sysadmin/internal/monitor"
	"github.com/spf13/cobra"
)

var alertsCmd = &cobra.Command{
	Use:   "alerts",
	Short: "Configure alerts for normal condition",
	Run: func(cmd *cobra.Command, args []string) {

		cfg := monitor.NewConfig(10, 20)
		mon := monitor.NewMonitor(cfg)

		// set up alerting rules
		mon.SetRules(20, 30)

		fmt.Println("Alerts configured successfully")
	},
}

func init() {
	rootCmd.AddCommand(alertsCmd)
}
