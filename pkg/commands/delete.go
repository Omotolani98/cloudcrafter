package commands

import (
	"cloudcrafter/pkg/providers"
	"cloudcrafter/pkg/services"
	"fmt"

	"github.com/spf13/cobra"
)

func DeleteCommand() *cobra.Command {
	var provider string
	var resourceType string
	var resourceID string

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a resource by provider and ID",
		RunE: func(cmd *cobra.Command, args []string) error {
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

	cmd.Flags().StringVarP(&provider, "provider", "p", "", "Cloud provider (e.g., aws, azure, gcp)")
	cmd.Flags().StringVarP(&resourceType, "resource", "r", "", "Resource type (e.g., storage, vm)")
	cmd.Flags().StringVarP(&resourceID, "id", "i", "", "Resource ID")
	_ = cmd.MarkFlagRequired("provider")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}
