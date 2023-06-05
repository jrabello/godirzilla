package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

func configFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("unable to get user home directory: %w", err)
	}

	return filepath.Join(homeDir, ".godirzilla", "config.json"), nil
}

func LoadConfig() (*Config, error) {
	var cfg Config

	homeDir, _ := os.UserHomeDir()
	configFilePath := path.Join(homeDir, ".godirzilla", "config.json")

	data, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("Error reading config file: %w", err)
	}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling config file: %w", err)
	}

	return &cfg, nil
}

func SaveConfig(config *Config) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("unable to encode config: %w", err)
	}

	filePath, err := configFilePath()
	if err != nil {
		return err
	}

	// Create the directory if it doesn't exist
	err = os.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil {
		return fmt.Errorf("unable to create config directory: %w", err)
	}

	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("unable to write config file: %w", err)
	}

	return nil
}
