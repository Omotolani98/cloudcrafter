package services

import (
	"cloudcrafter/pkg/models"
	"cloudcrafter/pkg/providers"
	"fmt"
	"time"
)

type ProvisioningService struct {
	providerRegistry *providers.ProviderRegistry
}

func NewProvisioningService(providerRegistry *providers.ProviderRegistry) *ProvisioningService {
	return &ProvisioningService{
		providerRegistry: providerRegistry,
	}
}

func (s *ProvisioningService) CreateResource(config *models.Configuration) ([]*models.ResourceMetadata, error) {
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	provider, err := s.providerRegistry.Get(config.Provider)
	if err != nil {
		return nil, fmt.Errorf("failed to find provider: %w", err)
	}

	var provisionedResources []*models.ResourceMetadata

	for _, resourceMap := range config.Resources {
		for resourceType, resource := range resourceMap {
			switch resourceType {
			case "vm":
				resourceMetadata, err := provider.CreateResource(resource)
				if err != nil {
					return nil, fmt.Errorf("failed to provision VM resource: %w", err)
				}
				provisionedResources = append(provisionedResources, resourceMetadata)
			case "storage":
				bucketName, ok := resource.Properties["bucketName"]
				if !ok || bucketName == "" {
					return nil, fmt.Errorf("missing required property: bucketName for storage resource")
				}
				err := provider.CreateBucket(bucketName)
				if err != nil {
					return nil, fmt.Errorf("failed to create bucket: %w", err)
				}
				provisionedResources = append(provisionedResources, &models.ResourceMetadata{
					ID:        bucketName, 
					Name:      bucketName,
					Type:      "storage",
					Provider:  config.Provider,
					CreatedAt: time.Now(),
				})
			default:
				return nil, fmt.Errorf("unsupported resource type: %s", resourceType)
			}
		}
	}

	logProvisionedResources(provisionedResources)
	return provisionedResources, nil
}

func logProvisionedResources(resources []*models.ResourceMetadata) {
	fmt.Println("Provisioned resources:")
	for _, resource := range resources {
		fmt.Printf("Resource: %+v\n", *resource) 
	}
}

func (s *ProvisioningService) DeleteResource(providerName, resourceType string, resourceID string) error {
	provider, err := s.providerRegistry.Get(providerName)
	if err != nil {
		return fmt.Errorf("failed to find provider: %w", err)
	}

	switch resourceType {
	case "vm":
		return provider.DeleteResource(resourceID)
	case "storage":
		return provider.DeleteBucket(resourceID)
	default:
		return fmt.Errorf("unsupported resource type: %s", resourceType)
	}
}

func (s *ProvisioningService) GetResource(providerName, resourceID string) (*models.ResourceMetadata, error) {
	
	provider, err := s.providerRegistry.Get(providerName)
	if err != nil {
		return nil, fmt.Errorf("failed to find provider: %w", err)
	}

	
	resourceMetadata, err := provider.GetResource(resourceID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve resource metadata: %w", err)
	}

	return resourceMetadata, nil
}

func (s *ProvisioningService) ListResources(providerName string, resourceType string) ([]models.ResourceMetadata, error) {
	provider, err := s.providerRegistry.Get(providerName)
	if err != nil {
		return nil, fmt.Errorf("provider %s not found: %w", providerName, err)
	}

	switch resourceType {
	case "vm":
		return provider.ListResources()
	case "storage":
		buckets, err := provider.ListBuckets()
		if err != nil {
			return nil, err
		}

		var resources []models.ResourceMetadata
		for _, bucket := range buckets {
			resources = append(resources, models.ResourceMetadata{
				ID:        bucket.Name,
				Name:      bucket.Name,
				Type:      "storage",
				Provider:  providerName,
				CreatedAt: *bucket.CreationDate,
			})
		}
		return resources, nil
	default:
		return nil, fmt.Errorf("unsupported resource type: %s", resourceType)
	}
}

