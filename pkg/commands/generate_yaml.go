package commands

import (
	"cloudcrafter/pkg/utils"
	"fmt"
	"github.com/urfave/cli/v2"
)

// GenerateYAMLCommand returns the CLI command for generating YAML
func GenerateYAMLCommand() *cli.Command {
	return &cli.Command{
		Name:  "generate-yaml",
		Usage: "Generate a YAML file interactively",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "Output file path",
			},
		},
		Action: func(c *cli.Context) error {
			resource, err := utils.CollectInteractiveResourceData()
			if err != nil {
				return fmt.Errorf("failed to collect resource data: %v", err)
			}

			outputPath := c.String("output")

			err = utils.WriteYaml(resource, outputPath)
			if err != nil {
				return fmt.Errorf("failed to generate YAML: %v", err)
			}

			fmt.Printf("YAML file generated successfully: %s\n", outputPath)
			return nil
		},
	}
}
