package services

import (
	"fmt"
	"multicloud-provisioner/pkg/models"
	"multicloud-provisioner/pkg/providers"
)

// ProvisioningService handles business logic for resource provisioning
type ProvisioningService struct {
	providerRegistry *providers.ProviderRegistry
}

// NewProvisioningService creates a new instance of ProvisioningService
func NewProvisioningService(providerRegistry *providers.ProviderRegistry) *ProvisioningService {
	return &ProvisioningService{
		providerRegistry: providerRegistry,
	}
}

// CreateResource provisions a new resource using the appropriate provider
func (s *ProvisioningService) CreateResource(config models.Configuration) ([]*models.ResourceMetadata, error) {
	// Validate the entire configuration
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	// Get the appropriate provider
	provider, err := s.providerRegistry.Get(config.Provider)
	if err != nil {
		return nil, fmt.Errorf("failed to find provider: %w", err)
	}

	// Iterate over resources and provision each
	var provisionedResources []*models.ResourceMetadata
	for _, resource := range config.Resources {
		resourceMetadata, err := provider.CreateResource(resource)
		if err != nil {
			return nil, fmt.Errorf("failed to provision resource '%s': %w", resource.Name, err)
		}
		provisionedResources = append(provisionedResources, resourceMetadata)
	}

	return provisionedResources, nil
}

// DeleteResource deletes a specific resource using its ID and provider
func (s *ProvisioningService) DeleteResource(providerName, resourceID string) error {
	// Get the appropriate provider
	provider, err := s.providerRegistry.Get(providerName)
	if err != nil {
		return fmt.Errorf("failed to find provider: %w", err)
	}

	// Delete the resource
	if err := provider.DeleteResource(resourceID); err != nil {
		return fmt.Errorf("failed to delete resource: %w", err)
	}

	return nil
}

// GetResource retrieves metadata about a specific resource
func (s *ProvisioningService) GetResource(providerName, resourceID string) (*models.ResourceMetadata, error) {
	// Get the appropriate provider
	provider, err := s.providerRegistry.Get(providerName)
	if err != nil {
		return nil, fmt.Errorf("failed to find provider: %w", err)
	}

	// Fetch resource metadata
	resourceMetadata, err := provider.GetResource(resourceID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve resource metadata: %w", err)
	}

	return resourceMetadata, nil
}
