package okex

import "github.com/theghostmac/web3trakka/external/crypto"

// OKEXClient represents a client for the OK-Ex API.
type OKEXClient struct {
}

func NewOKEXClient() (*OKEXClient, error) {
	return &OKEXClient{}, nil
}

func (oc *OKEXClient) GetSymbolDetails(pair string) (*crypto.SymbolDetails, error) {
	return nil, nil
}

func (oc *OKEXClient) ExecuteTrade() {}
