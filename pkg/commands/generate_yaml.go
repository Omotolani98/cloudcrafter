package commands

import (
	"cloudcrafter/pkg/utils"
	"fmt"

	"github.com/spf13/cobra"
)

func GenerateYAMLCommand() *cobra.Command {
	var output string

	cmd := &cobra.Command{
		Use:   "generate-yaml",
		Short: "Generate a YAML file interactively",
		RunE: func(cmd *cobra.Command, args []string) error {
			resource, err := utils.CollectInteractiveResourceData()
			if err != nil {
				return fmt.Errorf("failed to collect resource data: %v", err)
			}

			if err := utils.WriteYaml(resource, output); err != nil {
				return fmt.Errorf("failed to generate YAML: %v", err)
			}

			fmt.Printf("YAML file generated successfully: %s\n", output)
			return nil
		},
	}

	cmd.Flags().StringVarP(&output, "output", "o", "", "Output file path")

	return cmd
}
