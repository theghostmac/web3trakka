package arbitrage

import (
	"fmt"
	"github.com/theghostmac/web3trakka/external/crypto"
	"github.com/theghostmac/web3trakka/internal/housekeeper"
)

// Arbitrage model subject to implementation.
type Arbitrage struct {
	Exchanges []ExchangeClient
}

// NewArbitrage creates a new Arbitrage instance with given exchanges.
func NewArbitrage(exchanges []ExchangeClient) *Arbitrage {
	return &Arbitrage{Exchanges: exchanges}
}

var logger = housekeeper.NewCustomLogger()

// FindArbitrage finds arbitrage opportunities across exchanges.
func (a *Arbitrage) FindArbitrage(symbolPair string) error {
	var symbolDetailsMap = make(map[string]*crypto.SymbolDetails)

	for _, exchange := range a.Exchanges {
		details, err := exchange.GetSymbolDetails(symbolPair)
		if err != nil {
			errMsg := fmt.Sprintf("failed to fetch details: %v", err)
			logger.Error(errMsg)
			return err
		}
		symbolDetailsMap[details.Symbol] = details
	}

	// Compare prices from different exchanges.
	for symbol, details := range symbolDetailsMap {
		// TODO: move from a
		detailsMsg := fmt.Sprintf("Details for %s: %+v", symbol, details)
		logger.Info(detailsMsg)
	}

	// TODO: call ExecuteTrade.
	return nil
}

// ExecuteTrade places a trade on the exchange.
func (a *Arbitrage) ExecuteTrade() {
	// TODO: Compare prices and calculate profit potential.
	// If the profit is considerably good,
	// Log or act on opportunity.
}
