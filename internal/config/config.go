package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// represent the configuration for monitoring rules
type Config struct {
	CPUThreshold    float64 `yaml:"cpu_threshold"`
	MemoryThreshold float64 `yaml:"memory_threshold"`
}

// load configs from a file
func LoadConfig(f string) (*Config, error) {

	data, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %v", err)
	}

	return &cfg, nil
}

// save configurations to file
func SaveConfig(config *Config, f string) error {

	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("error marshalling config: %v", err)
	}
	err = ioutil.WriteFile(f, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing config file: %v", err)
	}

	return nil
}
