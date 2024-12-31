package commands

import (
	"cloudcrafter/pkg/logger"
	"cloudcrafter/pkg/providers"
	"cloudcrafter/pkg/services"
	"cloudcrafter/pkg/utils"
	"fmt"

	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func ProvisionCommand() *cli.Command {
	return &cli.Command{
		Name:  "provision",
		Usage: "Provision resources using a YAML file",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "file",
				Aliases:  []string{"f"},
				Usage:    "Path to the YAML configuration file",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			filePath := c.String("file")

			logger.Log.Info("Executing 'provision' command", zap.String("file", filePath))

			// Parse the YAML file
			config, err := utils.ParseYAMLConfig(filePath)
			if err != nil {
				logger.Log.Error("Error parsing YAML file", zap.Error(err))
				return fmt.Errorf("error parsing YAML file: %w", err)
			}

			logger.Log.Info("YAML file parsed successfully", zap.String("provider", config.Provider))

			// Initialize ProviderRegistry based on the provider context in the YAML
			providerRegistry, err := providers.InitializeRegistry(config.Provider)
			if err != nil {
				logger.Log.Error("Error initializing provider registry", zap.Error(err))
				return fmt.Errorf("error initializing provider registry: %w", err)
			}

			// Initialize Provisioning Service
			provisioningService := services.NewProvisioningService(providerRegistry)

			// Provision resources
			provisionedResources, err := provisioningService.CreateResource(*config)
			if err != nil {
				logger.Log.Error("Error provisioning resources", zap.Error(err))
				return fmt.Errorf("error provisioning resources: %w", err)
			}

			logger.Log.Info("Resources provisioned successfully",
				zap.String("provider", config.Provider),
				zap.Int("count", len(provisionedResources)),
			)

			fmt.Println("Successfully provisioned the following resources:")
			for _, resource := range provisionedResources {
				fmt.Printf("ID: %s, Name: %s, Type: %s, Region: %s\n",
					resource.ID, resource.Name, resource.Type, resource.Region)
			}

			return nil
		},
	}
}
