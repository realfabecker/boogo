package config

import (
	"encoding/json"
	"fmt"
	"github.com/realfabecker/bogo/internal/core/domain"
	"github.com/realfabecker/bogo/internal/core/ports"
	"os"
	"path/filepath"
)

// JsonConfigRepository struct definition
type JsonConfigRepository struct{}

// NewJsonConfigRepository json config repository constructor
func NewJsonConfigRepository() ports.ConfigRepository {
	return &JsonConfigRepository{}
}

// Get obtain the config from a json repository
func (c JsonConfigRepository) Get() (*domain.BogoConfig, error) {
	h, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	p := filepath.Join(h, ".bogo", "config.json")
	d, err := os.ReadFile(p)
	if err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}

	var config domain.BogoConfig
	err = json.Unmarshal(d, &config)
	if err != nil {
		return nil, fmt.Errorf("json: %w", err)
	}
	return &config, nil
}

// Save marshal a config struct into a json
func (c JsonConfigRepository) Save(config *domain.BogoConfig) error {
	h, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("save: %w", err)
	}
	p := filepath.Join(h, ".bogo", "config.json")

	d, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("json: %w", err)
	}
	err = os.WriteFile(p, d, 0644)
	if err != nil {
		return fmt.Errorf("save: %w", err)
	}
	return nil
}
