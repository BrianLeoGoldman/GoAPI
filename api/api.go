package api

import (
	"encoding/json"
	"net/http"
)

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

// writeError returns an error to the caller of the endpoint
func writeError(w http.ResponseWriter, message string, code int) {
	// We create the Error struct
	resp := Error{
		Code:    code,
		Message: message,
	}
	// We set values to the header of the response
	w.Header().Set("Content-Type", "application/json")
	// We set the response code
	w.WriteHeader(code)
	// We set the response itself
	json.NewEncoder(w).Encode(resp)
}

// Wrapper to the writeError function
var (
	// The RequestErrorHandler is for specific error messages when there is a problem with the request
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	// The InternalErrorHandler is for generic error messages
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An unexpected error occurred", http.StatusInternalServerError)
	}
)
