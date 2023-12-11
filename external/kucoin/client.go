package kucoin

import (
	"github.com/Kucoin/kucoin-go-sdk"
	"github.com/theghostmac/web3trakka/external/crypto"
)

type KucoinClient struct {
	Client *kucoin.ApiService
}

type Opts struct {
}

func NewKucoinClient(key, secret, passphrase string) *KucoinClient {
	client := kucoin.NewApiService()
	return &KucoinClient{Client: client}
}

func (kc *KucoinClient) GetSymbolDetails(cryptoPair string) (*crypto.SymbolDetails, error) {

	return nil, nil
}

func (kc *KucoinClient) ExecuteTrade() {

}
