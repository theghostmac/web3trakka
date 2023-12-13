package kucoin

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/theghostmac/web3trakka/external/crypto"
	"github.com/theghostmac/web3trakka/internal/housekeeper"
)

// KucoinClient represents a client for Kucoin API.
type KucoinClientV2 struct {
	BaseURL string
	APIKey string
	APISecret string
	Version string
	Client *http.Client
}

// NewKucoinClientV2 creates a new Kucoin client.
func NewKucoinClientV2(apiKey, apiSecret, baseURL, version string) (*KucoinClientV2, error) {
	if apiKey == "" || apiSecret == "" || baseURL == "" {
		return nil, errors.New("missing Kucoin API credentials or bae URL")
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	
	return &KucoinClientV2{
		BaseURL: baseURL,
		APIKey: apiKey,
		APISecret: apiSecret,
		Version: version,
		Client: client,
	}, nil
}

var logger = housekeeper.CustomLogger{}

// GetSymbolDetails fetches and returns details for a specific pair.
func (kc *KucoinClientV2) GetSymbolDetails(cryptoPair string) (*crypto.SymbolDetails, error) {
	url := fmt.Sprintf("%s/api/%s/market/stats", kc.BaseURL, kc.Version)
	// query := map[string]string{"symbol": cryptoPair}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		logger.Error(fmt.Sprintf("Error creating request: %v", err))
		return nil, err
	}

	req.URL.Query().Add("symbol", cryptoPair)
	resp, err := kc.Client.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("Error making request: %v", err)
		logger.Error(errMsg)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
        errMsg := fmt.Sprintf("Unsuccessful response for GetSymbolDetails(%s): %d", cryptoPair, resp.StatusCode)
        logger.Error(errMsg)
        return nil, errors.New(errMsg)
    }

    var tickerData map[string]interface{}
    err = json.NewDecoder(resp.Body).Decode(&tickerData)
    if err != nil {
        errMsg := fmt.Sprintf("Error parsing ticker data: %v", err)
		logger.Error(errMsg)
        return nil, err
    }

    return kc.parseKucoinToSymbolDetails(cryptoPair, tickerData)
}

// parseKucoinToSymbolDetails converts Kucoin ticker data to crypto.SymbolDetails format.
func (kc *KucoinClientV2) parseKucoinToSymbolDetails(cryptoPair string, tickerData map[string]interface{}) (*crypto.SymbolDetails, error) {
    lastTrade, ok := tickerData["last"]
    if !ok {
        return nil, errors.New("missing 'last' field in ticker data")
    }
    lastPrice, err := kc.convertToFloat64(lastTrade)
    if err != nil {
        return nil, fmt.Errorf("error converting last price: %v", err)
    }

    high, ok := tickerData["high"]
    if !ok {
        return nil, errors.New("missing 'high' field in ticker data")
    }
    high24h, err := kc.convertToFloat64(high)
    if err != nil {
        return nil, fmt.Errorf("error converting high 24h: %v", err)
    }

    low, ok := tickerData["low"]
    if !ok {
        return nil, errors.New("missing 'low' field in ticker data")
    }
    low24h, err := kc.convertToFloat64(low)
    if err != nil {
        return nil, fmt.Errorf("error converting low 24h: %v", err)
    }

    bid, ok := tickerData["buy"]
    if !ok {
        return nil, errors.New("missing 'buy' field in ticker data")
    }
    bidPrice, err := kc.convertToFloat64(bid)
    if err != nil {
        return nil, fmt.Errorf("error converting bid price: %v", err)
    }

    ask, ok := tickerData["sell"]
    if !ok {
        return nil, errors.New("missing 'sell' field in ticker data")
    }
    askPrice, err := kc.convertToFloat64(ask)
    if err != nil {
        return nil, fmt.Errorf("error converting ask price: %v", err)
    }

    volume, ok := tickerData["vol"]
    if !ok {
        return nil, errors.New("missing 'vol' field in ticker data")
    }
    volume24h, err := kc.convertToFloat64(volume)
    if err != nil {
        return nil, fmt.Errorf("error converting volume 24h: %v", err)
    }

    // TODO: Extract other relevant fields from tickerData and set them on SymbolDetails ...

    return &crypto.SymbolDetails{
        Symbol:     strings.ToUpper(cryptoPair), // TODO: Update with actual crypto pair extraction
        // TODO: Populate other fields from extracted data
        LastPrice:  lastPrice,
		HighPrice: high24h,
		LowPrice: low24h,
        BidPrice:   bidPrice,
        AskPrice:   askPrice,
        Volume: volume24h,
    }, nil
}

// convertToFloat64 safely converts an interface{} value to a float64.
func (kc *KucoinClientV2) convertToFloat64(value interface{}) (float64, error) {
    switch t := value.(type) {
    case float64:
        return t, nil
    case string:
        f, err := strconv.ParseFloat(t, 64)
        if err != nil {
            return 0, err
        }
        return f, nil
    default:
        return 0, fmt.Errorf("invalid type for conversion: %T", value)
    }
}
