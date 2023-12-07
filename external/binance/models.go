package binance

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
	PriceChange            string   `json:"priceChange"`
	PriceChangePercent     string   `json:"priceChangePercent"`
	WeightedAvgPrice       string   `json:"weightedAvgPrice"`
	PrevClosePrice         string   `json:"prevClosePrice"`
	LastPrice              string   `json:"lastPrice"`
	BidPrice               string   `json:"bidPrice"`
	AskPrice               string   `json:"askPrice"`
	OpenPrice              string   `json:"openPrice"`
	HighPrice              string   `json:"highPrice"`
	LowPrice               string   `json:"lowPrice"`
	Volume                 string   `json:"volume"`
	QuoteVolume            string   `json:"quoteVolume"`
	OpenTime               uint64   `json:"openTime"`
	CloseTime              uint64   `json:"closeTime"`
	FirstId                uint64   `json:"firstId"`
	LastId                 uint64   `json:"lastId"`
	Count                  uint64   `json:"count"`
}

// Filter represents various filters applied to a symbol.
type Filter struct {
	FilterType       string `json:"filterType"`
	MinPrice         string `json:"minPrice"`
	MaxPrice         string `json:"maxPrice"`
	TickSize         string `json:"tickSize"`
	MinQty           string `json:"minQty"`
	MaxQty           string `json:"maxQty"`
	StepSize         string `json:"stepSize"`
	MinNotional      string `json:"minNotional"`
	Limit            uint   `json:"limit"`
	MaxNumAlgoOrders int64  `json:"maxNumAlgoOrders"`
}
