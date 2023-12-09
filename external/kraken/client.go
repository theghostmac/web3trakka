package kraken

import (
	"fmt"
	"github.com/beldur/kraken-go-api-client"
	"github.com/theghostmac/web3trakka/external/crypto"
	"github.com/theghostmac/web3trakka/internal/housekeeper"
	"os"
	"strconv"
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

// convertToFloat64 converts an interface to float64.
func convertToFloat64(value interface{}) (float64, error) {
	switch v := value.(type) {
	case string:
		return strconv.ParseFloat(v, 64)
	case float64:
		return v, nil
	default:
		return 0, fmt.Errorf("unexpected type: %T", v)
	}
}

// convertToUint64 converts an interface to uint64.
func convertToUint64(value interface{}) (uint64, error) {
	switch v := value.(type) {
	case string:
		return strconv.ParseUint(v, 10, 64)
	case float64:
		return uint64(v), nil
	default:
		return 0, fmt.Errorf("unexpected type: %T", v)
	}
}

// GetSymbolDetails fetches and returns details for a specific pair
func (kc *KrakenClient) GetSymbolDetails(pair string) (*crypto.SymbolDetails, error) {
	// Convert pair to Kraken's format.
	krakenPair := crypto.ConvertPairName(pair)

	// Query Kraken API
	response, err := kc.Client.Query("Ticker", map[string]string{"pair": krakenPair})
	if err != nil {
		logger.Error("Failed to fetch Ticker prices from Kraken Client")
		return nil, fmt.Errorf("error querying Kraken API: %v", err)
	}

	// Extract the data from the response
	data, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected response format from Kraken API")
	}

	// Extract the pair data using the correct key.
	pairData, ok := data[krakenPair].(map[string]interface{})
	if !ok {
		// If the key doesn't match, log the keys for debugging
		keys := make([]string, 0, len(data))
		for k := range data {
			keys = append(keys, k)
		}
		logger.Info(fmt.Sprintf("available keys in response: %v", keys))
		return nil, fmt.Errorf("pair data not found in response for %s", krakenPair)
	}

	return NormalizePairDetails(pairData)
}

func (kc *KrakenClient) ExecuteTrade(symbol, orderType string, price float64) error {
	return nil
}

// NormalizePairDetails normalizes the data returned from Kraken API.
func NormalizePairDetails(pairData map[string]interface{}) (*crypto.SymbolDetails, error) {
	lastTrade, ok := pairData["c"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected or missing 'c' field in response")
	}
	lastTradePrice, err := convertToFloat64(lastTrade[0])
	if err != nil {
		return nil, fmt.Errorf("error converting last trade price: %v", err)
	}

	high, ok := pairData["h"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected or missing 'h' in response")
	}
	high24h, err := convertToFloat64(high[1])
	if err != nil {
		return nil, fmt.Errorf("error converting high 24hrs: %v", err)
	}

	low, ok := pairData["l"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected or missing 'l' in response")
	}
	low24h, err := convertToFloat64(low[1])
	if err != nil {
		return nil, fmt.Errorf("error converting low 24hrs: %v", err)
	}

	bidData, ok := pairData["b"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected or missing 'b' in response")
	}
	bidPrice, err := convertToFloat64(bidData[0])
	if err != nil {
		return nil, fmt.Errorf("error converting bid price: %v", err)
	}

	askData, ok := pairData["a"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected or missing 'a' in response")
	}
	askPrice, err := convertToFloat64(askData[0])
	if err != nil {
		return nil, fmt.Errorf("error converting ask price: %v", err)
	}

	// Extracting the 'o' field (open price)
	openPriceVal, ok := pairData["o"]
	if !ok {
		return nil, fmt.Errorf("unexpected or missing 'o' field in response")
	}
	openPrice, err := convertToFloat64(openPriceVal)
	if err != nil {
		return nil, fmt.Errorf("error converting open price: %v", err)
	}

	volume, ok := pairData["v"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected or missing 'v' field in response")
	}
	volume24h, err := convertToFloat64(volume[1])
	if err != nil {
		return nil, fmt.Errorf("error converting volume 24hrs: %v", err)
	}

	avgPriceData, ok := pairData["p"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected average price data format from Kraken API")
	}
	avgPriceStr := avgPriceData[1].(string) // Assuming 24h average price

	avgPrice, err := strconv.ParseFloat(avgPriceStr, 64)
	if err != nil {
		return nil, fmt.Errorf("could not convert average price to float64: %s", err)
	}

	totalTradesData, ok := pairData["t"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected total trades data format from Kraken API")
	}
	totalTrades, err := convertToUint64(totalTradesData[1])
	if err != nil {
		return nil, fmt.Errorf("error converting total trades: %v", err)
	}

	// Building the SymbolDetails struct
	details := &crypto.SymbolDetails{
		Symbol:           "", // TODO: set the Symbol in the caller function.
		LastPrice:        lastTradePrice,
		HighPrice:        high24h,
		LowPrice:         low24h,
		Volume:           volume24h,
		BidPrice:         bidPrice,
		AskPrice:         askPrice,
		OpenPrice:        openPrice,
		WeightedAvgPrice: avgPrice,
		Count:            totalTrades,
	}

	return details, nil
}
