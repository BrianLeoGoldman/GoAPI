package api

// CoinBalanceParams parameters for the endpoint
type CoinBalanceParams struct {
	Username string
}

// CoinBalanceResponse response of the endpoint
type CoinBalanceResponse struct {
	// Success code, usually 200
	Code int
	// Account balance
	Balance int64
}

// Error the error to be returned if something failed
type Error struct {
	// Error code
	Code int
	// Error message
	Message string
}
