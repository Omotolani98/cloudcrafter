package models

import "time"

// ResourceMetadata defines metadata for a provisioned cloud resource
type ResourceMetadata struct {
	ID        string    `json:"id"`        // Unique identifier (e.g., UUID)
	Name      string    `json:"name"`      // Resource name
	Type      string    `json:"type"`      // Resource type (e.g., "vm", "storage")
	Provider  string    `json:"provider"`  // Cloud provider (e.g., AWS, Azure, GCP)
	Region    string    `json:"region"`    // Cloud region
	Status    string    `json:"status"`    // Current status (e.g., "running", "stopped")
	CreatedAt time.Time `json:"createdAt"` // Timestamp for when the resource was created
	UpdatedAt time.Time `json:"updatedAt"` // Timestamp for when the resource was last updated
}
