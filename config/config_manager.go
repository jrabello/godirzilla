package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Config struct {
	Directories []string `json:"directories"`
}

func configFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("unable to get user home directory: %w", err)
	}

	return filepath.Join(homeDir, ".godirzilla", "config.json"), nil
}

func LoadConfig() (Config, error) {
	var config Config

	filePath, err := configFilePath()
	if err != nil {
		return config, err
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		// If the file doesn't exist, that's okay; we'll create it when we save the config
		if os.IsNotExist(err) {
			return config, nil
		}

		return config, fmt.Errorf("unable to read config file: %w", err)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("unable to parse config file: %w", err)
	}

	return config, nil
}

func SaveConfig(config Config) error {
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
