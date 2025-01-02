package utils

import (
	"cloudcrafter/pkg/models"
	"gopkg.in/yaml.v2"
	"os"
)

// ParseYAML parses a YAML file into ResourceConfig
func ParseYAML(filePath string) (*models.Configuration, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config models.Configuration
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// GenerateYAML generates a YAML representation of a ResourceConfig
func GenerateYAML(resource models.Resource) ([]byte, error) {
	config := models.Configuration{
		Provider: "aws",
		Resources: []models.Resource{
			resource,
		},
	}
	return yaml.Marshal(config)
}
