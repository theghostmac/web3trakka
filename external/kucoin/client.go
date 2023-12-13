package kucoin

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Kucoin/kucoin-go-sdk"
	"github.com/theghostmac/web3trakka/external/crypto"
)

// KucoinClient represents a client for Kucoin API.
type KucoinClient struct {
    Client *kucoin.ApiService
}

// NewKucoinClient creates a new Kucoin client.
func NewKucoinClient() (*KucoinClient, error) {
    apiKey := os.Getenv("KUCOIN_API_KEY")
    apiSecret := os.Getenv("KUCOIN_API_SECRET")
    passphrase := os.Getenv("KUOIN_PASSPHRASE")

    if apiKey == "" || apiSecret == "" || passphrase == "" {
        return nil, fmt.Errorf("API Key, Secret, or Passphrase are not set")
    }

    client := kucoin.NewApiService(
        kucoin.ApiKeyOption(apiKey),
        kucoin.ApiSecretOption(apiSecret),
        kucoin.ApiPassPhraseOption(passphrase),
    )

    return &KucoinClient{Client: client}, nil
}

// GetSymbolDetails fetches and returns details for a specific pair.
func (kc *KucoinClient) GetSymbolDetails(cryptoPair string) (*crypto.SymbolDetails, error) {
    resp, err := kc.Client.TickerLevel1(cryptoPair)
    if err != nil {
        return nil, fmt.Errorf("failed to fetch ticker for pair %s: %v", cryptoPair, err)
    }

	var tickerLevel1 kucoin.TickerLevel1Model
	err = resp.ReadData(&tickerLevel1)
	if err != nil {
		return nil, fmt.Errorf("failed to read ticker data for pair %s: %v", cryptoPair, err)
	}

	lastPrice, err := strconv.ParseFloat(tickerLevel1.Price, 64)
	if err != nil {
        return nil, fmt.Errorf("error converting last price: %v", err)
    }

    bidPrice, err := strconv.ParseFloat(tickerLevel1.BestBid, 64)
    if err != nil {
        return nil, fmt.Errorf("error converting bid price: %v", err)
    }

    askPrice, err := strconv.ParseFloat(tickerLevel1.BestAsk, 64)
    if err != nil {
        return nil, fmt.Errorf("error converting ask price: %v", err)
    }

    // Estimate high and low price based on last trade size (assumes constant price during trade)
    size, err := strconv.ParseFloat(tickerLevel1.Size, 64)
    if err != nil {
        return nil, fmt.Errorf("error converting size: %v", err)
    }

	// -------> EVERYTHING HERE IS NOT ACCURATE <-------
	
    highPrice := lastPrice + size/100 // Assuming 1% size impact on price
    lowPrice := lastPrice - size/100

	// Estimate open price (replace with actual value if available)
    openPrice := lastPrice // Replace with actual open price if available

    // Estimate volume based on accumulated last trade size (not comprehensive)
    volume := size // Accumulate size of all trades for accurate volume

    // Calculate quote volume
    quoteVolume := lastPrice * size

    return &crypto.SymbolDetails{
        Symbol:     cryptoPair,
        LastPrice:  lastPrice,
        BidPrice:   bidPrice,
        AskPrice:   askPrice,
		OpenPrice:  openPrice,
		LowPrice:   lowPrice,
        HighPrice:  highPrice,
		Volume: volume,
		QuoteVolume: quoteVolume,
    }, nil
}

func (kc *KucoinClient) ExecuteTrade() {
    // Implementation for executing a trade
}
