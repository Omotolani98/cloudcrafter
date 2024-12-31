package commands

import (
	"cloudcrafter/pkg/logger"
	"cloudcrafter/pkg/providers"
	"cloudcrafter/pkg/services"
	"fmt"

	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func ListCommand() *cli.Command {
	return &cli.Command{
		Name:  "list",
		Usage: "List all provisioned resources for a specific provider",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "provider",
				Aliases:  []string{"p"},
				Usage:    "Cloud provider (e.g., aws, azure, gcp)",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			provider := c.String("provider")

			// Initialize ProviderRegistry based on the provider context
			providerRegistry, err := providers.InitializeRegistry(provider)
			if err != nil {
				return fmt.Errorf("error initializing provider registry: %w", err)
			}

			// Initialize Provisioning Service
			provisioningService := services.NewProvisioningService(providerRegistry)

			// List resources
			resources, err := provisioningService.ListResources(provider)
			if err != nil {
				return fmt.Errorf("error listing resources: %w", err)
			}

			fmt.Printf("Resources for provider %s:\n", provider)
			for _, resource := range resources {
				logger.Log.Info("Resource discovered",
					zap.String("id", resource.ID),
					zap.String("name", resource.Name),
					zap.String("status", resource.Status),
					zap.String("region", resource.Region),
				)
			}
			return nil
		},
	}
}
