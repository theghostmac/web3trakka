package binance

import (
	"context"
	"fmt"
	"os"

	binance_connector "github.com/binance/binance-connector-go"
	"github.com/theghostmac/web3trakka/internal/housekeeper"
)

var logger = housekeeper.NewCustomLogger()

// BinanceClient represents a client for Binance API
type BinanceClient struct {
	Client *binance_connector.Client
}

// NewBinanceClient creates a new Binance API client
func NewBinanceClient() (*BinanceClient, error) {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")
	baseURL := os.Getenv("BINANCE_BASE_URL") // change to BINANCE_BASE_URL_PRODUCTION when ready.

	if apiKey == "" || secretKey == "" {
		return nil, fmt.Errorf("API key or Secret key is missing")
	}

	binanceClient := binance_connector.NewClient(apiKey, secretKey, baseURL)
	return &BinanceClient{Client: binanceClient}, nil
}

// GetSymbolDetails fetches and returns details for a specific symbol
func (bc *BinanceClient) GetSymbolDetails(pairSymbol string) (*SymbolDetails, error) {
	// Fetch exchange information
	exchangeInfo, err := bc.Client.NewExchangeInfoService().Do(context.Background())
	if err != nil {
		errMsg := fmt.Sprintf("Failed to fetch exchange information: %v", err)
		logger.Error(errMsg)
		return nil, err
	}

	// Iterate over the symbols to find the specific one
	for _, s := range exchangeInfo.Symbols {
		if s.Symbol == pairSymbol {
			details := SymbolDetails{
				Symbol:                 s.Symbol,
				Status:                 s.Status,
				BaseAsset:              s.BaseAsset,
				BaseAssetPrecision:     s.BaseAssetPrecision,
				QuoteAsset:             s.QuoteAsset,
				QuotePrecision:         s.QuotePrecision,
				OrderTypes:             s.OrderTypes,
				IcebergAllowed:         s.IcebergAllowed,
				IsSpotTradingAllowed:   s.IsSpotTradingAllowed,
				IsMarginTradingAllowed: s.IsMarginTradingAllowed,

				// ... Populate other fields from s (the SymbolInfo)
			}

			// Optional: Fetch additional data like 24hr ticker price change statistics
			ticker24hr, err := bc.Client.NewTicker24hrService().Symbol(pairSymbol).Do(context.Background())
			//logger.Info(fmt.Sprintf("Fetching ticker24hr: %+v", ticker24hr))
			if err != nil {
				errMsg := fmt.Sprintf("Failed to fetch 24hr ticker data for pairSymbol %s: %v", pairSymbol, err)
				logger.Warning(errMsg)
			} else {
				// Populate the price change statistics fields in details
				details.PriceChange = ticker24hr.PriceChange
				details.PriceChangePercent = ticker24hr.PriceChangePercent
				details.WeightedAvgPrice = ticker24hr.WeightedAvgPrice
				details.PrevClosePrice = ticker24hr.PrevClosePrice
				details.LastPrice = ticker24hr.LastPrice
				details.AskPrice = ticker24hr.AskPrice
				details.BidPrice = ticker24hr.BidPrice
				details.OpenPrice = ticker24hr.OpenPrice
				details.HighPrice = ticker24hr.HighPrice
				details.LowPrice = ticker24hr.LowPrice
				details.Volume = ticker24hr.Volume
				details.QuoteVolume = ticker24hr.QuoteVolume
				details.OpenTime = ticker24hr.OpenTime
				details.CloseTime = ticker24hr.CloseTime

				// ... Populate other fields from ticker24hr
			}

			return &details, nil
		}
	}

	return nil, fmt.Errorf("pairSymbol %s not found", pairSymbol)
}
