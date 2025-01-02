package commands

import (
	"cloudcrafter/pkg/utils"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

// GenerateYAMLCommand returns the CLI command for generating YAML
func GenerateYAMLCommand() *cli.Command {
	return &cli.Command{
		Name:  "generate-yaml",
		Usage: "Generate a YAML file interactively",
		Action: func(c *cli.Context) error {
			resource, err := utils.CollectInteractiveResourceData()
			if err != nil {
				return err
			}

			yamlData, err := utils.GenerateYAML(resource)
			if err != nil {
				return fmt.Errorf("failed to generate YAML: %v", err)
			}

			err = os.WriteFile("generated.yaml", yamlData, 0644)
			if err != nil {
				return fmt.Errorf("failed to save YAML file: %v", err)
			}
			fmt.Println("YAML file generated successfully: generated.yaml")
			return nil
		},
	}
}
