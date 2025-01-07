package models

import (
	"fmt"
	"time"
)

type Configuration struct {
	Provider  string                `json:"provider" yaml:"provider"` // Cloud provider (e.g., AWS, Azure, GCP)
	Resources []map[string]Resource `json:"resources" yaml:"resources"`
}

// Resource defines a single cloud resource configuration
type Resource struct {
	Type            string            `json:"type" yaml:"type"`             // Resource type (e.g., vm, storage, etc.)
	Properties      map[string]string `json:"properties" yaml:"properties"` // Generic properties for different resource types
	NestedResources []NestedResource  `json:"nestedResources,omitempty" yaml:"nestedResources,omitempty"`
}

type NestedResource struct {
	Type       string            `json:"type" yaml:"type"` // Nested resource type
	Properties map[string]string `json:"properties" yaml:"properties"`
}

// S3Bucket defines the structure for S3 bucket metadata
type S3Bucket struct {
	Name         string     `json:"name"`
	CreationDate *time.Time `json:"creationDate"`
}

func (c *Configuration) Validate() error {
	if c.Provider == "" {
		return fmt.Errorf("provider is required")
	}
	if len(c.Resources) == 0 {
		return fmt.Errorf("at least one resource must be defined")
	}
	for _, resourceMap := range c.Resources {
		for resourceType, resource := range resourceMap {
			if resourceType == "" {
				return fmt.Errorf("resource type is required")
			}
			if resource.Properties == nil || len(resource.Properties) == 0 {
				return fmt.Errorf("resource properties are required for type '%s'", resourceType)
			}
		}
	}
	return nil
}
