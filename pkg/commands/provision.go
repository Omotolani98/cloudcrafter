package commands

import (
	"cloudcrafter/pkg/providers"
	"cloudcrafter/pkg/services"
	"cloudcrafter/pkg/utils"
	"fmt"

	"github.com/spf13/cobra"
)

func ProvisionCommand() *cobra.Command {
	var file string
	var provider string

	cmd := &cobra.Command{
		Use:   "provision",
		Short: "Provision resources (YAML-based or interactive)",
		RunE: func(cmd *cobra.Command, args []string) error {
			return provisionFromYAML(file, provider)
		},
	}

	cmd.Flags().StringVarP(&file, "file", "f", "", "Path to the YAML configuration file")
	cmd.Flags().StringVarP(&provider, "provider", "p", "", "Provider to use")

	return cmd
}

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

	fmt.Printf("Provisioned resource: %+v\n", metadata)

	return nil
}
