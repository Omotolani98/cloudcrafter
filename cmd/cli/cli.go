package main

import (
	"cloudcrafter/pkg/commands"
	"cloudcrafter/pkg/logger"
	"context"
	"os"

	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"
)

func main() {
	defer logger.SyncLogger()

	rootCmd := &cobra.Command{
		Use:   "cloudcrafter",
		Short: "Provision and manage cloud resources across multiple providers",
	}

	rootCmd.AddCommand(
		commands.ProvisionCommand(),
		commands.GenerateYAMLCommand(),
		commands.ListCommand(),
		commands.DeleteCommand(),
		commands.PlanCommand(),
	)

	if err := fang.Execute(context.Background(), rootCmd); err != nil {
		os.Exit(1)
	}
}
