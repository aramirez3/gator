package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/aramirez3/gator/internal/database"
	"github.com/google/uuid"
)

const (
	configFileName = ".gatorconfig.json"
)

type Config struct {
	DBUrl           string    `json:"db_url"`
	CurrentUserName string    `json:"current_user_name"`
	CurrentUserId   uuid.UUID `json:"current_user_id"`
}

func getConfigFilepath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error getting home directory: %w", err)
	}
	return filepath.Join(homeDir, configFileName), nil
}

func Read() (Config, error) {
	config := Config{}
	filePath, err := getConfigFilepath()
	if err != nil {
		return config, err
	}

	configFile, err := os.Open(filePath)
	if err != nil {
		return config, err
	}
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}
	return config, nil
}

func (c *Config) SetUser(user database.User) error {
	c.CurrentUserName = user.Name
	c.CurrentUserId = user.ID

	write(*c)
	return nil
}

func write(cfg Config) error {
	filePath, err := getConfigFilepath()
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}
	return nil
}
