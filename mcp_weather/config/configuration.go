package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	DefaultWeatherServerConfigPath = "config/weaterserverconfig.json"
)

// WeatherServerConfig is loaded from weaterserverconfig.json.
// Durations in the file are in seconds. Nil fields were not set in JSON.
type WeatherServerConfig struct {
	RequestTimeout int `json:"requestTimeout,omitempty"`
	ToolTimeout    int `json:"toolTimeout,omitempty"`
	CacheTTL       int `json:"cacheTTL,omitempty"`
}

// ResolvePath finds path relative to the working directory or any parent (up to module root).
func ResolvePath(path string) (string, error) {
	if filepath.IsAbs(path) {
		if _, err := os.Stat(path); err != nil {
			return "", fmt.Errorf("config file not found: %q", path)
		}
		return path, nil
	}

	if _, err := os.Stat(path); err == nil {
		return filepath.Abs(path)
	}

	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		candidate := filepath.Join(dir, path)
		if _, err := os.Stat(candidate); err == nil {
			return filepath.Abs(candidate)
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	return "", fmt.Errorf("config file not found: %q (searched from working directory upward)", path)
}

// Load reads and parses the JSON config file at path.
func Load(path string) (*WeatherServerConfig, error) {
	resolved, err := ResolvePath(path)
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(resolved)
	if err != nil {
		return nil, fmt.Errorf("read config %q: %w", resolved, err)
	}

	return Parse(data)
}

// Parse unmarshals config JSON from data.
func Parse(data []byte) (*WeatherServerConfig, error) {
	var cfg WeatherServerConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}
	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (c *WeatherServerConfig) validate() error {
	if c.RequestTimeout <= 0 {
		return fmt.Errorf("requestTimeout must be positive, got %d", c.RequestTimeout)
	}

	if c.ToolTimeout <= 0 {
		return fmt.Errorf("toolTimeout must be positive, got %d", c.ToolTimeout)
	}

	if c.CacheTTL <= 0 {
		return fmt.Errorf("cacheTTL must be positive, got %d", c.CacheTTL)
	}

	return nil
}
