package monitor

import (
	"fmt"
	"time"

	"github.com/chigaji/sysadmin/internal/config"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

// represent the system resource monitor
type Monitor struct {
	config *config.Config
}

// create a new monitor instance
func NewMonitor(cfg *config.Config) *Monitor {
	return &Monitor{
		config: cfg,
	}
}

// create new instance of config
func NewConfig(cpuThreshold float64, memThreshold float64) *config.Config {
	return &config.Config{
		CPUThreshold:    cpuThreshold,
		MemoryThreshold: memThreshold,
	}
}

// check for abnormal resource usages
func (m *Monitor) isAbnormal(cpuPercent, memUsedPercent float64) bool {

	if cpuPercent > m.config.CPUThreshold || memUsedPercent > m.config.MemoryThreshold {
		return true
	}
	return false
}

// start monitoring
func (m *Monitor) Start() (<-chan float64, <-chan float64, <-chan error) {
	cpuUsageCh := make(chan float64)
	memoryUsageCh := make(chan float64)
	chanError := make(chan error)

	fmt.Println("Starting Monitoring....")

	go func() {

		defer close(cpuUsageCh)
		defer close(memoryUsageCh)
		defer close(chanError)

		ticker := time.NewTicker(time.Second * 1)
		defer ticker.Stop()

		for range ticker.C {
			// fetch system metrics
			cpuPercent, err := cpu.Percent(time.Second, false)
			if err != nil {
				chanError <- fmt.Errorf("error getting cpu Usage: %v", err)
				continue
			}

			memInfo, err := mem.VirtualMemory()

			if err != nil {
				chanError <- fmt.Errorf("error getting Memory Usage: %v", err)
				continue
			}

			// send metrics through the channel
			cpuUsageCh <- cpuPercent[0]
			memoryUsageCh <- memInfo.UsedPercent

			// check for abnormal conditions
			if m.isAbnormal(cpuPercent[0], memInfo.UsedPercent) {
				m.GenerateAlert(cpuPercent[0], memInfo.UsedPercent)
			}
		}
	}()

	return cpuUsageCh, memoryUsageCh, chanError
}

// printCn := make(chan struct{})
var printCn = make(chan struct{})

// stop monitoring
func (m *Monitor) Stop() {
	fmt.Println("Monitoring stopped...")
	close(printCn)
}

// retrieves real-time metrics
func (m *Monitor) GetMetrics() (float64, float64, error) {

	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0, 0, fmt.Errorf("error getting cpu Usage: %v", err)
	}

	memInfo, err := mem.VirtualMemory()

	if err != nil {
		return 0, 0, fmt.Errorf("error getting Memory Usage: %v", err)
	}

	return cpuPercent[0], memInfo.UsedPercent, nil
}

// Print system usage metrics
func (m *Monitor) PrintMetrics(cpuUsageCh <-chan float64, memoryUsageCh <-chan float64, errorCh <-chan error) {

	// Start goroutine to print metrics
	go func() {
		for {
			select {
			case cpuUsage := <-cpuUsageCh:
				fmt.Printf("CPU Usage : %.2f%%\n", cpuUsage)
			case memoryUsage := <-memoryUsageCh:
				fmt.Printf("Memory Usage : %.2f%%\n", memoryUsage)
			case err := <-errorCh:
				fmt.Printf("Error : %v\n", err)
			case <-printCn:
				//Stop Printing metrics if the program is terminated
				fmt.Println("Stopping Metrics service")
				return
			}
		}
	}()
}

// configures monitoring rules
func (m *Monitor) SetRules(CPUThreshold, MemoryThreshold float64) error {
	m.config.CPUThreshold = CPUThreshold
	m.config.MemoryThreshold = MemoryThreshold
	return nil
}

// generates alert for abnormal condition
func (m *Monitor) GenerateAlert(cpuPercent float64, memUsedPercent float64) {
	redColor := "\033[31m"
	colorReset := "\033[0m"
	fmt.Printf("%sAlert!! Abnormal conditions detected- CPU: %.2f%%, Memory: %.2f%%%s\n", redColor, cpuPercent, memUsedPercent, colorReset)
}
