package config

import "log"

var globalConfig *Config = nil

// GetGlobalConfig etrieves the global configuration instance
func GetGlobalConfig() *Config {
	if globalConfig == nil {
		log.Fatal("Global config not initialized")
	}
	return globalConfig
}

// LoadGlobalConfig loads config file into memory and creates a new one if it does not exist
func LoadGlobalConfig() error {
	if !isConfigFileExists() {
		cfg := Config{
			CurrentGroup: "main",
			Groups:       make(map[Group][]Directory),
		}
		cfg.Groups["main"] = []Directory{}
		cfg.version = "1.0.0"

		SaveConfig(&cfg)
		globalConfig = &cfg
		return nil
	}

	var err error
	globalConfig, err = LoadConfig()
	return err
}
