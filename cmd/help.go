package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Display use instructions",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Root().UsageString())
	},
}

func init() {
	rootCmd.AddCommand(helpCmd)
}
