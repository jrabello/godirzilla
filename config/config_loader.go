package config

import "log"

// LoadGlobalConfig loads config file into memory and creates a new one if it does not exist
func LoadGlobalConfig() (*Config, error) {
	if !isConfigFileExists() {
		cfg := Config{
			CurrentGroup: "main",
			Groups:       make(map[Group][]Directory),
		}
		cfg.Groups["main"] = []Directory{}
		cfg.version = "1.0"

		err := SaveConfig(&cfg)
		if err != nil {
			log.Fatalf("Error trying to save config file: %v", err)
		}
		return &cfg, nil
	}

	globalConfig, err := LoadConfig()
	if err != nil {
		log.Fatalf("Error trying to load config file: %v", err)
	}

	return globalConfig, nil
}
