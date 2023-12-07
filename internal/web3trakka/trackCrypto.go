package web3trakka

func NewCryptoTracker() *TrackCrypto {
	return &TrackCrypto{}
}

func (tc *TrackCrypto) TrackCrypto(cryptoName string) {
	// TODO: implement logic to track the specified cryptocurrency.
	// use binance api to fetch the latest price of the specified cryptocurrency.
}
