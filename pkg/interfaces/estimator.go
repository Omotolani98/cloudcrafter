package providers

type CostEstimator interface {
	EstimateVMCost(resource *map[string]string) (float64, error)
	EstimateStorageCost(resource *map[string]string) (float64, error)
	EstimateDatabasesCost(resource *map[string]string) (float64, error)
}
