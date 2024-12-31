package providers

import (
	providers "cloudcrafter/pkg/interfaces"
	"cloudcrafter/pkg/logger"
	"fmt"

	"go.uber.org/zap"
)

type ProviderRegistry struct {
	providers map[string]providers.Provider
}

func NewProviderRegistry() *ProviderRegistry {
	return &ProviderRegistry{
		providers: make(map[string]providers.Provider),
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
	logger.Log.Info("Provider retrieved", zap.String("provider", providerName))
	return provider, nil
}

// InitializeRegistry initializes the ProviderRegistry based on the specified provider context
func InitializeRegistry(provider string) (*ProviderRegistry, error) {
	logger.Log.Info("Initializing provider registry", zap.String("provider", provider))

	registry := NewProviderRegistry()

	switch provider {
	case "aws":
		awsProvider, err := NewAWSProvider("us-east-1") // Region can be made dynamic
		if err != nil {
			logger.Log.Error("Failed to initialize AWS provider", zap.Error(err))
			return nil, fmt.Errorf("failed to initialize AWS provider: %w", err)
		}
		registry.Register("aws", awsProvider)
		logger.Log.Info("AWS provider registered", zap.String("region", "us-east-1"))
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
		logger.Log.Error("Unsupported provider specified", zap.String("provider", provider))
		return nil, fmt.Errorf("unsupported provider: %s", provider)
	}

	return registry, nil
}
