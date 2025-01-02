package main

import (
	"cloudcrafter/pkg/commands"
	"cloudcrafter/pkg/logger"
	"os"

	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	logger.InitLogger("development", zapcore.DebugLevel)
	defer logger.SyncLogger()

	logger.Log.Info("CloudCrafter CLI starting...")

	app := &cli.App{
		Name:  "CloudCrafter",
		Usage: "Provision and manage cloud resources across multiple providers",
		Commands: []*cli.Command{
			commands.ProvisionCommand(),
			commands.GenerateYAMLCommand(),
			commands.ListCommand(),
			commands.DeleteCommand(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Log.Fatal("Application terminated", zap.Error(err))
	}
}
