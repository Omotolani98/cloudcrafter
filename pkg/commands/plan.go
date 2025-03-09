package commands

import (
	"fmt"

	"github.com/Omotolani98/cloudcrafter/pkg/providers"
	"github.com/Omotolani98/cloudcrafter/pkg/services"
	"github.com/Omotolani98/cloudcrafter/pkg/utils"
	"github.com/urfave/cli/v2"
)

// PlanCommand returns the CLI command for planning resources
func PlanCommand() *cli.Command {
	return &cli.Command{
		Name:  "plan",
		Usage: "Plan resources (YAML-based with optional cost estimation)",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "file",
				Aliases: []string{"f"},
				Usage:   "Path to the YAML configuration file",
			},
			&cli.BoolFlag{
				Name:    "estimate-cost",
				Aliases: []string{"e"},
				Usage:   "Estimate the cost of the resources",
			},
			//&cli.StringFlag{
			//	Name:    "provider",
			//	Aliases: []string{"p"},
			//	Usage:   "Provider to use",
			//},
		},
		Action: func(c *cli.Context) error {
			file := c.String("file")
			estimateCost := c.Bool("estimate-cost")

			// Call the plan logic
			return planFromYAML(file, estimateCost)
		},
	}
}

func planFromYAML(filePath string, estimateCost bool) error {
	// Parse the YAML configuration
	config, err := utils.ParseYAML(filePath)
	if err != nil {
		return fmt.Errorf("failed to parse YAML: %v", err)
	}

	// Initialize the provider registry
	providerRegistry, err := providers.InitializeRegistry(config.Provider)
	if err != nil {
		return fmt.Errorf("failed to initialize provider registry: %v", err)
	}

	// Initialize the provisioning service
	estimatorService := services.NewEstimatorService(providerRegistry)

	// If cost estimation is requested
	if estimateCost {
		totalCost, err := estimatorService.EstimateCosts(config)
		if err != nil {
			return fmt.Errorf("failed to estimate costs: %v", err)
		}

		fmt.Printf("Estimated Cost: $%.2f/month\n", totalCost)
		return nil
	}

	// If no cost estimation, just plan the resources
	metadata, err := estimatorService.EstimateCosts(config)
	if err != nil {
		return fmt.Errorf("failed to plan resources: %v", err)
	}

	fmt.Printf("Planned resources: %+v\n", metadata)

	return nil
}
