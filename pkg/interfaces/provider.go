package providers

import "cloudcrafter/pkg/models"

// import "multi-cloud-provisioner/pkg/models"

// Provider defines the methods that every cloud provider must implement
type Provider interface {
	CreateResource(resource models.Resource) (*models.ResourceMetadata, error)
	DeleteResource(resourceID string) error
	GetResource(resourceID string) (*models.ResourceMetadata, error)
}
