package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure Monitoring rules",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Configuring Monitoring Rules...")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
