package services

import (
	"cloudcrafter/pkg/models"
	"cloudcrafter/pkg/providers"
	"fmt"
)

type EstimatorService struct {
	providerRegistry *providers.ProviderRegistry
}

func NewEstimatorService(providerRegistry *providers.ProviderRegistry) *EstimatorService {
	return &EstimatorService{
		providerRegistry: providerRegistry,
	}
}

func (s *EstimatorService) EstimateCosts(config *models.Configuration) (float64, error) {
	
	fmt.Printf("Available cost estimators: %+v\n", s.providerRegistry.CostEstimators)

	
	estimator, err := s.providerRegistry.GetCostEstimator(config.Provider)
	if estimator == nil {
		fmt.Println("Error: No cost estimator found for provider", config.Provider)
		return 0, fmt.Errorf("no cost estimator found for provider %s", config.Provider)
	}
	if err != nil {
		return 0, fmt.Errorf("failed to find cost estimator: %w", err)
	}

	var totalCost float64

	for _, resourceMap := range config.Resources {
		for resourceType, resource := range resourceMap {
			fmt.Printf("Processing resource type: %s, properties: %+v\n", resourceType, resource.Properties)

			var cost float64
			var err error
			switch resourceType {
			case "vm":
				cost, err = estimator.EstimateVMCost(&resource.Properties)
			case "storage":
				cost, err = estimator.EstimateStorageCost(&resource.Properties)
			case "database":
				cost, err = estimator.EstimateDatabasesCost(&resource.Properties)
			default:
				return 0, fmt.Errorf("unsupported resource type: %s", resourceType)
			}

			if err != nil {
				return 0, fmt.Errorf("failed to estimate cost for resource type %s: %w", resourceType, err)
			}

			totalCost += cost
		}
	}

	fmt.Printf("Total estimated cost: $%.2f\n", totalCost)
	return totalCost, nil
}
