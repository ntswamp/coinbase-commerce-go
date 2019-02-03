package coinbase

import "os"

func NewTestClient() *CoinbaseClient {
	apiKey := os.Getenv("COINBASE_COMMERCE_API_KEY")
	return NewHttpClient(apiKey)
}
