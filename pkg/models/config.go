package models

import (
	"fmt"
	"time"
)

type Configuration struct {
        Provider  string                `json:"provider" yaml:"provider"`
        Resources []map[string]Resource `json:"resources" yaml:"resources"`
}

type Resource struct {
	Type            string            `json:"type" yaml:"type"`             
	Properties      map[string]string `json:"properties" yaml:"properties"` 
	NestedResources []NestedResource  `json:"nestedResources,omitempty" yaml:"nestedResources,omitempty"`
}

type NestedResource struct {
	Type       string            `json:"type" yaml:"type"` 
	Properties map[string]string `json:"properties" yaml:"properties"`
}

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
