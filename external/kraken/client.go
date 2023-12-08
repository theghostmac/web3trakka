package kraken

import (
	"fmt"
	"github.com/beldur/kraken-go-api-client"
	"github.com/theghostmac/web3trakka/external/crypto"
	"github.com/theghostmac/web3trakka/internal/housekeeper"
	"os"
)

// KrakenClient represents a client for Kraken API.
type KrakenClient struct {
	Client *krakenapi.KrakenAPI
}

var logger = housekeeper.NewCustomLogger()

// NewKrakenClient creates a new Kraken API client.
func NewKrakenClient() (*KrakenClient, error) {
	apiKey := os.Getenv("KRAKEN_API_KEY")
	secretKey := os.Getenv("KRAKEN_SECRET_KEY")

	if apiKey == "" || secretKey == "" {
		return nil, fmt.Errorf("kraken API key or Private key is not correct")
	}

	krakenClient := krakenapi.New(apiKey, secretKey)
	return &KrakenClient{Client: krakenClient}, nil
}

// GetSymbolDetails fetches and returns details for a specific pair
func (kc *KrakenClient) GetSymbolDetails(pair string) (*crypto.SymbolDetails, error) {
	response, err := kc.Client.Query("Ticker", map[string]string{"pair": pair})
	if err != nil {
		logger.Error("Failed to fetch Ticker prices from Kraken Client")
		return nil, fmt.Errorf("error querying Kraken API: %v", err)
	}

	// TODO: delete these.
	fmt.Println(response)
	return nil, err
}
