package commands

import (
	"cloudcrafter/pkg/providers"
	"cloudcrafter/pkg/services"
	"cloudcrafter/pkg/utils"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"time"

	"github.com/spf13/cobra"
)

func ListCommand() *cobra.Command {
	var provider string
	var resourceType string

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List provisioned resources for a provider",
		RunE: func(cmd *cobra.Command, args []string) error {
			providerRegistry, err := providers.InitializeRegistry(provider)
			if err != nil {
				fmt.Printf("Error initializing provider: %s\n", err.Error())
			}

			bar := progressbar.NewOptions(100,
				progressbar.OptionSetDescription("Fetching resources..."),
				progressbar.OptionSetTheme(progressbar.Theme{
					Saucer:        "=",
					SaucerPadding: " ",
					BarStart:      "[",
					BarEnd:        "]",
				}),
			)

			go func() {
				for i := 0; i <= 100; i += 10 {
					if err := bar.Add(10); err != nil {
						return
					}
					time.Sleep(100 * time.Millisecond)
				}
			}()

			provisioningService := services.NewProvisioningService(providerRegistry)
			resources, err := provisioningService.ListResources(provider, resourceType)
			if err != nil {
				return fmt.Errorf("failed to list resources: %w", err)
			}

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

	cmd.Flags().StringVarP(&provider, "provider", "p", "", "Cloud provider (e.g., aws, azure, gcp)")
	cmd.Flags().StringVarP(&resourceType, "resource", "r", "", "Resource type (e.g., storage, vm)")
	_ = cmd.MarkFlagRequired("provider")

	return cmd
}
