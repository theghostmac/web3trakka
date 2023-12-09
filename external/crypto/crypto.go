package crypto

// SymbolDetails represents detailed information about a cryptocurrency symbol.
type SymbolDetails struct {
	Symbol                 string   `json:"symbol"`
	Status                 string   `json:"status"`
	BaseAsset              string   `json:"baseAsset"`
	BaseAssetPrecision     int64    `json:"baseAssetPrecision"`
	QuoteAsset             string   `json:"quoteAsset"`
	QuotePrecision         int64    `json:"quotePrecision"`
	OrderTypes             []string `json:"orderTypes"`
	IcebergAllowed         bool     `json:"icebergAllowed"`
	OcoAllowed             bool     `json:"ocoAllowed"`
	IsSpotTradingAllowed   bool     `json:"isSpotTradingAllowed"`
	IsMarginTradingAllowed bool     `json:"isMarginTradingAllowed"`
	Filters                []Filter `json:"filters"`
	Permissions            []string `json:"permissions"`
	PriceChange            float64  `json:"priceChange"`
	PriceChangePercent     float64  `json:"priceChangePercent"`
	WeightedAvgPrice       float64  `json:"weightedAvgPrice"`
	PrevClosePrice         float64  `json:"prevClosePrice"`
	LastPrice              float64  `json:"lastPrice"`
	BidPrice               float64  `json:"bidPrice"`
	AskPrice               float64  `json:"askPrice"`
	OpenPrice              float64  `json:"openPrice"`
	HighPrice              float64  `json:"highPrice"`
	LowPrice               float64  `json:"lowPrice"`
	Volume                 float64  `json:"volume"`
	QuoteVolume            float64  `json:"quoteVolume"`
	OpenTime               uint64   `json:"openTime"`
	CloseTime              uint64   `json:"closeTime"`
	FirstId                uint64   `json:"firstId"`
	LastId                 uint64   `json:"lastId"`
	Count                  uint64   `json:"count"`
}

// Filter represents various filters applied to a symbol.
type Filter struct {
	FilterType       string  `json:"filterType"`
	MinPrice         float64 `json:"minPrice"`
	MaxPrice         float64 `json:"maxPrice"`
	TickSize         float64 `json:"tickSize"`
	MinQty           float64 `json:"minQty"`
	MaxQty           float64 `json:"maxQty"`
	StepSize         float64 `json:"stepSize"`
	MinNotional      float64 `json:"minNotional"`
	Limit            uint    `json:"limit"`
	MaxNumAlgoOrders int64   `json:"maxNumAlgoOrders"`
}
