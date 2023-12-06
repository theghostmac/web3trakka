package main

import (
	"fmt"
	"github.com/theghostmac/web3trakka/internal/housekeeper"
	"github.com/theghostmac/web3trakka/internal/runner"
	"github.com/theghostmac/web3trakka/internal/web3trakka"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	// Initialize the runner.
	startRunner := runner.NewStartRunner()

	// initialize the housekeeper.
	logger := housekeeper.NewCustomLogger()

	trackCrypto := web3trakka.NewCryptoTracker()
	viewPortfolio := web3trakka.NewPortfolioViewer()
	setAlert := web3trakka.NewAlertSetter()

	// Define web3trakka
	commands := []*cli.Command{
		{
			Name:  "start",
			Usage: "Starts the web3trakka server",
			Action: func(c *cli.Context) error {
				startRunner.StartServer()
				return nil
			},
		},
		{
			Name:  "track",
			Usage: "Track a new cryptocurrency",
			Action: func(c *cli.Context) error {
				cryptoName := c.Args().First() // Gets the first argument.
				if cryptoName == "" {
					// TODO: write a proper response message for wrong argument formats, raising the 'help' for that command.
					errMsg := fmt.Sprintf("you did not specify the name of the crypto to be tracked.")
					logger.Error(errMsg)
				}
				trackCrypto.TrackCrypto(cryptoName)
				return nil
			},
		},
		{
			Name:  "portfolio",
			Usage: "View your cryptocurrency portfolio",
			Action: func(c *cli.Context) error {
				viewPortfolio.ViewPortfolio()
				return nil
			},
		},
		{
			Name:  "set-alert",
			Usage: "Set an alert for a cryptocurrency",
			Action: func(c *cli.Context) error {
				setAlert.SetAlert()
				return nil
			},
		},
		// TODO: Add additional web3trakka commands here
	}

	// Initialize the application.
	app := &cli.App{
		Commands: commands,
	}

	err := app.Run(os.Args)
	if err != nil {
		errMsg := fmt.Sprintf("error: %s", err)
		logger.Fatal(errMsg)
		os.Exit(1)
	}
}