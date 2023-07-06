package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/jrabello/godirzilla/file"
)

var cognusCraft = ".cognuscraft"
var godirzilla = "godirzilla"
var configFileName = "config.json"

func UpsertNewConfigDirectory() error {
	filePath := getConfigFilePath()

	// Create the directory if it doesn't exist
	err := os.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil {
		return fmt.Errorf("unable to create config directory: %w", err)
	}

	return nil
}

func isConfigFileExists() bool {
	configFilePath := getConfigFilePath()
	return file.IsFileExists(configFilePath)
}

func getConfigFilePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Unable to get user home directory: %v", err)
	}

	return filepath.Join(homeDir, cognusCraft, godirzilla, configFileName)
}

// LoadConfig loads the config file into memory
func LoadConfig() (*Config, error) {
	configFilePath := getConfigFilePath()
	data, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("Error reading config file: %w", err)
	}

	var cfg Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling config file: %w", err)
	}

	return &cfg, nil
}

// SaveConfig saves config file to disk
func SaveConfig(config *Config) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("unable to encode config: %w", err)
	}

	configFilePath := getConfigFilePath()
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
