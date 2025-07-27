package config

import (
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

const configPath = "/etc/samba/mounts.yaml"

// Mount represents a single mount point
type Mount struct {
	Src    string `yaml:"src"`
	Dst    string `yaml:"dst"`
	Type   string `yaml:"type"`
	Option string `yaml:"option,omitempty"`
}

// DefaultOptions represents the default options for mounting
type DefaultOptions struct {
	Type    string `yaml:"type"`
	Options string `yaml:"options"`
}

// Config represents the structure of the mounts.yaml file
type Config struct {
	Mounts struct {
		DefaultOptions DefaultOptions `yaml:"default-options"`
		Points         []Mount        `yaml:"points"`
	} `yaml:"mounts"`
}

// LoadConfig loads the configuration from the specified file
func LoadConfig(filename string) (*Config, error) {
	slog.Debug("Loading config from file", "filename", filename)

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decode the YAML
	var config Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	// Apply default options to mount points if not specified
	for i := range config.Mounts.Points {
		if config.Mounts.Points[i].Type == "" {
			config.Mounts.Points[i].Type = config.Mounts.DefaultOptions.Type
		}
		if config.Mounts.Points[i].Option == "" {
			config.Mounts.Points[i].Option = config.Mounts.DefaultOptions.Options
		}
	}

	return &config, nil
}

// GetConfigPath returns the path to the configuration file
func GetConfigPath() (string, error) {
	return configPath, nil
}
