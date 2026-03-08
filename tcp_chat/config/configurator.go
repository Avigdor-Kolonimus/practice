package config

import (
	"encoding/json"
	"io"
	"os"
)

const (
	ServerConfigPath = "config/practice/conn.json"
)

// ------------------------------------------------------------ ServerConfig ------------------------------------------------------------

type ServerConfig struct {
	Port int `json:"port"`
}

func LoadServerConfig(path string) (*ServerConfig, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var config ServerConfig
	err = json.Unmarshal(jsonData, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}