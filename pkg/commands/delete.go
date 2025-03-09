package commands

import (
	"fmt"

	"github.com/Omotolani98/cloudcrafter/pkg/providers"
	"github.com/Omotolani98/cloudcrafter/pkg/services"

	"github.com/urfave/cli/v2"
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
				Name:    "resource",
				Aliases: []string{"r"},
				Usage:   "Resource type (e.g., storage, vm)",
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
			resourceType := c.String("resource")
			resourceID := c.String("id")

			providerRegistry, err := providers.InitializeRegistry(provider)
			if err != nil {
				return fmt.Errorf("error initializing provider registry: %w", err)
			}

			provisioningService := services.NewProvisioningService(providerRegistry)

			err = provisioningService.DeleteResource(provider, resourceType, resourceID)
			if err != nil {
				return fmt.Errorf("error deleting resource: %w", err)
			}

			fmt.Printf("Successfully deleted resource with ID %s from provider %s\n", resourceID, provider)
			return nil
		},
	}
}
