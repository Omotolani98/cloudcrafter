package commands

import (
	"fmt"

	"github.com/Omotolani98/cloudcrafter/pkg/providers"
	"github.com/Omotolani98/cloudcrafter/pkg/services"
	"github.com/Omotolani98/cloudcrafter/pkg/utils"
	"github.com/urfave/cli/v2"
)

// ProvisionCommand returns the CLI command for provisioning
func ProvisionCommand() *cli.Command {
	return &cli.Command{
		Name:  "provision",
		Usage: "Provision resources (YAML-based or interactive)",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "file",
				Aliases: []string{"f"},
				Usage:   "Path to the YAML configuration file",
			},
		},
		Action: func(c *cli.Context) error {
			file := c.String("file")
			return provisionFromYAML(file)
		},
	}
}

func provisionFromYAML(filePath string) error {
	config, err := utils.ParseYAML(filePath)
	if err != nil {
		return fmt.Errorf("failed to parse YAML: %v", err)
	}

	providerRegistry, err := providers.InitializeRegistry(config.Provider)
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
