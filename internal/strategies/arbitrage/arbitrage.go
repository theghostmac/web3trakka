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

// FindArbitrageOpportunities finds arbitrage opportunities across exchanges.
func (a *Arbitrage) FindArbitrageOpportunities(symbolPair string) error {
	var symbolDetailsMap = make(map[string]*crypto.SymbolDetails)

	// Go through all exchanges and get the symbol details of the pair on them.
	for _, exchange := range a.Exchanges {
		details, err := exchange.GetSymbolDetails(symbolPair)
		if err != nil {
			errMsg := fmt.Sprintf("failed to fetch details: %v", err)
			logger.Error(errMsg)
			return err
		}
		symbolDetailsMap[details.Symbol] = details
	}

	// Compare prices of the pair from the different exchanges.
	for symbol, details := range symbolDetailsMap {
		// TODO: move from this, to the proper logic.
		detailsMsg := fmt.Sprintf("Details for %s: %+v", symbol, details)
		logger.Info(detailsMsg)
	}

	// TODO: call ExecuteTradeOnExchanges.
	return nil
}

// ExecuteTradeOnExchanges places a trade on the exchange.
func (a *Arbitrage) ExecuteTradeOnExchanges() {
	// TODO: Compare prices and calculate profit potential.
	// If the profit is considerably good,
	// Log or act on opportunity.
}
