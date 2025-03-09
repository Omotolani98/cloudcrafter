package providers

import (
	"fmt"

	providers "github.com/Omotolani98/cloudcrafter/pkg/interfaces"
	"github.com/Omotolani98/cloudcrafter/pkg/logger"

	"go.uber.org/zap"
)

type ProviderRegistry struct {
	providers      map[string]providers.Provider
	CostEstimators map[string]providers.CostEstimator
}

func NewProviderRegistry() *ProviderRegistry {
	return &ProviderRegistry{
		providers:      make(map[string]providers.Provider),
		CostEstimators: make(map[string]providers.CostEstimator),
	}
}

func (r *ProviderRegistry) Register(providerName string, provider providers.Provider) {
	r.providers[providerName] = provider
}

func (r *ProviderRegistry) Get(providerName string) (providers.Provider, error) {
	provider, exists := r.providers[providerName]
	if !exists {
		availableProviders := make([]string, 0, len(r.providers))
		for k := range r.providers {
			availableProviders = append(availableProviders, k)
		}
		logger.Log.Error("Provider not found",
			zap.String("requested_provider", providerName),
			zap.Strings("available_providers", availableProviders),
		)
		return nil, fmt.Errorf("provider %s not registered", providerName)
	}
	fmt.Printf("Provider retrieved: %s\n", providerName)
	return provider, nil
}

func (r *ProviderRegistry) RegisterCostEstimator(providerName string, estimator providers.CostEstimator) {
	r.CostEstimators[providerName] = estimator
}

func (r *ProviderRegistry) GetCostEstimator(providerName string) (providers.CostEstimator, error) {
	estimator, ok := r.CostEstimators[providerName]
	if !ok {
		return nil, fmt.Errorf("cost estimator for provider %s not found", providerName)
	}
	return estimator, nil
}

// InitializeRegistry initializes the ProviderRegistry based on the specified provider context
func InitializeRegistry(provider string) (*ProviderRegistry, error) {
	fmt.Printf("Initializing provider registry: %s\n", provider)

	registry := NewProviderRegistry()

	switch provider {
	case "aws":
		awsProvider, err := NewAWSProvider("us-east-1") // Region can be made dynamic
		if err != nil {
			//logger.Log.Error("Failed to initialize AWS provider", zap.Error(err))
			return nil, fmt.Errorf("failed to initialize AWS provider: %w", err)
		}

		registry.RegisterCostEstimator("aws", awsProvider)
		// Register AWS provider
		registry.Register("aws", awsProvider)
		fmt.Printf("AWS provider registered \n")
	case "azure":
		// Add Azure provider initialization logic here
		// azureProvider, err := NewAzureProvider()
		// if err != nil {
		//     logger.Log.Error("Failed to initialize Azure provider", zap.Error(err))
		//     return nil, fmt.Errorf("failed to initialize Azure provider: %w", err)
		// }
		// registry.Register("azure", azureProvider)
	case "gcp":
		// Add GCP provider initialization logic here
		// gcpProvider, err := NewGCPProvider()
		// if err != nil {
		//     logger.Log.Error("Failed to initialize GCP provider", zap.Error(err))
		//     return nil, fmt.Errorf("failed to initialize GCP provider: %w", err)
		// }
		// registry.Register("gcp", gcpProvider)
	default:
		//logger.Log.Error("Unsupported provider specified", zap.String("provider", provider))
		return nil, fmt.Errorf("unsupported provider: %s", provider)
	}

	return registry, nil
}
