package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const (
	configFileName = ".gatorconfig.json"
)

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFile(readWriteMode bool) (*os.File, error) {
	// must call Close() where consuming the returned file
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("error getting home directory: %w", err)
	}

	filePath := filepath.Join(homeDir, configFileName)
	if readWriteMode {
		configFile, err := os.OpenFile(filePath, os.O_RDWR, 0644)
		if err != nil {
			return nil, fmt.Errorf("error opening config file: %w", err)
		}

		return configFile, nil
	} else {
		configFile, err := os.Open(filePath)
		if err != nil {
			return nil, fmt.Errorf("error opening config file: %w", err)
		}

		return configFile, nil
	}
}

func Read() (Config, error) {
	config := Config{}
	configFile, err := getConfigFile(false)

	if err != nil {
		return config, fmt.Errorf("error getting config file: %w", err)
	}

	defer configFile.Close()

	byteData, err := io.ReadAll(configFile)
	if err != nil {
		return config, fmt.Errorf("error reading config file: %w", err)
	}

	json.Unmarshal(byteData, &config)
	fmt.Println("Config file:")
	fmt.Printf("    db_url: %s\n", config.DBUrl)
	fmt.Printf("    current_user_name: %s\n", config.CurrentUserName)
	return config, nil
}

func SetUser(userName string) error {
	cfg, err := Read()
	if err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}
	cfg.CurrentUserName = userName

	write(cfg)
	return nil
}

func write(cfg Config) error {
	configFile, err := getConfigFile(true)
	if err != nil {
		return fmt.Errorf("error getting config file: %w", err)
	}
	defer configFile.Close()

	encoder := json.NewEncoder(configFile)
	err = encoder.Encode(cfg)
	if err != nil {
		return fmt.Errorf("error encoding json file: %w", err)
	}
	return nil
}
