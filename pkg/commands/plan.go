package commands

import (
	"cloudcrafter/pkg/providers"
	"cloudcrafter/pkg/services"
	"cloudcrafter/pkg/utils"
	"fmt"

	"github.com/spf13/cobra"
)

func PlanCommand() *cobra.Command {
	var file string
	var estimate bool

	cmd := &cobra.Command{
		Use:   "plan",
		Short: "Plan resources (YAML-based with optional cost estimation)",
		RunE: func(cmd *cobra.Command, args []string) error {
			return planFromYAML(file, estimate)
		},
	}

	cmd.Flags().StringVarP(&file, "file", "f", "", "Path to the YAML configuration file")
	cmd.Flags().BoolVarP(&estimate, "estimate-cost", "e", false, "Estimate the cost of the resources")

	return cmd
}
func planFromYAML(filePath string, estimateCost bool) error {
        config, err := utils.ParseYAML(filePath)
        if err != nil {
                return fmt.Errorf("failed to parse YAML: %v", err)
        }

        providerRegistry, err := providers.InitializeRegistry(config.Provider)
        if err != nil {
                return fmt.Errorf("failed to initialize provider registry: %v", err)
        }

        estimatorService := services.NewEstimatorService(providerRegistry)

        if estimateCost {
                totalCost, err := estimatorService.EstimateCosts(config)
                if err != nil {
                        return fmt.Errorf("failed to estimate costs: %v", err)
                }

		fmt.Printf("Estimated Cost: $%.2f/month\n", totalCost)
		return nil
	}
        metadata, err := estimatorService.EstimateCosts(config)
        if err != nil {
                return fmt.Errorf("failed to plan resources: %v", err)
        }

	fmt.Printf("Planned resources: %+v\n", metadata)

	return nil
}
