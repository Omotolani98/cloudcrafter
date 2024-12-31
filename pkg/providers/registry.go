package providers

import (
	providers "cloudcrafter/pkg/interfaces"
	"fmt"
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
		return nil, fmt.Errorf("provider %s not registered", providerName)
	}
	return provider, nil
}
