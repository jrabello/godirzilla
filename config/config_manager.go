package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type FilePath string

func CreateConfigFileIfNotExists() {
	if isConfigFileExists() {
		return
	}

	cfg := Config{
		CurrentGroup: "main",
		Groups:       make(map[Group][]Directory),
	}
	cfg.Groups["main"] = []Directory{}

	SaveConfig(&cfg)
}

func UpsertNewConfigDirectory() error {
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	// Create the directory if it doesn't exist
	err = os.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil {
		return fmt.Errorf("unable to create config directory: %w", err)
	}

	return nil
}

func isConfigFileExists() bool {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return false
	}

	_, err = os.Stat(configFilePath)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return false
}

func LoadConfig() (*Config, error) {
	var cfg Config

	configFilePath, err := getConfigFilePath()
	if err != nil {
		return nil, err
	}

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

	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	err = UpsertNewConfigDirectory()
	if err != nil {
		return fmt.Errorf("unable to create file: %w", err)
	}

	err = ioutil.WriteFile(configFilePath, data, 0644)
	if err != nil {
		return fmt.Errorf("unable to write config file: %w", err)
	}

	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("unable to get user home directory: %w", err)
	}

	return filepath.Join(homeDir, ".godirzilla", "config.json"), nil
}
