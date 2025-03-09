package test

import (
	"encoding/json"
	"testing"

	"github.com/Omotolani98/cloudcrafter/pkg/models"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestJSONParsing(t *testing.T) {
	// Example JSON configuration
	jsonData := `{
		"provider": "aws",
		"resources": [
			{
				"type": "vm",
				"name": "test-instance",
				"machineType": "t2.micro",
				"region": "us-east-1",
				"image": "ami-0abcdef1234567890"
			}
		]
	}`

	var config models.Configuration
	err := json.Unmarshal([]byte(jsonData), &config)

	assert.NoError(t, err, "JSON unmarshalling should not throw an error")
	assert.Equal(t, "aws", config.Provider)
	assert.Len(t, config.Resources, 1)
	assert.Equal(t, "vm", config.Resources[0]["type"])
	assert.Equal(t, "test-instance", config.Resources[0]["name"])
}

func TestYAMLParsing(t *testing.T) {
	// Example YAML configuration
	yamlData := `
		provider: aws
		resources:
		- type: vm
			name: test-instance
			machineType: t2.micro
			region: us-east-1
			image: ami-0abcdef1234567890
	`

	var config models.Configuration
	err := yaml.Unmarshal([]byte(yamlData), &config)

	assert.NoError(t, err, "YAML unmarshalling should not throw an error")
	assert.Equal(t, "aws", config.Provider)
	assert.Len(t, config.Resources, 1)
	assert.Equal(t, "vm", config.Resources[0]["type"])
	assert.Equal(t, "test-instance", config.Resources[0]["name"])
}
