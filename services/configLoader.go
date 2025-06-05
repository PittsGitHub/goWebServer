package services

import (
	"encoding/json"
	"log"
	"os"

	"github.com/PittsGitHub/goWebServer/types"
)

var DefaultConfig = types.ConfigFile{
	Port: "5000",
}

func LoadConfig(path string) types.ConfigFile {
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			// Config file doesn't exist, create with defaults
			log.Printf("Config file not found, creating default at %s\n", path)
			defaultBytes, err := json.MarshalIndent(DefaultConfig, "", "  ")
			if err != nil {
				log.Fatalf("Failed to marshal default config: %v", err)
			}

			err = os.WriteFile(path, defaultBytes, 0644)
			if err != nil {
				log.Fatalf("Failed to write default config file: %v", err)
			}

			return DefaultConfig
		}
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer file.Close()

	bytes, err := os.ReadFile(file.Name())
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	var config types.ConfigFile
	if err := json.Unmarshal(bytes, &config); err != nil {
		log.Fatalf("Failed to parse config JSON: %v", err)
	}

	return config
}
