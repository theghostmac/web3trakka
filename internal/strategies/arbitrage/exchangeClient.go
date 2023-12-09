package arbitrage

import "github.com/theghostmac/web3trakka/external/crypto"

type ExchangeClient interface {
	GetSymbolDetails(symbolPair string) (*crypto.SymbolDetails, error)
	ExecuteTrade(symbol, orderType string, price float64) error
}
