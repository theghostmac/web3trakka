package kucoin

import (
	"github.com/Kucoin/kucoin-go-sdk"
	"github.com/theghostmac/web3trakka/external/crypto"
	"os"
)

type KucoinClient struct {
	Client *kucoin.ApiService
}

// NewKucoinClient creates a new Kucoin client.
func NewKucoinClient() *KucoinClient {
	apiKey := os.Getenv("KUCOIN_API_KEY")
	baseURL := os.Getenv("KUCOIN_PRODUCTION_URL")
	apiSecret := os.Getenv("KUCOIN_API_SECRET")

	client := kucoin.NewApiService(
		kucoin.ApiKeyOption(apiKey),
		kucoin.ApiBaseURIOption(baseURL),
		kucoin.ApiSecretOption(apiSecret),
		kucoin.ApiKeyVersionOption(kucoin.ApiKeyVersionV2),
	)

	return &KucoinClient{Client: client}
}

func (kc *KucoinClient) GetSymbolDetails(cryptoPair string) (*crypto.SymbolDetails, error) {

	return nil, nil
}

func (kc *KucoinClient) ExecuteTrade() {}
