package utils

import (
	"cloudcrafter/pkg/models"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// ParseYAMLConfig reads a YAML file and parses it into a Configuration struct
func ParseYAMLConfig(filePath string) (*models.Configuration, error) {
	data, err := os.ReadFile(filePath) // Updated to use os.ReadFile
	if err != nil {
		return nil, fmt.Errorf("failed to read YAML file: %w", err)
	}

	var config models.Configuration
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	return &config, nil
}
