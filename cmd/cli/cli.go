package main

import (
	"cloudcrafter/pkg/services"
	"cloudcrafter/pkg/utils"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "CloudCrafter",
		Usage: "Provision and manage cloud resources across multiple providers",
		Commands: []*cli.Command{
			{
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

					// Parse YAML configuration
					config, err := utils.ParseYAMLConfig(filePath)
					if err != nil {
						return fmt.Errorf("error parsing YAML file: %w", err)
					}

					// Initialize provisioning service
					provisioningService := services.NewProvisioningService(nil) // Add provider registry
					provisionedResources, err := provisioningService.CreateResource(*config)
					if err != nil {
						return fmt.Errorf("error provisioning resources: %w", err)
					}

					fmt.Printf("Successfully provisioned resources: %+v\n", provisionedResources)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
