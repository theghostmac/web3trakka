package web3trakka

import (
	"fmt"
	"github.com/theghostmac/web3trakka/external/binance"
	"github.com/theghostmac/web3trakka/external/crypto"
	"github.com/theghostmac/web3trakka/internal/housekeeper"
)

var logger = housekeeper.NewCustomLogger()

func NewCryptoTracker() *TrackCrypto {
	return &TrackCrypto{}
}

// TrackCrypto tracks a new cryptocurrency pair.
func (tc *TrackCrypto) TrackCrypto(cryptoSymbol string) (*binance.SymbolDetails, error) {
	binanceClient, err := binance.NewBinanceClient()
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to initialize Binance client: %v", err))
		return nil, err
	}

	details, err := binanceClient.GetSymbolDetails(cryptoSymbol)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to fetch symbol details: %v", err))
		return nil, err
	}

	// Format and display the details
	formattedDetails := crypto.ProvideSymbolDetailsResponse(details)
	fmt.Println(formattedDetails)

	return details, nil
}
