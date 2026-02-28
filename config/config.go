package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/yourusername/dbmigrate/database"
	"github.com/yourusername/dbmigrate/migration"
)

// Config represents the configuration for DbMigrate
type Config struct {
	Database database.Config `json:"database"`
	Migrations []migration.Config `json:"migrations"`
}

// LoadConfig loads the configuration from the given file path
func LoadConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	// Validate configuration
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// Validate validates the configuration
func (cfg *Config) Validate() error {
	if cfg.Database.Dialect == "" {
		return errors.New("database dialect is required")
	}

	if cfg.Database.Username == "" {
		return errors.New("database username is required")
	}

	if cfg.Database.Password == "" {
		return errors.New("database password is required")
	}

	if cfg.Database.Host == "" {
		return errors.New("database host is required")
	}

	if cfg.Database.Port == 0 {
		return errors.New("database port is required")
	}

	if cfg.Database.Name == "" {
		return errors.New("database name is required")
	}

	return nil
}