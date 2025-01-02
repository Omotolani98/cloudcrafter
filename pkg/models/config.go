package models

import "errors"

// Resource defines a single cloud resource configuration
type Resource struct {
	Name           string   `json:"name" yaml:"name"`
	MachineType    string   `json:"machineType" yaml:"machineType"`
	Region         string   `json:"region" yaml:"region"`
	Image          string   `json:"image" yaml:"image"`
	Subnet         string   `json:"subnet" yaml:"subnet"`
	SecurityGroups []string `json:"securityGroups" yaml:"securityGroups"`
	KeyName        string   `json:"keyName" yaml:"keyName"`
}

// Configuration defines the overall user-provided infrastructure configuration
type Configuration struct {
	Provider  string     `json:"provider" yaml:"provider"` // Cloud provider (e.g., AWS, Azure, GCP)
	Resources []Resource `json:"resources" yaml:"resources"`
}

func (c *Configuration) Validate() error {
	if c.Provider == "" {
		return errors.New("provider cannot be empty")
	}

	if len(c.Resources) == 0 {
		return errors.New("at least one resource must be specified")
	}

	// Validate each resource
	for _, resource := range c.Resources {
		if err := ValidateResourceConfig(resource); err != nil {
			return err
		}
	}

	return nil
}
