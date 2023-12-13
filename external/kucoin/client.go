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
		Symbol:      cryptoPair,
		LastPrice:   lastPrice,
		BidPrice:    bidPrice,
		AskPrice:    askPrice,
		OpenPrice:   openPrice,
		LowPrice:    lowPrice,
		HighPrice:   highPrice,
		Volume:      volume,
		QuoteVolume: quoteVolume,
	}, nil
}

// ExecuteTrade executes a trade on Kucoin.
// Parameters:
// - symbol: the trading pair
// - side: "buy" or "sell"
// - tradeType: "market" or "limit"
// - size: the amount to buy or sell
// - price: the price to buy/sell at (only for limit orders)
func (kc *KucoinClient) ExecuteTrade(symbol string, orderType string, price float64) error {
	var side string
	var tradeType string

	// Determine the trade type and side based on orderType parameter
	switch orderType {
	case "buy-market":
		tradeType = "market"
		side = "buy"
	case "sell-market":
		tradeType = "market"
		side = "sell"
	case "buy-limit":
		tradeType = "limit"
		side = "buy"
	case "sell-limit":
		tradeType = "limit"
		side = "sell"
	default:
		return fmt.Errorf("invalid order type: %s", orderType)
	}

	// Construct the order request
	var order kucoin.CreateOrderModel
	order.ClientOid = fmt.Sprintf("oid-%d", os.Getpid()) // Unique order ID
	order.Symbol = symbol
	order.Side = side
	order.Type = tradeType
	order.Price = strconv.FormatFloat(price, 'f', -1, 64)
	// Size needs to be set, this example assumes a fixed size, modify as needed
	order.Size = "1" // Placeholder size, adjust based on your logic

	// Execute the order
	resp, err := kc.Client.CreateOrder(&order)
	if err != nil {
		return fmt.Errorf("failed to execute trade: %v", err)
	}

	var orderResp kucoin.ApiService
	err = resp.ReadData(&orderResp)
	if err != nil {
		return fmt.Errorf("failed to read order response: %v", err)
	}

	fmt.Printf("Order executed successfully: %s\n", orderResp)
	return nil
}
