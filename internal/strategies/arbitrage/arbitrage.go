package arbitrage

import (
	"fmt"
	"github.com/theghostmac/web3trakka/external/crypto"
	"github.com/theghostmac/web3trakka/internal/housekeeper"
	"reflect"
	"sort"
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
	symbolDetailsMap := a.getSymbolDetailsFromExchanges(symbolPair)

	arbitrageOpportunities := identifyArbitrageOpportunities(symbolDetailsMap)

	if len(arbitrageOpportunities) > 0 {
		a.ExecuteTradeOnExchanges(arbitrageOpportunities)
	}

	return nil
}

// ArbitrageOpportunity represents an arbitrage opportunity on different exchanges.
type ArbitrageOpportunity struct {
	BuyExchange  string
	SellExchange string
	BuyPrice     float64
	SellPrice    float64
	Profit       float64
}

func identifyArbitrageOpportunities(detailsMap map[string]*crypto.SymbolDetails) []ArbitrageOpportunity {
	var opportunities []ArbitrageOpportunity

	// Sort exchanges for consistent order processing.
	exchanges := make([]string, 0, len(detailsMap))
	for exchange := range detailsMap {
		exchanges = append(exchanges, exchange)
	}

	sort.Strings(exchanges)

	// Compare each exchange with every other exchange.
	for i := 0; i < len(exchanges)-1; i++ {
		for j := i + 1; j < len(exchanges); j++ {
			buyExchange := exchanges[i]
			sellExchange := exchanges[j]

			buyDetails := detailsMap[buyExchange]
			sellDetails := detailsMap[sellExchange]

			// Check for potential arbitrage opportunity.
			if buyDetails.AskPrice < sellDetails.BidPrice {
				profit := sellDetails.BidPrice - buyDetails.AskPrice
				opportunities = append(opportunities, ArbitrageOpportunity{
					BuyExchange:  buyExchange,
					SellExchange: sellExchange,
					BuyPrice:     buyDetails.AskPrice,
					SellPrice:    sellDetails.BidPrice,
					Profit:       profit,
				})
			}
		}
	}

	return opportunities
}

// ExecuteTradeOnExchanges places a trade on the exchange.
func (a *Arbitrage) ExecuteTradeOnExchanges(opportunities []ArbitrageOpportunity) {
	// TODO: implement trade execution based on identified arbitrage opportunities.
}

func (a *Arbitrage) getSymbolDetailsFromExchanges(symbolPair string) map[string]*crypto.SymbolDetails {
	// TODO: implement fetching symbol details for all exchanges.
	symbolDetailsMap := make(map[string]*crypto.SymbolDetails)

	// Go through all exchanges and get the symbol details of the pair on them.
	for _, exchange := range a.Exchanges {
		details, err := exchange.GetSymbolDetails(symbolPair)
		if err != nil {
			errMsg := fmt.Sprintf("failed to fetch details: %v", err)
			logger.Error(errMsg)
			continue
		}
		// Use a unique identifier for each exchange.
		exchangeID := a.getExchangeIdentifier(exchange)
		symbolDetailsMap[exchangeID] = details
	}

	return symbolDetailsMap
}

// getExchangeIdentifier returns a unique identifier for an exchange client.
func (a *Arbitrage) getExchangeIdentifier(exchange ExchangeClient) string {
	exchangeType := reflect.TypeOf(exchange)

	if exchangeType.Kind() == reflect.Ptr {
		exchangeType = exchangeType.Elem()
	}

	return exchangeType.Name()
}
