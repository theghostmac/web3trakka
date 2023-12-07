package binance

import (
	"fmt"
	binance_connector "github.com/binance/binance-connector-go"
	"github.com/theghostmac/web3trakka/internal/housekeeper"
	"os"
)

var logger = housekeeper.NewCustomLogger()

type BinanceClient struct {
	APIKey    string
	SecretKey string
	BaseURL   string
	Client    *binance_connector.Client
}

// NewBinanceClient creates a new Binance API Client with provided parameters.
func NewBinanceClient() (*BinanceClient, error) {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")
	baseURL := os.Getenv("BINANCE_BASE_URL")

	if apiKey == "" || secretKey == "" {
		return nil, fmt.Errorf("API key or Secret key is missing")
	}

	// Initialize Binance Connector Client
	binanceClient := binance_connector.NewClient(apiKey, secretKey, baseURL)

	return &BinanceClient{
		APIKey:    apiKey,
		SecretKey: secretKey,
		BaseURL:   baseURL,
		Client:    binanceClient,
	}, nil
}

// ConnectToWebsocket initializes a websocket connection for Market/User Data Stream.
func (bc *BinanceClient) ConnectToWebsocket(isCombined bool) error {
	// Create Websocket Client
	websocketClient := binance_connector.NewWebsocketStreamClient(isCombined, bc.BaseURL)

	// Logic for handling websocket stream.
	// For example, subscribing to a diff. depth stream
	wsHandler := func(event *binance_connector.WsDepthEvent) {
		fmt.Println(binance_connector.PrettyPrint(event))
	}

	errHandler := func(err error) {
		errMsg := fmt.Sprintf("Websocket Error: %v", err)
		logger.Error(errMsg)
	}

	_, _, err := websocketClient.WsDepthServe("BTCUSDT", wsHandler, errHandler)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to subscribe to Websocket Stream: %v", err)
		logger.Error(errMsg)
		return err
	}

	// Add logic to handle stream, like a go routine to keep it running or a way to stop it.
	return nil
}
