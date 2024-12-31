package models

import "errors"

// Validate validates the Configuration model
func (config *Configuration) Validate() error {
	if config.Provider == "" {
		return errors.New("provider is required")
	}

	if len(config.Resources) == 0 {
		return errors.New("at least one resource must be defined")
	}

	for _, resource := range config.Resources {
		if resource.Type == "" || resource.Name == "" || resource.Region == "" {
			return errors.New("resource type, name, and region are required")
		}

		// Additional validation for VMs
		if resource.Type == "vm" && resource.MachineType == "" {
			return errors.New("machineType is required for VM resources")
		}
	}

	return nil
}
