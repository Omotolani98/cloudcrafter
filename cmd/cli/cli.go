package main

import (
	"fmt"
	"os"

	"github.com/Omotolani98/cloudcrafter/pkg/commands"
	"github.com/Omotolani98/cloudcrafter/pkg/logger"

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
			commands.PlanCommand(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("Application terminated")
		_ = fmt.Errorf("%v", err)
	}
}
