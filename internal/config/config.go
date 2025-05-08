package config

import (
	"encoding/json"
	// "fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

// Config represents the JSON file structure
type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name,omitempty"`
}

// Read reads the JSON file and returns a Config struct
func Read() (Config, error) {
	var cfg Config

	path, err := getConfigFilePath()
	if err != nil {
		return cfg, err
	}

	file, err := os.Open(path)
	if err != nil {
		return cfg, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}

// SetUser sets the current_user_name and writes the config back to disk
func (cfg *Config) SetUser(name string) error {
	cfg.CurrentUserName = name
	return write(*cfg)
}

// getConfigFilePath returns the full path to the config file
func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, configFileName), nil
}

// write writes the Config struct to the JSON file
func write(cfg Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // pretty-print
	if err := encoder.Encode(cfg); err != nil {
		return err
	}

	return nil
}
