package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Context struct {
	ServerURL string `json:"server_url"`
	APIKey    string `json:"api_key"`
}

type Config struct {
	CurrentContext string             `json:"current_context"`
	Contexts       map[string]Context `json:"contexts"`
}

func GetDefaultConfigPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ".flexcli.json"
	}
	return filepath.Join(home, ".flexcli.json")
}

func LoadConfig(path string) (*Config, error) {
	if path == "" {
		path = GetDefaultConfigPath()
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &Config{
				Contexts: make(map[string]Context),
			}, nil
		}
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Try to unmarshal into new format
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err == nil && cfg.Contexts != nil {
		return &cfg, nil
	}

	// Migration logic: handle old single-object format
	var tempMap map[string]interface{}
	if err := json.Unmarshal(data, &tempMap); err != nil {
		// If we're here, the file exists but it's not valid JSON at all
		return nil, fmt.Errorf("malformed config file (%s): invalid JSON syntax. Please check for missing commas or quotes", filepath.Base(path))
	}

	if _, ok := tempMap["server_url"]; ok {
		// This is the old format
		newCfg := &Config{
			CurrentContext: "default",
			Contexts: map[string]Context{
				"default": {
					ServerURL: fmt.Sprintf("%v", tempMap["server_url"]),
					APIKey:    fmt.Sprintf("%v", tempMap["api_key"]),
				},
			},
		}
		return newCfg, nil
	}

	return &Config{
		Contexts: make(map[string]Context),
	}, nil
}
func SaveConfig(path string, cfg *Config) error {
	if path == "" {
		path = GetDefaultConfigPath()
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	// Create with 0600 permissions
	return os.WriteFile(path, data, 0600)
}
