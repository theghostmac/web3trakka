package okex

import (
	"context"
	"fmt"
	"github.com/amir-the-h/okex"
	"github.com/amir-the-h/okex/api"
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
	return nil, nil
}

func (oc *OKEXClient) ExecuteTrade() {}
