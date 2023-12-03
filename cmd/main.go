package main

import (
	"fmt"
	"github.com/theghostmac/web3trakka/internal/housekeeper"
	"github.com/theghostmac/web3trakka/internal/runner"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	// Initialize the runner.
	startRunner := runner.NewStartRunner()

	// initialize the housekeeper.
	logger := housekeeper.NewCustomLogger()

	cmd := &cli.Command{
		Name:         "start",
		Aliases:      nil,
		Usage:        "starts the web3trakka server",
		UsageText:    "",
		Description:  "",
		ArgsUsage:    "",
		Category:     "",
		BashComplete: nil,
		Before:       nil,
		After:        nil,
		Action: func(c *cli.Context) error {
			startRunner.StartServer()
			return nil
		},
		OnUsageError:           nil,
		Subcommands:            nil,
		Flags:                  nil,
		SkipFlagParsing:        false,
		HideHelp:               false,
		HideHelpCommand:        false,
		Hidden:                 false,
		UseShortOptionHandling: false,
		HelpName:               "",
		CustomHelpTemplate:     "",
	}

	// Initialize the application.
	app := &cli.App{
		Commands: []*cli.Command{cmd},
	}

	err := app.Run(os.Args)
	if err != nil {
		errMsg := fmt.Sprintf("error: %s", err)
		logger.Fatal(errMsg)
		os.Exit(1)
	}
}
