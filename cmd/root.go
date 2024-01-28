package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/chigaji/sysadmin/internal/config"
	"github.com/chigaji/sysadmin/internal/monitor"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var rootCmd = &cobra.Command{
	Use:   "sysadmin",
	Short: "A CLI tool to monitor system resources",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to sysadmin - a CLI tool for monitoring system resources")
		fmt.Println("For usage instructions, run 'sysadmin help'")
		//display usage
		// fmt.Println(cmd.UsageString())
	},
}

func init() {

	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is config.yaml)")

	//Set up viper for configuration
	// viper.SetConfigName("config")
	// viper.SetConfigType("yaml")
	// viper.AddConfigPath(".")
	// viper.AutomaticEnv()
	// viper.SetEnvPrefix("sysadmin")

	// // set defaults
	// viper.SetDefault("cpu_threshold", 30)
	// viper.SetDefault("memory_threshold", 50)
}
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {

		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AutomaticEnv()
		viper.SetEnvPrefix("sysadmin")

		// set defaults
		// viper.SetDefault("cpu_threshold", 30)
		// viper.SetDefault("memory_threshold", 50)
	}

	// read configuration file
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		os.Exit(1)
	}

	// unmarshal configuration into struct
	var cfg config.Config

	if err := viper.Unmarshal(&cfg); err != nil {
		fmt.Printf("Error unmarshalling config: %v\n", err)
		os.Exit(1)
	}

	//Initialize monitor with configuration
	mon := monitor.NewMonitor(&cfg)
	//start monitoring and display real-time metrics
	cpuUsageCh, memoryUsageCh, errorCh := mon.Start()
	mon.PrintMetrics(cpuUsageCh, memoryUsageCh, errorCh)
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	// <-sigChan
	go func() {
		<-sigChan
		fmt.Println("\nShutting down Gracefully...")
		os.Exit(0)
	}()

}
