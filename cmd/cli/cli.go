package main

import (
	"cloudcrafter/pkg/commands"
	"cloudcrafter/pkg/logger"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	//logger.InitLogger("development", zapcore.DebugLevel)
	defer logger.SyncLogger()

	fmt.Println("CloudCrafter CLI starting...")

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
		fmt.Println("Application terminated")
		_ = fmt.Errorf("%v", err)
	}
}
