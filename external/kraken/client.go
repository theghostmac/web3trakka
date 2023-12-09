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
		logger.Error(fmt.Sprintf("kraken API key: %s or Private key: %s", apiKey, secretKey))
		return nil, fmt.Errorf("kraken API key or Private key is not correct")
	}

	krakenClient := krakenapi.New(apiKey, secretKey)
	return &KrakenClient{Client: krakenClient}, nil
}

// NormalizePairDetails normalizes the data returned from Kraken API.
func NormalizePairDetails(pairData map[string]interface{}) (*crypto.SymbolDetails, error) {
	lastTrade, ok := pairData["c"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected last trade data format from Kraken API")
	}
	lastTradePrice := lastTrade[0].(string)

	volume, ok := pairData["v"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected volume data format from Kraken API")
	}
	volume24h := volume[1].(string) // Assuming the 24h volume is the second element

	high, ok := pairData["h"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected high data format from Kraken API")
	}
	high24h := high[1].(string) // Assuming the 24h high is the second element

	low, ok := pairData["l"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected low data format from Kraken API")
	}
	low24h := low[1].(string) // I assume the 24h low is the second element

	// Extracting additional details
	bidData, ok := pairData["b"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected bid data format from Kraken API")
	}
	bidPrice := bidData[0].(string)

	askData, ok := pairData["a"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected ask data format from Kraken API")
	}
	askPrice := askData[0].(string)

	details := &crypto.SymbolDetails{
		Symbol:                 pair,
		Status:                 "",
		BaseAsset:              "",
		BaseAssetPrecision:     0,
		QuoteAsset:             "",
		QuotePrecision:         0,
		OrderTypes:             nil,
		IcebergAllowed:         false,
		OcoAllowed:             false,
		IsSpotTradingAllowed:   false,
		IsMarginTradingAllowed: false,
		Filters:                nil,
		Permissions:            nil,
		PriceChange:            "",
		PriceChangePercent:     "",
		WeightedAvgPrice:       "",
		PrevClosePrice:         "",
		LastPrice:              lastTradePrice,
		BidPrice:               bidPrice,
		AskPrice:               askPrice,
		OpenPrice:              "",
		HighPrice:              high24h,
		LowPrice:               low24h,
		Volume:                 volume24h,
		QuoteVolume:            "",
		OpenTime:               0,
		CloseTime:              0,
		FirstId:                0,
		LastId:                 0,
		Count:                  0,
	}

	return details, nil
}

// GetSymbolDetails fetches and returns details for a specific pair
func (kc *KrakenClient) GetSymbolDetails(pair string) (*crypto.SymbolDetails, error) {
	// Convert pair to Kraken's format.
	krakenPair := crypto.ConvertPairName(pair)
	response, err := kc.Client.Query("Ticker", map[string]string{"pair": krakenPair})
	if err != nil {
		logger.Error("Failed to fetch Ticker prices from Kraken Client")
		return nil, fmt.Errorf("error querying Kraken API: %v", err)
	}

	// Log the raw response for debugging
	logger.Info(fmt.Sprintf("Raw Kraken response: %+v", response))

	// Asserting response to be of type map[string]interface{}
	responseMap, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected type for Kraken response")
	}

	// Check and log the structure of the 'result' field
	logger.Info(fmt.Sprintf("Result field type: %T", responseMap["result"]))

	result, ok := responseMap["result"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected type for result field in Kraken response")
	}

	pairData, ok := result[krakenPair].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("pair data not found in response for %s", krakenPair)
	}

	return NormalizePairDetails(pairData)
}
