package commands

import (
	"cloudcrafter/pkg/providers"
	"cloudcrafter/pkg/services"
	"cloudcrafter/pkg/utils"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"github.com/urfave/cli/v2"
	"time"
)

// ListCommand returns the CLI command for listing resources
func ListCommand() *cli.Command {
	return &cli.Command{
		Name:  "list",
		Usage: "List provisioned resources for a provider",
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
		},
		Action: func(c *cli.Context) error {
			provider := c.String("provider")
			resourceType := c.String("resource")

			providerRegistry, err := providers.InitializeRegistry(provider)
			if err != nil {
				fmt.Printf("Error initializing provider: %s\n", err.Error())
			}

			// Create a progress bar
			bar := progressbar.NewOptions(100,
				progressbar.OptionSetDescription("Fetching resources..."),
				progressbar.OptionSetTheme(progressbar.Theme{
					Saucer:        "=",
					SaucerPadding: " ",
					BarStart:      "[",
					BarEnd:        "]",
				}),
			)

			// Simulate progress for the bar
			go func() {
				for i := 0; i <= 100; i += 10 {
					err := bar.Add(10)
					if err != nil {
						return
					}
					time.Sleep(100 * time.Millisecond) // Simulate work
				}
			}()

			provisioningService := services.NewProvisioningService(providerRegistry)
			// Fetch the list of resources
			resources, err := provisioningService.ListResources(provider, resourceType)
			if err != nil {
				return fmt.Errorf("failed to list resources: %w", err)
			}

			// Render the resources in a table
			if len(resources) == 0 {
				fmt.Println("No resources found.")
				return nil
			}

			headers := []string{"ID", "Name", "Type", "Region", "Status", "Created At"}
			rows := [][]string{}
			for _, resource := range resources {
				rows = append(rows, []string{
					resource.ID,
					resource.Name,
					resource.Type,
					resource.Region,
					resource.Status,
					resource.CreatedAt.Format("2006-01-02 15:04:05"),
				})
			}

			utils.RenderTable(headers, rows)
			return nil
		},
	}
}
