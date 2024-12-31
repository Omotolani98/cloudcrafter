package models

// Resource defines a single cloud resource configuration
type Resource struct {
	Type        string `json:"type" yaml:"type"`                                   // e.g., "vm", "storage"
	Name        string `json:"name" yaml:"name"`                                   // Resource name
	MachineType string `json:"machineType,omitempty" yaml:"machineType,omitempty"` // For VMs
	Region      string `json:"region" yaml:"region"`                               // Cloud region
	Image       string `json:"image,omitempty" yaml:"image,omitempty"`             // OS image (for VMs)
}

// Configuration defines the overall user-provided infrastructure configuration
type Configuration struct {
	Provider  string     `json:"provider" yaml:"provider"` // Cloud provider (e.g., AWS, Azure, GCP)
	Resources []Resource `json:"resources" yaml:"resources"`
}
