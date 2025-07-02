package utils

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/Omotolani98/cloudcrafter/pkg/models"
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

func WriteYaml(config models.Configuration, path string) error {
	if path == "" {
		path = "config.yml"
	}
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error creating YAML file:", err)
		return err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	encoder := yaml.NewEncoder(file)
	if err := encoder.Encode(&config); err != nil {
		fmt.Println("Error writing YAML:", err)
		return err
	}

	return nil
}
