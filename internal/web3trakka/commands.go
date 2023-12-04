package web3trakka

type TrackCrypto struct {
}

func NewCryptoTracker() *TrackCrypto {
	return &TrackCrypto{}
}

func (tc *TrackCrypto) TrackCrypto(cryptoName string) {
	// TODO: implement logic to track the specified cryptocurrency.
}

type ViewPortfolio struct {
}

func NewPortfolioViewer() *ViewPortfolio {
	return &ViewPortfolio{}
}

func (pv *ViewPortfolio) ViewPortfolio() {

}

type SetAlert struct {
}

func NewAlertSetter() *SetAlert {
	return &SetAlert{}
}

func (sa *SetAlert) SetAlert() {}
