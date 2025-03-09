package models

import (
	"fmt"
	"time"
)

// NetworkingConfig represents a generic networking configuration for cloud providers.
type NetworkingConfig struct {
	VPC struct {
		CIDRBlock string `json:"cidrBlock"`
	} `json:"vpc"`
	Subnet struct {
		CIDRBlock string `json:"cidrBlock"`
	} `json:"subnet"`
	SecurityGroup struct {
		Description  string        `json:"description"`
		IngressRules []IngressRule `json:"ingressRules"`
	} `json:"securityGroup"`
}

// IngressRule represents a rule for allowing inbound traffic.
type IngressRule struct {
	Protocol string `json:"protocol"`
	FromPort int    `json:"fromPort"`
	ToPort   int    `json:"toPort"`
	CIDRIP   string `json:"cidrIp"`
}

// Configuration represents the overall configuration including networking.
type Configuration struct {
	Provider   string                `json:"provider"`
	Networking NetworkingConfig      `json:"networking"`
	Resources  []map[string]Resource `json:"resources"`
}

// Resource represents a cloud resource.
type Resource struct {
	Type       string            `json:"type"`
	Properties map[string]string `json:"properties"`
}

// Resource defines a single cloud resource configuration
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
