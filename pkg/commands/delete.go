package commands

import (
	"cloudcrafter/pkg/logger"
	"cloudcrafter/pkg/providers"
	"cloudcrafter/pkg/services"
	"fmt"

	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func DeleteCommand() *cli.Command {
	return &cli.Command{
		Name:  "delete",
		Usage: "Delete a resource by provider and ID",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "provider",
				Aliases:  []string{"p"},
				Usage:    "Cloud provider (e.g., aws, azure, gcp)",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "id",
				Aliases:  []string{"i"},
				Usage:    "Resource ID",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			provider := c.String("provider")
			resourceID := c.String("id")

			logger.Log.Info("Executing 'delete' command",
				zap.String("provider", provider),
				zap.String("resource_id", resourceID),
			)

			// Initialize ProviderRegistry
			providerRegistry, err := providers.InitializeRegistry(provider)
			if err != nil {
				logger.Log.Error("Error initializing provider registry", zap.Error(err))
				return fmt.Errorf("error initializing provider registry: %w", err)
			}

			// Initialize Provisioning Service
			provisioningService := services.NewProvisioningService(providerRegistry)

			// Delete the resource
			err = provisioningService.DeleteResource(provider, resourceID)
			if err != nil {
				logger.Log.Error("Error deleting resource", zap.Error(err))
				return fmt.Errorf("error deleting resource: %w", err)
			}

			logger.Log.Info("Resource deleted successfully",
				zap.String("provider", provider),
				zap.String("resource_id", resourceID),
			)

			fmt.Printf("Successfully deleted resource with ID %s from provider %s\n", resourceID, provider)
			return nil
		},
	}
}
