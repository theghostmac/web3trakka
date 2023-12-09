package crypto

import (
	"fmt"
	"github.com/theghostmac/web3trakka/external/binance"
	"strings"
)

// ProvideSymbolDetailsResponse creates a user-friendly string representation of the crypto details.
func ProvideSymbolDetailsResponse(details *binance.SymbolDetails) string {
	if details == nil {
		return "No details available."
	}

	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("Symbol: %s\n", details.Symbol))
	builder.WriteString(fmt.Sprintf("Status: %s\n", details.Status))
	builder.WriteString(fmt.Sprintf("Base Asset: %s\n", details.BaseAsset))
	builder.WriteString(fmt.Sprintf("Quote Asset: %s\n", details.QuoteAsset))
	builder.WriteString(fmt.Sprintf("Is Spot Trading Allowed: %t\n", details.IsSpotTradingAllowed))
	builder.WriteString(fmt.Sprintf("Is Margin Trading Allowed: %t\n", details.IsMarginTradingAllowed))
	builder.WriteString(fmt.Sprintf("Price Change: %s\n", details.PriceChange))
	builder.WriteString(fmt.Sprintf("Price Change Percent: %s%%\n", details.PriceChangePercent))
	builder.WriteString(fmt.Sprintf("Last Price: %s\n", details.LastPrice))
	builder.WriteString(fmt.Sprintf("High Price: %s\n", details.HighPrice))
	builder.WriteString(fmt.Sprintf("Low Price: %s\n", details.LowPrice))
	builder.WriteString(fmt.Sprintf("Volume: %s\n", details.Volume))
	// TODO: Add more fields if necessary.

	return builder.String()
}

// ConvertPairName converts standard pair names to Kraken's format.
func ConvertPairName(pairName string) string {
	return strings.Replace(pairName, "BTC", "XBT", 1)
}
