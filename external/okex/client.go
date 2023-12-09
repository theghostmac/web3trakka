package okex

import (
	"context"
	"fmt"
	"github.com/amir-the-h/okex"
	"github.com/amir-the-h/okex/api"
	market2 "github.com/amir-the-h/okex/models/market"
	"github.com/amir-the-h/okex/requests/rest/market"
	"github.com/theghostmac/web3trakka/external/crypto"
	"github.com/theghostmac/web3trakka/internal/housekeeper"
	"os"
)

// OKEXClient represents a client for the OK-Ex API.
type OKEXClient struct {
	Client *api.Client
}

var logger = housekeeper.NewCustomLogger()

func NewOKEXClient() (*OKEXClient, error) {
	apiKey := os.Getenv("OKEX_API_KEY")
	apiSecret := os.Getenv("OKEX_SECRET_KEY")
	passphrase := os.Getenv("OKEX_PASSPHRASE")

	if apiSecret == "" || apiKey == "" {
		return nil, fmt.Errorf("API Key and/or Secret are not set")
	}

	ctx := context.Background()

	okexClient, err := api.NewClient(ctx, apiKey, apiSecret, passphrase, okex.AwsServer)
	if err != nil {
		errMsg := fmt.Sprintf("failed to create a new OkEX client: %v", err)
		logger.Error(errMsg)
		return nil, err
	}

	return &OKEXClient{
		Client: okexClient,
	}, nil
}

func (oc *OKEXClient) GetSymbolDetails(pair string) (*crypto.SymbolDetails, error) {
	// Fetch ticker information using a request.
	tickerReq := market.GetTickers{
		InstType: okex.SpotInstrument,
	}

	// Fetch the ticker data as a response.
	tickerResp, err := oc.Client.Rest.Market.GetTickers(tickerReq)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch ticker information: %v", err)
	}

	// Find the specific ticker for the given pair.
	var tickerData *market2.Ticker
	for _, t := range tickerResp.Tickers {
		if t.InstID == pair {
			tickerData = t
			break
		}
	}

	if tickerData == nil {
		return nil, fmt.Errorf("symbol %s not found", pair)
	}

	// Map the data to the SymbolDetails struct.
	details := &crypto.SymbolDetails{
		Symbol:                 string(tickerData.InstType),
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
		PriceChange:            0,
		PriceChangePercent:     0,
		WeightedAvgPrice:       0,
		PrevClosePrice:         0,
		LastPrice:              float64(tickerData.Last),
		BidPrice:               float64(tickerData.BidPx),
		AskPrice:               float64(tickerData.AskPx),
		OpenPrice:              float64(tickerData.Open24h),
		HighPrice:              float64(tickerData.High24h),
		LowPrice:               float64(tickerData.Low24h),
		Volume:                 float64(tickerData.Vol24h),
		QuoteVolume:            0,
		OpenTime:               0,
		CloseTime:              0,
		FirstId:                0,
		LastId:                 0,
		Count:                  0,
	}
	return details, nil
}

func (oc *OKEXClient) ExecuteTrade() {}
