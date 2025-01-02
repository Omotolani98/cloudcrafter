package commands

import (
	"cloudcrafter/pkg/providers"
	"cloudcrafter/pkg/services"
	"cloudcrafter/pkg/utils"
	"fmt"
	"github.com/urfave/cli/v2"
)

// ProvisionCommand returns the CLI command for provisioning
func ProvisionCommand() *cli.Command {
	return &cli.Command{
		Name:  "provision",
		Usage: "Provision resources (YAML-based or interactive)",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "file",
				Usage: "Path to the YAML configuration file",
			},
			&cli.StringFlag{
				Name:  "provider",
				Usage: "Provider to use",
			},
		},
		Action: func(c *cli.Context) error {
			file := c.String("file")
			provider := c.String("provider")

			if file != "" {
				return provisionFromYAML(file, provider)
			}
			return provisionInteractive()
		},
	}
}

// YAML-based provisioning
func provisionFromYAML(filePath string, providerName string) error {
	config, err := utils.ParseYAML(filePath)
	if err != nil {
		return fmt.Errorf("failed to parse YAML: %v", err)
	}

	providerRegistry, err := providers.InitializeRegistry(providerName)
	if err != nil {
		return fmt.Errorf("failed to initialize provider registry: %v", err)
	}

	provisionService := services.NewProvisioningService(providerRegistry)
	metadata, err := provisionService.CreateResource(config)
	if err != nil {
		return fmt.Errorf("failed to provision resource: %v", err)
	}
	//provider, err := providers.NewAWSProvider(config.Provider)
	//if err != nil {
	//	return fmt.Errorf("failed to initialize AWS provider: %v", err)
	//}
	//
	//for _, resource := range config.Resources {
	//	metadata, err := provider.CreateResource(resource)
	//	if err != nil {
	//		return fmt.Errorf("failed to provision resource %s: %v", resource.Name, err)
	//	}
	//	fmt.Printf("Provisioned resource: %+v\n", metadata)
	//}

	fmt.Printf("Provisioned resource: %+v\n", metadata)

	return nil
}

// Interactive provisioning
func provisionInteractive() error {
	resource, err := utils.CollectInteractiveResourceData()
	if err != nil {
		return err
	}

	provider, err := providers.NewAWSProvider(resource.Region)
	if err != nil {
		fmt.Printf("Failed to initialize AWS provider: %v\n", err)
	}
	metadata, err := provider.CreateResource(resource)
	if err != nil {
		return fmt.Errorf("failed to provision resource interactively: %v", err)
	}
	fmt.Printf("Provisioned resource: %+v\n", metadata)

	return nil
}
