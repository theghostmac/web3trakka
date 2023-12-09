package main

import (
	"fmt"
	"github.com/theghostmac/web3trakka/external/binance"
	"github.com/theghostmac/web3trakka/external/kraken"
	"github.com/theghostmac/web3trakka/external/okex"
	"github.com/theghostmac/web3trakka/internal/housekeeper"
	"github.com/theghostmac/web3trakka/internal/runner"
	"github.com/theghostmac/web3trakka/internal/strategies/arbitrage"
	"github.com/theghostmac/web3trakka/internal/web3trakka"
	"github.com/urfave/cli/v2"
	"os"
)

// initialize the housekeeper.
var logger = housekeeper.NewCustomLogger()

func main() {
	// Initialize the runner.
	startRunner := runner.NewStartRunner()

	var exchanges []arbitrage.ExchangeClient

	// Initialize clients for exchanges.
	binanceClient, err := binance.NewBinanceClient()
	if err != nil {
		errMsg := fmt.Sprintf("failed to initialize binance client due to: %v", err)
		logger.Error(errMsg)
	} else {
		exchanges = append(exchanges, binanceClient)
	}

	krakenClient, err := kraken.NewKrakenClient()
	if err != nil {
		errMsg := fmt.Sprintf("failed to initialize kraken client due to: %v", err)
		logger.Error(errMsg)
	} else {
		exchanges = append(exchanges, krakenClient)
	}

	okEXClient, err := okex.NewOKEXClient()
	if err != nil {
		errMsg := fmt.Sprintf("failed to initialize binance client due to: %v", err)
		logger.Error(errMsg)
	} else {
		exchanges = append(exchanges, okEXClient)
	}

	trackCrypto := web3trakka.NewCryptoTracker()
	viewPortfolio := web3trakka.NewPortfolioViewer()
	setAlert := web3trakka.NewAlertSetter()

	arbitrager := arbitrage.NewArbitrage([]arbitrage.ExchangeClient{
		binanceClient, krakenClient,
	})

	// Define web3trakka commands
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
				symbolPair := c.Args().First() // Gets the first argument.
				if symbolPair == "" {
					errMsg := "Error: No cryptocurrency symbol provided. Please specify a symbol."
					logger.Error(errMsg)
					cli.ShowCommandHelp(c, "track")
					return fmt.Errorf(errMsg)
				}

				details, err := trackCrypto.TrackCrypto(symbolPair)
				if err != nil {
					logger.Error(err.Error())
					return err
				}

				fmt.Printf("Details for %s:\n%+v\n", symbolPair, details)
				return nil
			},
		},
		{
			Name:  "arbitrage",
			Usage: "Find arbitrage opportunities across exchanges",
			Action: func(c *cli.Context) error {
				symbolPair := c.Args().First() // Gets the first argument.
				if symbolPair == "" {
					errMsg := "Error: No cryptocurrency pair symbol provided. Please specify a pair."
					logger.Error(errMsg)
					cli.ShowCommandHelp(c, "arbitrage")
					return fmt.Errorf(errMsg)
				}
				err := arbitrager.FindArbitrage(symbolPair)
				if err != nil {
					errMsg := fmt.Sprintf("Failed to call the FindArbitrage method against the pair: %v", err)
					logger.Error(errMsg)
				}

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

	err = app.Run(os.Args)
	if err != nil {
		errMsg := fmt.Sprintf("error: %s", err)
		logger.Fatal(errMsg)
		os.Exit(1)
	}
}
