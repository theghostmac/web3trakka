package arbitrage

import "github.com/theghostmac/web3trakka/external/crypto"

type ExchangeClient interface {
	GetSymbolDetails(symbolPair string) (*crypto.SymbolDetails, error)
	ExecuteTrade() // TODO: add necessary parameters here.
}
